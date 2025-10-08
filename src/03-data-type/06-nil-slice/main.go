package main

import "fmt"

func main() {
	var s1 []string
	s2 := []string(nil) // 空スライスを作成するシンタックスシュガー
	s3 := []string{}
	s4 := make([]string, 0)
	log(1, s1) // empty=false, nil=true
	log(2, s2) // empty=false, nil=true
	log(3, s3) // empty=true, nil=false
	log(4, s4) // empty=true, nil=false
}

func log(num int, s []string) {
	fmt.Printf("%d: empty=%t, nil=%t\n", num, len(s) == 0, s == nil)
}
