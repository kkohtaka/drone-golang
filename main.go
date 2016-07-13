// Copyright (c) 2016 Kazumasa Kohtaka. All rights reserved.
// This file is available under the MIT license.

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kkohtaka/drone-golang/pkg/service"
)

func main() {
	svc := service.NewService()

	quitCh := make(chan struct{})
	go func() {
		log.Println("HTTP service starts")
		svc.Run()
		close(quitCh)
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	select {
	case <-sig:
		log.Println("Received SIGINT, HTTP service exitting gracefully...")
		svc.Stop()
	case <-quitCh:
		log.Println("HTTP service exits")
	}
}
