package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var debugEnabled = false

func debugf(format string, args ...interface{}) {
	if debugEnabled {
		log.Printf(format, args...)
	}
}

type RPC struct {
	addr string
}

func NewRPC(addr string) *RPC {
	return &RPC{
		addr: addr,
	}
}

type Request struct {
	Method   string
	Resource string
	Query    url.Values
	Headers  http.Header
	In       interface{}
	Out      interface{}
}

func (r *RPC) Do(req *Request) (http.Header, error) {
	url := r.addr + req.Resource
	if req.Query != nil {
		url += "?" + req.Query.Encode()
	}

	var in_body io.Reader
	if req.In != nil {
		data, err := json.Marshal(req.In)
		if err != nil {
			return nil, err
		}
		in_body = bytes.NewReader(data)
	}

	http_req, err := http.NewRequest(req.Method, url, in_body)
	if err != nil {
		return nil, err
	}
	if req.Headers != nil {
		http_req.Header = req.Headers
	}
	http_req.Header.Set("Content-Type", "application/json")
	http_req.Header.Set("Accept", "application/json")
	// debugf("%+v", http_req)

	resp, err := http.DefaultClient.Do(http_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// debugf("response: %+v", resp)
	debugf("body: %s", out_body)

	if req.Out == nil {
		return resp.Header, nil
	}
	return resp.Header, json.Unmarshal(out_body, req.Out)
}

var prefix = []string{
	"420Weedbot",
	"YoloBlazeit",
	"Jerkbot",
	"DankMemes",
	"Gross",
	"Loser",
}

var suffix = []string{
	" 2000",
	" Yolo",
	" SteelBeams",
	" InsideJob",
	"inator",
	" ==>",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomName() string {
	return "JEFF WENDLING " + prefix[rand.Intn(len(prefix))] +
		suffix[rand.Intn(len(suffix))]
}
