package main

import (
	"fmt"
	// "modtest"

	modtest "github.com/heifetzwu/lab1/modtest"

	log "github.com/siruspen/logrus"
)

func init() {
	fmt.Println("init")
}
func main() {
	log.Println("logger")
	fmt.Println("hello world")
	modtest.Exec()
}
func Exec() {
	fmt.Print("exec")
}
