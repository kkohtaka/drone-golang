// Copyright (c) 2016 Kazumasa Kohtaka. All rights reserved.
// This file is available under the MIT license.

package service

import (
	"fmt"
	"net"
	"net/http"

	"gopkg.in/redis.v4"
)

// Service represents a HTTP service
type Service struct {
	listener net.Listener
	client   *redis.Client
}

// NewService returns an instance of a HTTP service
func NewService() Service {
	return Service{}
}

// Run starts the HTTP service
func (s *Service) Run() {
	s.listener, _ = net.Listen("tcp", ":8080")

	s.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/entrypoint", func(w http.ResponseWriter, r *http.Request) {
		err := s.client.Set("key", "value", 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Hello, Web App.")
	})

	http.Serve(s.listener, mux)
}

// Stop stops the HTTP service
func (s *Service) Stop() {
	s.client.Close()
	s.listener.Close()
}
