package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/PuRainDev/webview"
)

const (
	windowWidth  = 480
	windowHeight = 320
)

var indexHTML =`
<!doctype html>
<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Georama&display=swap" rel="stylesheet">
		<style>body { margin: 0; padding: 0; box-sizing: border-box; color: white; background-color: #16161f; font-family: 'Georama', sans-serif;}
		#header { cursor: default; text-align: left; background-color: black; padding: 10px 10px 10px 15px; font-size: 15px; display: flex; justify-content: space-between;}
		#close_btn {color: #ffffff; text-align: right; cursor: pointer; transition: .2s;}
		#close_btn:hover {color: #00a3d9;}</style>
	</head>
	<body style="text-align: center;">
		<div id="header">Simple frameless demo<span id="close_btn">X</span></div>
		<br>
		<br>
		<br>
		<br>
		<h1>Frameless-go</h1>
		<p>Simple example window without border</p>
		<br>
		<br>
		<br>
		<br>
		<br>
	</body>
	<script>
	header.onmouseover = function(){
		external.invoke('drag_on');
	}
	header.onmouseout = function(){
		external.invoke('drag_false');
	}
	close_btn.onclick = function(){
		external.invoke('close');
	}
	</script>
</html>
`

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(indexHTML))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func handleRPC(w webview.WebView, data string) {
	fmt.Println(data);
	switch {
	case data == "close":
		w.Terminate()
		w.Exit()
	case data == "drag_on":
		w.SetDraggable(true)
	case data == "drag_false":
		w.SetDraggable(false)
	}
}

func main() {
	url := startServer()
	w := webview.New(webview.Settings{
		Width:  windowWidth,
		Height: windowHeight,
		Title:  "Simple frameless demo",
		URL:    url,
		ExternalInvokeCallback: handleRPC,
		Borderless: true,
	})
	w.Run()
}
