package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

/*
Bài tập worker pool: tạo bằng tay file dưới. file.txt sau đó đọc từng dòng file này nạp dữ liệu vào
 1 buffer channel có size 10, Điều kiện đọc file từng dòng.
Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ xong
nâng cao. In ra số lượng goroutine đã khởi tạo. hoàn thiện để tối ưu, thu hồi channel và goroutine đã cấp phát.

Nâng cao 1: Tạo 1 struct Line có trường gồm có: số dòng hiện tại, giá trị của dòng đó.
In ra màn hình cú pháp ${line_number} giá trị là: ${data}.
Nâng cao 2: Khi kết thúc chương trình đã cho đóng những vòng lặp vô hạn của các goroutine lại.
Viết chương trình đó. Giợi ý sử dụng biến make([]chan bool, n)
*/
type DataLine struct {
	content   string
	indentity int
}

func b4() {
	buffReadData := make(chan *DataLine, 10)
	defer close(buffReadData)
	var wg sync.WaitGroup

	f, _ := os.Open("file.txt")
	defer f.Close()

	for i := 1; i <= 3; i++ {
		go printData(buffReadData, &wg)
	}

	scanner := bufio.NewScanner(f)

	count := 1

	for scanner.Scan() {
		dataLine := &DataLine{content: scanner.Text(), indentity: count}
		count++
		buffReadData <- dataLine
		wg.Add(1)
	}

	wg.Wait()

}

func printData(jobs chan *DataLine, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-jobs:
			log.Printf("Hang %v : %v xong!\n", data.indentity, data.content)
			wg.Done()
		}
	}

}
