package centrifugoKit

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/centrifugal/gocent/v3"
	"github.com/richelieu042/chimera/v3/src/idKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

func TestNewClient(t *testing.T) {
	addrs := []string{"http://localhost:8000/api"}
	//addrs := []string{"http://localhost:8000/apiii"}

	m := map[string]interface{}{
		"msg": "hellO" + idKit.NewUUID(),
	}
	jsonStr, err := jsonKit.MarshalToString(m)
	if err != nil {
		panic(err)
	}

	client, err := NewHttpClient(addrs, "4e75aa67-3423-4e10-86b3-cfab4febef55", nil)
	if err != nil {
		panic(err)
	}
	rst, err := client.Publish(context.TODO(), "test-channel", []byte(jsonStr))
	if err != nil {
		var codeErr gocent.ErrStatusCode
		if errors.As(err, &codeErr) {
			panic(codeErr.Code)
		}
		panic(err)
	}
	fmt.Println("Epoch:", rst.Epoch)
	fmt.Println("Offset:", rst.Offset)
}
