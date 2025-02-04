package redis

import "fmt"

func init() {
	fmt.Println("redis init")
}

func Store(key, value string) error {
	fmt.Println("redis store")
	return nil
}
