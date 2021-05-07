package main

import "simpleDbMirror/app"

func main() {

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("test!!!!!"))
	//})
	//http.ListenAndServe(":4000", nil)

	app.StartApplication()
}
