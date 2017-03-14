package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/Akagi201/light"
	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"golang.org/x/net/websocket"
)

var opts struct {
	ListenAddr string `long:"listen" default:"0.0.0.0:8327" description:"HTTP listen address and port"`
	StaticPath string `long:"static" default:"./static" description:"Path of static files"`
	MediaFile  string `long:"media" default:"./static/test.mp4" description:"Path of media file for vod"`
	Buffer     int    `long:"buffer" default:"10240000" description:"Send media buffer size"`
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
	b := make([]byte, opts.Buffer)

	f, err := os.OpenFile(opts.MediaFile, os.O_RDONLY, 0600)
	if err != nil {
		log.Printf("Open file failed, err: %v", err)
		return
	}
	defer f.Close()
	for {
		n, err := f.Read(b)
		if err != nil {
			log.Printf("Read file failed, err: %v", err)
			return
		}
		err = websocket.Message.Send(ws, b[0:n])
		if err != nil {
			log.Printf("WebSocket send media buffer failed, err: %v", err)
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

	root.ServeFiles("/vod/*filepath", http.Dir("./static"))
	root.Get("/", http.RedirectHandler("/vod", http.StatusFound).ServeHTTP)
	root.Get("/media", websocket.Handler(handleMedia).ServeHTTP)

	log.Printf("HTTP listening at: %v", opts.ListenAddr)
	http.ListenAndServe(opts.ListenAddr, root)
}
