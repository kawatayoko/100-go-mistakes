package main

import (
	"fmt"

	"github.com/kawatayoko/100-go-mistakes/src/02-code-project-organization/03-init-function/redis"
)

// 1) 以下のコマンドを実行
// go run src/02-code-project-organization/03-init-function/main/main.go
//
// 2) 以下のように出力される
// redis init
// var
// init
// main
// redis store
//
// 3) 結果から以下のことがわかる
// 依存先のパッケージから初期化される。（依存先のinit()が先に呼ばれる）

var a = func() int {
	fmt.Println("var")
	return 0
}()

// - init関数(p14)
//   - 引数を取らない
//   - 結果も返さない
//   - パッケージが初期化される際にinit関数が実行される(main関数が呼ばれる前にinitがよばれる)
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	redis.Store("key", "value")
}
