package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	getUser()
}

func sign(usr string) {
	url := "https://h5.ele.me/restapi/member/v2/users/" + usr + "/sign_in"
	post(url)
}

func getUser() {
	url := "https://h5.ele.me/restapi/eus/v2/current_user?info_raw={}"
	usr := get(url)
	sign(usr)
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

func get(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req.Header)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func post(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, nil)
	setHeaders(req.Header)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println((string(body)))
}
