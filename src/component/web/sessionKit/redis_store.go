package sessionKit

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/randomKit"
	"github.com/richelieu042/chimera/v3/src/serialize/gobKit"

	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
)

// RedisStore stores gorilla sessions in Redis
type RedisStore struct {
	// Richelieu: MaxAge == 0的情况下，session键的超时时间
	expirationForZeroMaxAge time.Duration

	// client to connect to redis
	client redis.UniversalClient
	// default options to use when a new session is created
	options sessions.Options
	// key prefix with which the session will be stored
	keyPrefix string
	// key generator
	keyGen KeyGenFunc
	// session serializer
	serializer SessionSerializer
}

// KeyGenFunc defines a function used by store to generate a key
type KeyGenFunc func() (string, error)

// NewRedisStore returns a new RedisStore with default configuration
/*
PS:
(1) cookie的value 由 传参keyGen 生成（即Redis中session键的后半部分）.

e.g.
	同一个Session，cookie的值: "01H815Q2YFYEFDZHAQ2478FFS9", Redis中session键: "session:01H815Q2YFYEFDZHAQ2478FFS9".

@param redisKeyPrefix 			Redis中session键的前半部分（固定）
@param keyGen					Redis中session键的后半部分（不固定，动态生成）|| cookie的value
								(a) 可以为nil，将采用默认值
@param opts 					Cookie的属性
@param expirationForZeroMaxAge	(a) sessions.Options.MaxAge == 0的情况下，此属性才有效
								(b) 传值可以参考 redisKit.Client 的 Set().
*/
func NewRedisStore(ctx context.Context, client redis.UniversalClient, redisKeyPrefix string, keyGen KeyGenFunc,
	opts sessions.Options, expirationForZeroMaxAge time.Duration) (*RedisStore, error) {
	// Richelieu
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	store := &RedisStore{
		expirationForZeroMaxAge: expirationForZeroMaxAge,
		//options: sessions.Options{
		//	Path:   "/",
		//	MaxAge: 86400 * 30,
		//},
		client: client,
		//keyPrefix:  "session:",
		//keyGen:     generateRandomKey,
		serializer: GobSerializer{},
	}

	// Redis中session键的前半部分（固定）
	store.KeyPrefix(redisKeyPrefix)

	// Redis中session键的后半部分（不固定，动态生成）
	if keyGen == nil {
		keyGen = generateRandomKey
	}
	store.KeyGen(keyGen)

	store.Options(opts)

	return store, nil
}

// Get returns a session for the given name after adding it to the registry.
func (s *RedisStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(s, name)
}

// New returns a session for the given name without adding it to the registry.
func (s *RedisStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(s, name)
	opts := s.options
	session.Options = &opts
	session.IsNew = true

	c, err := r.Cookie(name)
	if err != nil {
		return session, nil
	}
	session.ID = c.Value

	err = s.load(r.Context(), session)
	if err == nil {
		session.IsNew = false
	} else if err == redis.Nil {
		err = nil // no data stored
	}
	return session, err
}

// Save adds a single session to the response.
//
// If the Options.MaxAge of the session is <= 0 then the session file will be
// deleted from the store. With this process it enforces the properly
// session cookie handling so no need to trust in the cookie management in the
// web browser.
func (s *RedisStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	// Delete if max-age is <= 0

	// Richelieu
	//if session.Options.MaxAge <= 0 {
	if session.Options.MaxAge < 0 {
		if err := s.delete(r.Context(), session); err != nil {
			return err
		}
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return nil
	}

	if session.ID == "" {
		id, err := s.keyGen()
		if err != nil {
			return errors.New("redisstore: failed to generate session id")
		}
		session.ID = id
	}
	if err := s.save(r.Context(), session); err != nil {
		return err
	}

	http.SetCookie(w, sessions.NewCookie(session.Name(), session.ID, session.Options))
	return nil
}

