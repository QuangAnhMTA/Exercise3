package main

import (
	"fmt"
	"strconv"
	"sync"
)

//tạo 1 biến X map[string]string và 3 goroutine cùng thêm dữ liệu vào X. Mỗi goroutine thêm 1000 key khác nhau.
//Sao cho quá trình đủ 15 key không mất mát dữ liệu. Lưu ý sử dụng mutex

func b2() {
	X := make(map[string]string)
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := "key" + strconv.Itoa(i) + strconv.Itoa(j)
				value := "value" + strconv.Itoa(i) + strconv.Itoa(j)
				mutex.Lock()
				X[key] = value
				mutex.Unlock()
			}
		}(i)
	}
	//time.Sleep(10 * time.Second)
	// for _, value := range X {
	// 	fmt.Println(value)
	// }
	wg.Wait()
	fmt.Println(len(X))

}
