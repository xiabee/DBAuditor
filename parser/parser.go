package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Readlines(filename string) string {
	fd, err := os.Open(filename)
	var data byte
	if err != nil {
		fmt.Println("read error:", err)
	}
	defer fd.Close()
	buff := bufio.NewReader(fd)
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		fmt.Println(string(data))
	}
	return string(data)
}

// 按行读入文件

func flex() {
	return
}
