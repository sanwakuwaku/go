package main

import "fmt"

func main() {
	fmt.Printf(execise519("abc", "def") + "\n")
}

// 練習問題5.19
// return文を含んでいないのに、ゼロ値ではない値を返す関数をpanicとrecoverを使って書きなさい。
func execise519(msg1, msg2 string) (result string) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case bailout{}:
			result = msg1 + ":" + msg2
		default:
			panic(p)
		}
	}()

	panic(bailout{})
}
