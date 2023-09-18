package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUserData(userId int, respCh chan<- string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respCh <- fmt.Sprintf("user %d data", userId)
	wg.Done()
}

func fetchUserRecommendation(userId int, respCh chan<- string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respCh <- fmt.Sprintf("user %d recommendation", userId)
	wg.Done()
}

func main() {
	start := time.Now()
	respCh := make(chan string, 128)

	wg := &sync.WaitGroup{}
	go fetchUserData(1, respCh, wg)
	wg.Add(1)
	go fetchUserRecommendation(1, respCh, wg)
	wg.Add(1)

	wg.Wait()
	close(respCh)
	for resp := range respCh {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(start))
}
