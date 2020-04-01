package main

# hotfix email address
#issue 53 
# issue 53 - finish the new footer

#test

# test2

# test 3

# test 4

# Test FETCH_HEAD

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f := fib()

		res := &response{Message: "Happy Friday"}

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			res.EnvVars = append(res.EnvVars, pair[0]+"="+pair[1])
		}

		for i := 1; i <= 90; i++ {
			res.Fib = append(res.Fib, f())
		}

		// Beautify the JSON output
		out, _ := json.MarshalIndent(res, "", "  ")

		// Normally this would be application/json, but we don't want to prompt downloads
		w.Header().Set("Content-Type", "text/plain")

		io.WriteString(w, string(out))

		fmt.Println("Hello world - the log message")
	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string   `json:"message"`
	EnvVars []string `json:"env"`
	Fib     []int    `json:"fib"`
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

#test3

#test4
