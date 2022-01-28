package main

import (
	"fmt"
	"net/http" //package used to build web applications.
)

func main() {
	http.HandleFunc("/", Myfunc1) //Handlefunc()  routes the incoming request to a specified function.
	http.HandleFunc("/John", Myfunc2)
	http.ListenAndServe(":8080", nil) //used to create a remote server which listens to the particularrv host.
}
func Myfunc1(w http.ResponseWriter, r *http.Request) { //takes the incoming request and after processing it provides the response.
	fmt.Fprint(w, "Hello World\n")
}
func Myfunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello John\n")
}
