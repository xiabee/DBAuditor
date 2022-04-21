package pre

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Readlines(filename string) []string {
	fd, err := os.Open(filename)
	var data []string
	if err != nil {
		fmt.Println("Read Error:", err)
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

func Get_files(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Dir Error:", err)
		os.Exit(0)
	}
	// 获取文件，并输出它们的名字
	var filename []string
	for _, file := range files {
		filename = append(filename, file.Name())
	}
	return filename
}

// 获得指定目录的文件名，返回一个切片

func HasPrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// 判断字符串是否以prefix开头

func HasSuffix(s string, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 判断字符串是否以suffix结尾
