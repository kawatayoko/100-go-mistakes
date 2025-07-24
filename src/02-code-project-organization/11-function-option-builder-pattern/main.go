package main

import (
	"errors"
	"fmt"
)

// 関数オプションビルダーパターン
func main() {
	builder := ConfigBuilder{}
	cfg, err := builder.Port(8080).Schema("http").Build()
	if err != nil {
		fmt.Println("error!!")
		return
	}
	fmt.Printf("success builder: %+v \n", cfg)
}

type Config struct {
	Port   int
	Schema string
}

type ConfigBuilder struct {
	port   *int
	schema *string
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Schema(schema string) *ConfigBuilder {
	b.schema = &schema
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}
	if b.port == nil {
		cfg.Port = 80
	} else {
		if *b.port <= 0 {
			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}
	if b.schema == nil {
		cfg.Schema = "https"
	} else {
		cfg.Schema = *b.schema
	}
	return cfg, nil
}
