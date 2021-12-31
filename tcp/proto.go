package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
)

//对数据进行封包/拆包，解决tcp发送消息粘在一起的问题
//包格式：
//32位消息长度 + 消息
//思想就是加一个头，包含消息的长度

func Encode(msg string) ([]byte, error) {
	//func binary.Write(w io.Writer, order binary.ByteOrder, data interface{}) error
	//写入消息头

	//var pkg = new(bytes.Buffer)
	var pkg bytes.Buffer
	var mlen uint32 = uint32(len(msg))

	//写入长度
	err := binary.Write(&pkg, binary.LittleEndian, mlen)
	if err != nil {
		return nil, err
	}

	//写入消息
	err = binary.Write(&pkg, binary.LittleEndian, ([]byte)(msg))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("Error: invalid data!")
	}
	dlen := binary.LittleEndian.Uint32(data[:4])
	log.Printf("Decode dlen:%v\n", dlen)
	msg := string(data[4:])
	return msg, nil
}

func main() {
	msg := "hello world"
	bmsg, err := Encode(msg)
	if err != nil {
		log.Printf("Error:%v\n", err.Error())
	}

	//前32bit(4字节)是长度， 后面是msg
	log.Printf("bmsg:[%v]\n", string(bmsg[4:]))

	s, err := Decode(bmsg)
	if err != nil {
		log.Fatalf("Error:%v\n", err.Error())
	}
	log.Printf("s:[%v]\n", s)
}
