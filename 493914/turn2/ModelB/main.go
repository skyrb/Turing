package main

import (
    "fmt"
    "net/http"
)

type Request struct {
    method  string
    url     string
    headers http.Header
    body    []byte
}

func NewRequest(method string, url string, headers http.Header, body []byte) *Request {
    return &Request{
        method:  method,
        url:     url,
        headers: headers,
        body:    body,
    }
}

func main() {
    headers := http.Header{}
    headers.Add("Content-Type", "application/json")
    headers.Add("Authorization", "Bearer YOUR_TOKEN")