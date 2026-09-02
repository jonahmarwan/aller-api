package main

import (
	"fmt"
)

// methods

func (e HeartBeatResponse) parse() string {
	if e.err != nil {
		e.body = ""
		return fmt.Sprintf("%v", e.err)
	}
	return fmt.Sprintf("%v", e.body)
}

func (e DatabaseResponse) parse() string {
	return fmt.Sprintf("%v: {%v} -> Error? %q", e.index, e.body, e.err)
}

func (e DatabaseRequest) setBody(b string) DatabaseRequest {
	e.body = b
	return DatabaseRequest{index: e.index, body: e.body, err: e.err}
}
func (e HeartBeatRequest) setBody(b string) HeartBeatRequest {
	e.body = b
	return HeartBeatRequest{body: e.body, err: e.err}
}
