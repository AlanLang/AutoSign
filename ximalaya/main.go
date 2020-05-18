package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// User 用户信息
type User struct {
	isTickedToday bool
}

func main() {
	url := "https://m.ximalaya.com/starwar/lottery/check-in/record"
	fmt.Println(getUser(url).isTickedToday)
	getacc()
}

func getacc() {
	url := "https://m.ximalaya.com/starwar/task/listen/account"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req.Header)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	println("签到结果:", string(body))
}

func setHeaders(header http.Header) {
	fi, err := os.Open("./header")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		headerKeyValue := strings.Split(string(a), ":")
		header.Set(headerKeyValue[0], strings.TrimSpace(headerKeyValue[1]))
	}
}

func getUser(url string) User {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req.Header)
	resp, _ := client.Do(req)
	var u User
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &u)
	return u
}
