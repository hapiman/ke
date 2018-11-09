package controllers

import (
	"fmt"
	"testing"
	"time"
)

func TestAutoSync(t *testing.T) {
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()
	fmt.Printf("h:%d,m:%d,s:%d \n", h, m, s)
}

func TestExecuteOnTime(t *testing.T) {
	AutoSync()
	select {}
}
