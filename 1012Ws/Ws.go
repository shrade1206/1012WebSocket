package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	fmt.Println("Go WebSocket")

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
}

// func main() {
//     http.HandleFunc("/", handler)
//     if err := http.ListenAndServe(":8080", nil); err != nil {

// 		log.Fatal("ListenAndServe:", err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("WebSocket Test"))
// }

func Echo(ws *websocket.Conn) {

	var err error

	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("無接收訊息")
			break
		}
		fmt.Println("收到使用者訊息: " + reply)

		msg := "後台已收到"

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("傳送失敗")
			break
		}
	}
}
