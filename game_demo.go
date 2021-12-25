package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-idm.wax.io/v1/accounts/auto-accept/login", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api-idm.wax.io")
	//req.Header.Set("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	//req.Header.Set("sec-ch-ua-platform", ""Windows"")
	req.Header.Set("accept", "*/*")
	req.Header.Set("origin", "https://wallet.wax.io")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://wallet.wax.io/")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cookie", "_ga=GA1.2.1094603490.1636344808; sid=e24c8a6f-defa-4b32-af3f-f2c7dfd06ce4; _gid=GA1.2.777938078.1637742864; _fbp=fb.1.1637742865027.1185552348; prism_223721173=8262616d-1389-4bba-ba4e-6feda4ed842e; _hjSessionUser_994621=eyJpZCI6ImRkOWRkNTBmLTA3OTctNTE3Ni04NjIzLWJkNmNjNDQ4ZWM5ZiIsImNyZWF0ZWQiOjE2Mzc3NDI4NjU5MzYsImV4aXN0aW5nIjp0cnVlfQ==; _hjSession_994621=eyJpZCI6ImQxNDcyZGMyLTNlODQtNDkwOS04ZTI0LTQ5OTU1OGYwNzFmNCIsImNyZWF0ZWQiOjE2Mzc4MjA3MDY1MDd9; _hjAbsoluteSessionInProgress=0; __ssid=d83c419fb4e00901e5e5e214cc3e890; session_token=QW2dUGnT0kF1GFKrRftLCPekb5E51qxqJ3cPhFVs")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
