package socket

//import (
//	"fmt"
//	"github.com/gorilla/websocket"
//	"net/http"
//)
//
//var wsupgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//func Wshandler(w http.ResponseWriter, r *http.Request) {
//	conn, err := wsupgrader.Upgrade(w, r, nil)
//	if err != nil {
//		fmt.Println("Failed to set websocket upgrade: %+v", err)
//		return
//	}
//
//	for {
//		t, msg, err := conn.ReadMessage()
//
//		if err != nil {
//			break
//		}
//
//		conn.WriteMessage(t, msg)
//	}
//}
