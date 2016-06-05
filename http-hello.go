package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world!\n")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		// パスの1文字目(すなわち"/")以降を表示
}

func main() {
	flag.Parse()
	var subdirectory string = "/"
	if flag.NArg() > 0 {
		subdirectory += flag.Arg(0)
	}
	fmt.Printf(subdirectory)
	http.HandleFunc(subdirectory, handler)
	http.ListenAndServe(":8080", nil)
}
