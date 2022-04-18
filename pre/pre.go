package pre

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Readlines(filename string) []string {
	fd, err := os.Open(filename)
	var data []string
	if err != nil {
		fmt.Println("read error:", err)
		os.Exit(0)
	}
	defer fd.Close()
	buff := bufio.NewReader(fd)
	for {
		line, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		data = append(data, string(line))
		// fmt.Println(string(data))
	}
	return data
}

// 解析文本，返回一个切片

func Concat(res []string, op string) string {
	data := ""
	for i := 0; i < len(res)-1; i++ {
		data += res[i]
		data += op
	}
	data += res[len(res)-1]
	return data
}

// 将切片拼接为字符串
