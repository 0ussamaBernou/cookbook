package main

import "testing"

func Teststate(t testing.T) {
	var state int32

	state = 0
	for i := 0; i < 10; i++ {
		go func() {
			state++
		}()
	}
}
