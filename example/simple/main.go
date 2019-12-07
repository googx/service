// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/googx/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Printf("start...==>%v  \n", s)

	go p.run()
	return nil
}
func (p *program) run() {
	t := time.NewTicker(time.Second)
	for {
		i := <-t.C
		fmt.Printf("do dome work:%v \n", i)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Printf("stop...==>%v  \n", s)

	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	sysstatus, err := s.Status()
	if err != nil {
		log.Fatalf("查看服务状态失败 %v \n ", err)
	}
	log.Printf("服务状态信息:%v", sysstatus)

	log.Printf("s.Platform():%v", s.Platform())

	if err := s.Install(); err != nil {
		log.Printf("安装服务失败:%v", err)
		// 卸载服务
		if err := s.Uninstall(); err != nil {
			log.Printf("卸载服务失败:%v", err)
		} else {
			log.Printf("卸载服务成功...")
		}
	} else {
		log.Printf("安装服务成功...")
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
