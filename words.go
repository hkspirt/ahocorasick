package ahocorasick

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

//从csv生成敏感词正则表达式字串
func BuildSensitiveStr(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal("read sensitives_words.csv failed")
		return ""
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	ret := []string{}

	rd.ReadString('\n') //中文字段名
	rd.ReadString('\n') //类型
	rd.ReadString('\n') //字段名
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lineData := strings.Split(line, ",")
		if len(lineData) != 2 {
			continue
		}
		word := lineData[1]
		word = strings.Replace(word, " ", "", -1)
		word = strings.Replace(word, "\n", "", -1)
		word = strings.Replace(word, "\r", "", -1)
		word = strings.Replace(word, "|", "", -1)
		word = strings.Replace(word, ".", "", -1)
		word = strings.Replace(word, "+", "", -1)
		word = strings.Replace(word, "*", "", -1)
		word = strings.Replace(word, "?", "", -1)
		word = strings.Replace(word, "(", "", -1)
		word = strings.Replace(word, ")", "", -1)
		word = strings.Replace(word, "[", "", -1)
		word = strings.Replace(word, "]", "", -1)
		word = strings.Replace(word, "{", "", -1)
		word = strings.Replace(word, "}", "", -1)
		if word == "" {
			continue
		}
		ret = append(ret, word)
	}
	return strings.Join(ret, "|")
}

//从csv生成敏感词数组
func BuildSensitiveArray(path string) []string {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal("read sensitives_words.csv failed")
		return nil
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	ret := []string{}

	rd.ReadString('\n') //中文字段名
	rd.ReadString('\n') //类型
	rd.ReadString('\n') //字段名
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lineData := strings.Split(line, ",")
		if len(lineData) != 2 {
			continue
		}
		word := lineData[1]
		word = strings.Replace(word, " ", "", -1)
		word = strings.Replace(word, "\n", "", -1)
		word = strings.Replace(word, "\r", "", -1)
		if word == "" {
			continue
		}
		ret = append(ret, word)
	}
	return ret
}
