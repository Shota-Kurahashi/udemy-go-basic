package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func funcDefer() {
	// 関数をdeferで呼び出すと、その関数の最後に実行される
	// 複数ある場合は、後ろから順に実行される
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))

	for _, file := range files {
		out = append(out, strings.TrimSuffix(file, filepath.Ext(file)))
	}

	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)

	if err != nil {
		return "", errors.New("file not found")
	}

	defer f.Close()

	return f.Name(), nil
}

func addExtension(f func(file string) string, name string) {
	fmt.Println(f(name))
}

func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	count := 0

	return func(i int) int {
		return count + i
	}
}

func main() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))

	name, err := fileChecker("file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name)

	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)

	f1 := func(i int) int {
		return i + 1
	}

	fmt.Println(f1(2))

	f2 := func(file string) string {
		return file + ".csv"
	}

	addExtension(f2, "file")

	f3 := multiply()

	fmt.Println(f3(2))

	f4 := countUp()

	for i := 0; i < 10; i++ {
		fmt.Println(f4(i))
	}

}
