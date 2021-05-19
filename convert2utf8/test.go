package main

import (
	"bufio"
	"github.com/saintfish/chardet"
	"io"
	"os"
)

func main() {
	//8859-1.txt
	//utf-8.txt

	f, err := os.Open("./8859-1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}

		detector := chardet.NewTextDetector()
		charset, err := detector.DetectBest(line)
		if err != nil {
			panic(err)
		}

		println("=====line:", string(line))
		println("charset:", charset.Charset, " charset.Language:", charset.Language)
	}

	return

	rawBytes := []byte("some text")
	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(rawBytes)
	if err != nil {
		panic(err)
	}

	println(charset.Charset)
	println(charset.Language)
}
