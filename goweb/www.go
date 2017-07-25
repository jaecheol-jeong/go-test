package main

import (
	"log"
	"net/http"
	"time"

	calc "goweb/controllers"
	m "goweb/module"
)

func main() {

	http.HandleFunc("/exit", func(w http.ResponseWriter, req *http.Request) {
		go writeLog("WWW", "Good Bye!", req.RequestURI)
		time.Sleep(500 * time.Millisecond)

		w.Write([]byte("the end"))
		log.Fatalln()
	})

	htmlResponse("/", "<h3 style='color:black;'>welcome home</h3>")

	htmlResponse("/hab", calc.Hab2Html(10, 5))

	println("http://localhost:5000/")
	http.ListenAndServe(":5000", nil)
}

func htmlResponse(router string, html string) {

	http.HandleFunc(router, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))

		go writeLog("WWW", req.RemoteAddr, req.RequestURI)
	})

}

func writeLog(title string, log string, etc string) {
	if len(etc) >= 4 && etc[len(etc)-4:] == ".ico" {
		return
	}
	go func() {
		l := m.LOG{Title: title, Log: log, Etc: etc}
		l.SetLog()
	}()
}
