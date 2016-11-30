package main

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"golang.org/x/net/websocket"
)

var opts struct {
	OriginURL string `long:"origin" default:"http://localhost/" description:"Origin url"`
	WSURL     string `long:"ws" default:"ws://localhost:8327/media" description:"WebSocket URL to connect to"`
	Buffer    int    `long:"buffer" default:"1024" description:"WebSocket receive buffer size"`
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

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("error: %v", err)
		} else {
			return
		}
	}

	ws, err := websocket.Dial(opts.WSURL, "", opts.OriginURL)
	if err != nil {
		log.Fatal(err)
	}

	//message := []byte("hello world!")
	//_, err = ws.Write(message)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("WebSocket send: %v", message)

	msg := make([]byte, opts.Buffer)
	for {

		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("WebSocket receive: %v, %v", n, msg[:n])
	}
}
