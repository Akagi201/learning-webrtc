package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Akagi201/light"
	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"golang.org/x/net/websocket"
)

var opts struct {
	ListenAddr string `long:"listen" default:"0.0.0.0:8327" description:"HTTP listen address and port"`
	StaticPath string `long:"static" default:"./static" description:"Path of static files"`
	HTTPFLVURL string `long:"httpflv" default:"http://uplive.b0.upaiyun.com/live/loading.flv" description:"HTTP flv stream url"`
	Buffer     int    `long:"buffer" default:"1024" description:"Send media buffer size"`
}

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.InfoLevel
	f := new(logrus.TextFormatter)
	f.TimestampFormat = "2006-01-02 15:04:05"
	f.FullTimestamp = true
	log.Formatter = f
}

func handleMedia(ws *websocket.Conn) {
	r := ws.Request()

	log.Println("Proxy client", r.RemoteAddr, "to", opts.HTTPFLVURL)
	defer ws.Close()

	flv, err := http.Get(opts.HTTPFLVURL)
	if err != nil {
		fmt.Println("Connect backend", opts.HTTPFLVURL, "failed, err is", err)
		return
	}
	defer flv.Body.Close()

	b := make([]byte, opts.Buffer)
	for {
		var n int
		if n, err = flv.Body.Read(b); err != nil {
			fmt.Println("Recv from backend failed, err is", err)
			return
		}

		if err = websocket.Message.Send(ws, b[:n]); err != nil {
			fmt.Println("Send to ws failed, err is", err)
			return
		}
	}
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("error: %v", err)
		} else {
			return
		}
	}

	root := light.New()

	root.ServeFiles("/live/*filepath", http.Dir("./static"))
	root.Get("/", http.RedirectHandler("/live", http.StatusFound).ServeHTTP)
	root.Get("/media", websocket.Handler(handleMedia).ServeHTTP)

	log.Printf("HTTP listening at: %v", opts.ListenAddr)
	http.ListenAndServe(opts.ListenAddr, root)
}
