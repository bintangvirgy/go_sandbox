package main

// for importing new package, run :
// 0. make sure go mod init
// 1. go get -u github.com/gorilla/mux
// 2. import package like this
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// implement router in golang using mux.
// mux adapt http module

func main() {
	// create new router
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "This is main page")
	})

	r.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "This is about page")
	})

	r.HandleFunc("/calculate", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Please insert number")
	})

	r.HandleFunc("/calculate/{fNum}/{operate}/{sNum}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		res := 0

		fmt.Fprintf(rw, "Calculating %s %s %s\n", vars["fNum"], vars["operate"], vars["sNum"])
		fNum, err1 := strconv.Atoi(vars["fNum"])
		sNum, err2 := strconv.Atoi(vars["sNum"])
		if err1 != nil || err2 != nil {
			fmt.Fprint(rw, "Number undetected")
		} else {
			calculate(fNum, vars["operate"], sNum, &res)
			fmt.Fprintf(rw, "Result : %s", strconv.Itoa(res))
		}

	})

	http.ListenAndServe(":8080", r)
}

func calculate(fNum int, operate string, sNum int, res *int) {
	switch operate {
	case "+":
		*res = fNum + sNum
	case "-":
		*res = fNum - sNum
	case "*":
		*res = fNum * sNum
	case "/":
		*res = fNum / sNum
	default:
		*res = 0
	}
}