// Options set options to use when a new session is created
func (s *RedisStore) Options(opts sessions.Options) {
	s.options = opts
}

// KeyPrefix sets the key prefix to store session in Redis
func (s *RedisStore) KeyPrefix(keyPrefix string) {
	s.keyPrefix = keyPrefix
}

// KeyGen sets the key generator function
func (s *RedisStore) KeyGen(f KeyGenFunc) {
	s.keyGen = f
}

// Serializer sets the session serializer to store session
func (s *RedisStore) Serializer(ss SessionSerializer) {
	s.serializer = ss
}

// Close closes the Redis store
func (s *RedisStore) Close() error {
	return s.client.Close()
}

// save writes session in Redis
/*
@param genFlag Richelieu: true: session.ID是新生成的
*/
func (s *RedisStore) save(ctx context.Context, session *sessions.Session) error {
	b, err := s.serializer.Serialize(session)
	if err != nil {
		return err
	}

	// Richelieu
	//return s.client.Set(ctx, s.keyPrefix+session.ID, b, time.Duration(session.Options.MaxAge)*time.Second).Err()
	var expiration time.Duration
	if session.Options.MaxAge == 0 {
		expiration = s.expirationForZeroMaxAge
	} else {
		expiration = time.Duration(session.Options.MaxAge) * time.Second
	}
	if session.IsNew {
		// 要避免: key在Redis中已存在（uuid、ulid等并不可靠）
		for i := 0; i < 3; i++ {
			ok, err := s.client.SetNX(ctx, s.keyPrefix+session.ID, b, expiration).Result()
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
			// 重复了，需要重新生成 session.ID
			id, err := s.keyGen()
			if err != nil {
				return errorKit.Newf("fail to regenerate session id")
			}
			session.ID = id + "_" + strconv.Itoa(randomKit.Int(0, 123456))
		}
		return errorKit.Newf("multiple repetition")
	}
	return s.client.Set(ctx, s.keyPrefix+session.ID, b, expiration).Err()
}

// load reads session from Redis
func (s *RedisStore) load(ctx context.Context, session *sessions.Session) error {
	cmd := s.client.Get(ctx, s.keyPrefix+session.ID)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	b, err := cmd.Bytes()
	if err != nil {
		return err
	}

	return s.serializer.Deserialize(b, session)
}

// delete deletes session in Redis
func (s *RedisStore) delete(ctx context.Context, session *sessions.Session) error {
	return s.client.Del(ctx, s.keyPrefix+session.ID).Err()
}

// SessionSerializer provides an interface for serialize/deserialize a session
type SessionSerializer interface {
	Serialize(s *sessions.Session) ([]byte, error)
	Deserialize(b []byte, s *sessions.Session) error
}

type GobSerializer struct{}

func (gs GobSerializer) Serialize(s *sessions.Session) ([]byte, error) {
	// gob序列化
	data, err := gobKit.Marshal(s.Values)
	if err != nil {
		return nil, err
	}
	// base64编码（Richelieu: 额外hex编码是为了防国产TongRDS，set后get的值有问题）
	return base64Kit.Encode(data), nil

	//buf := new(bytes.Buffer)
	//enc := gob.NewEncoder(buf)
	//err := enc.Encode(s.Values)
	//if err != nil {
	//	return nil, err
	//}
	//data := buf.Bytes()
	//return hexKit.Encode(data), nil
}

func (gs GobSerializer) Deserialize(d []byte, s *sessions.Session) error {
	// base64解码
	data, err := base64Kit.Decode(d)
	if err != nil {
		return err
	}
	// gob反序列化
	return gobKit.Unmarshal(data, &s.Values)

	//dec := gob.NewDecoder(bytes.NewBuffer(d))
	//return dec.Decode(&s.Values)
}

// generateRandomKey returns a new random key
func generateRandomKey() (string, error) {
	k := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return "", err
	}
	return strings.TrimRight(base32.StdEncoding.EncodeToString(k), "="), nil
}
