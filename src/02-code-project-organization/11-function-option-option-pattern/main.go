package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	_, err := NewServer("localhost", WithPort(8080), WithTimeout(180))
	if err != nil {
		fmt.Printf("error!!")
	}

	fmt.Print("finish")
}

func WithPort(port int) Option {
	return func(op *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		op.port = &port
		return nil
	}
}

func WithTimeout(second int) Option {
	return func(op *options) error {
		if second < 0 {
			return errors.New("second should be positive")
		}
		op.timeout = &second
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options

	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = 80
	} else {
		port = *options.port
	}

	fmt.Printf("port is %d\n", port)

	var timeout int
	if options.timeout == nil {
		timeout = 60
	} else {
		timeout = *options.timeout
	}

	fmt.Printf("timeout is %d\n", timeout)

	return &http.Server{}, nil
}

type Option func(*options) error

type options struct {
	port    *int
	timeout *int
}
