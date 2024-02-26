package main

import cmd "apiModules/ProductManagment/Cmd"

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("About Page Running About Handler"))
// 	w.WriteHeader(http.StatusOK)
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Index Page Running Index Handler"))
// }

func main() {

	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/index", indexHandler)
	// http.HandleFunc("/about", aboutHandler)
	// http.ListenAndServe(":8080", nil)

	cmd.Cmd()

}
