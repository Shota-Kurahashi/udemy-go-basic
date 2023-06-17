package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")

	if err != nil {
		// 強制的にプログラムを終了する
		log.Fatal(err)
	}

	defer file.Close()

	// logに追加の情報を付与する
	flags := log.Lshortfile | log.LstdFlags
	warnLogger := log.New(io.MultiWriter(file, os.Stdout), "Warning: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "Error: ", flags)

	warnLogger.Println("Warning: This is a warning message")
	errorLogger.Fatalln("Error: This is an error message")
}
