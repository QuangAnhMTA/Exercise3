package main

import (
	"log"
	"sync"
	"time"
)

/*
dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message.
func chanRoutine() {
    log.Print("hello 1")
    go func() {
        time.Sleep(1 * time.Second)
        log.Print("hello 3")
    }
    log.Print("hello 2")
}

-- nâng cao. In ra các message theo thứ tự. -- In ra message 3 trước message 2.
Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)
*/
// cách 1: Dùng time sleep đơn giản
// func chanRoutine() {

// 	log.Print("hello 1")
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		log.Print("hello 3")
// 	}()
// 	time.Sleep(1 * time.Second)
// 	log.Print("hello 2")
// 	time.Sleep(5 * time.Second)

// }
// cách 2: Dùng chan
// func chanRoutine() {
// 	done := make(chan bool)
// 	log.Print("hello 1")
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		log.Print("hello 3")
// 		done <- true
// 	}()
// 	<-done
// 	log.Print("hello 2")
// }

//Cách 3: Dùng mutex

// Cách 4: Waitgroup
func chanRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")

}
