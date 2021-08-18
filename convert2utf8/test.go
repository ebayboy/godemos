package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/axgle/mahonia"
	"github.com/saintfish/chardet"
)

/*
mapping:
detector ->  decoder
UTF-8 -> utf8
ISO-8859-1 -> ISO-8859-1
GB-18030 -> gb18030
*/

func Str2UTF8(src string, datautf8 string) error {
	//建立charset识别字符集到转码字符集的映射
	chardet_to_chardet := map[string]string{
		"UTF-8":      "utf8",
		"ISO-8859-1": "ISO-8859-1",
		"GB-18030":   "gb18030",
		"Big5":       "big5",
	}

	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest([]byte(src))
	if err != nil {
		fmt.Println("continue utf8!")
		return err
	}

	if charset.Charset == "UTF-8" {
		datautf8 = src
		return nil
	}

	decCharset, exist := chardet_to_chardet[charset.Charset]
	if !exist {
		return errors.New(fmt.Sprintf("Error: chardet[%s] to [%s] not exit!", charset, decCharset))
	}

	d := mahonia.NewDecoder(decCharset)
	if d == nil {
		return errors.New(fmt.Sprintf("Error:Could not create decoder for %s convert to  %s", charset, decCharset))
	}

	datautf8 = d.ConvertString(string(src))

	fmt.Printf("[%s] -> [%s] datautf8:[%s]\n", charset.Charset, decCharset, datautf8)

	return nil
}

func main() {
	//8859-1.txt
	//utf-8.txt

	f, err := os.Open("./8859-1.txt")
	//f, err := os.Open("./utf8.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}

		fmt.Println("=====line:", string(line))
		datautf8 := ""
		Str2UTF8(line, datautf8)
		fmt.Println("Done=====datautf8:", datautf8)
	}
}
