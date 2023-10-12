package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 10

func checkBargain(url string, ch chan string, max_price float32) {
	for {
		time.Sleep(time.Second)
		price := rand.Float32() * 20
		if price > max_price {
			ch <- url
			// close(ch)
			break
		}
	}
}

func main() {
	chickenCh := make(chan string)
	tofuCh := make(chan string)
	urls := []string{"walmart.com", "costco.com", "wholefoods.com"}
	for _, url := range urls {
		go checkBargain(url, chickenCh, MAX_CHICKEN_PRICE)
		go checkBargain(url, tofuCh, MAX_TOFU_PRICE)
	}

	select {
	case url := <-chickenCh:
		fmt.Printf("Chicken Bargain found on : %v", url)
	case url := <-tofuCh:
		fmt.Printf("tofu Bargain found on : %v", url)
	}
}
