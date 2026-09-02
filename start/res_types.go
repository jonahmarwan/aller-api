package main

// interfaces

type Response interface {
	parse()
}

type Request[T any] interface {
	set_body() T
}

// structs

type HeartBeatResponse struct {
	body string
	err  error
}
type HeartBeatRequest struct {
	body string
	err  error
}
type DatabaseResponse struct {
	index int
	body  string
	err   error
}

type DatabaseRequest struct {
	index int
	body  string
	err   error
}
