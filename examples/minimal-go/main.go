package main

import "github.com/PuRainDev/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example",
		"https://en.m.wikipedia.org/wiki/Main_Page", 800, 600, true, false)
}
