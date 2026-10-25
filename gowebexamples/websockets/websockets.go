package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/coder/websocket"
)

func echoHandle(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{Subprotocols: []string{"echo"}, InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return
	}
	defer c.CloseNow()
	if c.Subprotocol() != "echo" {
		err := c.Close(websocket.StatusPolicyViolation, "Only ECHO over here!")
		if err != nil {
			log.Println(err)
		}
		return
	}
	ctx := r.Context()
	for {
		typ, cr, err := c.Reader(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		cw, err := c.Writer(ctx, typ)
		if err != nil {
			log.Println(err)
			return
		}
		tee := io.TeeReader(cr, os.Stdout)
		_, err = io.Copy(cw, tee)
		if err != nil {
			log.Println(err)
			return
		}
		err = cw.Close()
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/echo", echoHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
