package main

import (
	"log"
	"sync"
)

//chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
//Lý giải nguyên nhân lỗi.
func errFunc() {
	var mutex = &sync.Mutex{}
	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mutex.Unlock()
			}
		}()
	}
	log.Println(len(m))
	log.Print("done")
}
