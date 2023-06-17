package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	// 終了していないgoroutineがあるかチェック
	// goroutine leakがあるとテストが失敗する
	defer goleak.VerifyNone(t)
	main()

}
