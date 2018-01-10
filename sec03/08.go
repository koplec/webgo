//ハンドラ関数による実装
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HelloHello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Worlddd!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3008",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	//ハンドラ関数をHandlerに変換
	helloHandler := http.HandlerFunc(hello)
	http.Handle("/hello2", &helloHandler)

	server.ListenAndServe()
}
