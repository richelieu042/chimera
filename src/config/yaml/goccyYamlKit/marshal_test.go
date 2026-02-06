package goccyYamlKit

import (
	"fmt"
	"testing"

	"github.com/goccy/go-yaml"
)

/*
例子：先后进行 反序列化 和 序列化，注释未丢失.
*/
func TestMarshal_KeepComments(t *testing.T) {
	type Config struct {
		Server struct {
			Port int    `yaml:"port"`
			Host string `yaml:"host"`
		} `yaml:"server"`

		Database struct {
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
		} `yaml:"database"`
	}

	yamlText := `
# 服务器
# aaa
# bbb
# ccc
server:
  port: 80	# 端口号
  host: 127.0.0.1

# 数据库
# ddd
database:
  user: root
  pass: <PASSWORD>
`

	var cfg Config

	// 用来存注释的 Map
	comments := yaml.CommentMap{}

	// 反序列化（收集注释）
	if err := yaml.UnmarshalWithOptions([]byte(yamlText), &cfg, yaml.CommentToMap(comments)); err != nil {
		panic(err)
	}

	// ==== 修改配置 ====
	cfg.Server.Port = 9090

	// 序列化（带回注释）
	out, err := yaml.MarshalWithOptions(cfg, yaml.WithComment(comments))
	if err != nil {
		panic(err)
	}

	// 输出查看
	fmt.Println(string(out))
}
