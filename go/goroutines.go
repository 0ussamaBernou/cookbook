package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4"}
var results = []string{}

func dbCall(i int) {
	defer wg.Done()

	delay := time.Second
	time.Sleep(delay)
	fmt.Println("Result from db is:", dbData[i])
	save(dbData[i])
	log()
}

func save(result string) {
	// position of the lock is crucial (lock just the parts that can race)
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("Current results are: ", results)
	m.RUnlock()
}

func main() {
	t0 := time.Now()

	for i := range dbData {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Println("Total exection time:", time.Since(t0))
	fmt.Println("Resluts: ", results)
}
