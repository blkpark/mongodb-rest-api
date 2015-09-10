package main

import (
	"os"
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"

    "github.com/ghmlee/mongodb-rest-api/context"
    "github.com/ghmlee/mongodb-rest-api/route"
    "github.com/ghmlee/mongodb-rest-api/mongodb"

	"github.com/gorilla/mux"
)

var (
    port = getEnv("PORT", "8888")
    host = getEnv("HOST", "127.0.0.1")
    db = getEnv("DATABASE", "default")
    client = newMongoDB()
)

// to get an environment variable if it exists or default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func newMongoDB() *mongodb.MongoDB {
    m, _ := mongodb.NewMongoDB(host, db)
    return m
}

// A middleware to serve a context struct
type MuxHandler func(context.Context, http.ResponseWriter, *http.Request) (int, map[string]interface{})

func serveContext(next MuxHandler) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		queryParams := r.URL.Query()
		c := context.Context{
			mux.Vars(r),
			body,
			queryParams,
            client,
		}
        
		var js []byte
		var length string

		status, res := next(c, w, r)
		if res == nil {
			js = []byte("")
			length = "0"
		} else {
			js, _ = json.Marshal(res)
			length = strconv.Itoa(len(js))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", length)
		w.WriteHeader(status)
		w.Write(js)

		return
	})
}

// to run
func run() {
	mux := mux.NewRouter()

    mux.HandleFunc("/{collection}", serveContext(route.PostDocument)).Methods("POST")

	http.Handle("/", mux)
	http.ListenAndServe(":" + port, nil)
}

func init() {
}

func main() {
	run()
}
