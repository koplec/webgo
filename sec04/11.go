package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hoge struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
  <head><title>Go Web Programming</title></head>
  <body><h1>Hello World</h1></body>
  </html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Println(w, "そのようなサービスはありません。他を当たってください。")
}

//redirect sample
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.yahoo.co.jp")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Hoge{
		User:    "koplec78",
		Threads: []string{"1盤", "2番", "将棋"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
