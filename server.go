package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type timeResponse struct {
	Time string
}

func getTime(w http.ResponseWriter, req *http.Request) {
	dt := time.Now()
	if req.Method != "GET" {
		return
	}
	res := &timeResponse{Time: dt.Format(time.RFC3339)}
	jsonRes, _ := json.Marshal(res)
	io.WriteString(w, string(jsonRes))
}

func main() {
	http.HandleFunc("/time", getTime)
	http.ListenAndServe(":8795", nil)
}
