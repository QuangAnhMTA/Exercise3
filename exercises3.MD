## về goroutine, channel và locking

- Tinh túy của golang nằm ở khả năng tận dụng tối đa phần cứng và các lõi cpu và các luồng. Sử dụng đa luồng trong golang là bài quan trọng và là sự bắt buộc mà bất cứ cá nhân nào muốn học phải nắm vững.

- Riêng bài này, là bài tối quan trọng. Nếu chưa nắm vững đồng nghĩa việc đi tiếp là không có ý nghĩa. Nên nhớ đây là cốt lõi của ngôn ngữ và việc nắm bắt và sử dụng thành thạo là bắt buộc.

- Nội dung được biên soạn bởi [maTe](fb.com/0.anhsang.0)

1. dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message.

```go
func chanRoutine() {
    log.Print("hello 1")
    go func() {
        time.Sleep(1 * time.Second)
        log.Print("hello 3")
    }
    log.Print("hello 2")
}

```

-- nâng cao. In ra các message theo thứ tự.
-- In ra message 3 trước message 2.
Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)

3. tạo 1 biến X `map[string]string` và 3 goroutine cùng thêm dữ liệu vào X. Mỗi goroutine thêm 1000 key khác nhau. Sao cho quá trình đủ 15 key không mất mát dữ liệu.

Lưu ý sử dụng mutex


4. chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
- Lý giải nguyên nhân lỗi.

```go
func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
			}
		}()
	}

	log.Print("done")
}

```

-- nâng cao. In ra các message theo thứ tự.
-- In ra message 3 trước message 2.
Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)


4. bài tập worker pool: tạo bằng tay file dưới. `file.txt` sau đó đọc từng dòng file này nạp dữ liệu vào 1 buffer channel có size 10, Điều kiện đọc file từng dòng. Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ `xong`
```txt
"z9nnHLy8V8"
"6AVcSrDUkB"
"DezRGPwtx7"
"eSmXGjCmTq"
"9rfCMntQA5"
"Trk6xppMuM"
"2sb8BPaUsp"
"6AAh6zVFNA"
"gsY8kAuKp8"
"FQgb8QEpxg"
"hEXnKUkYrp"
"tchiG2Tiv4"
"daMPMJWaM6"
"WbBMpX89Sz"
"YVnsveajtj"
"L9TA7FE5d9"
"xBjE7UNe98"
"q6bLPeVjYr"
"oBppTK62nT"
"GxUjEDYBdG"
"ZTEpXFStLo"
"4XkynbWFvU"
"WFmmUSWzDv"
"nit8qjmvZH"
"iT8BqzHdXo"
"7N7mz3qzn2"
"KfhMZsHABi"
"M4mKWrGgDn"
"qLEduDF7so"
"YhigrGfLJr"
"f82gk2mrxv"
"q7TPNZB3Bv"
"eWLL5Yg6sG"
"GyPqxrXiUg"
"86dGJYRzPN"
"EWYtAVfXnd"
"8dNcD3F8uS"
"NLRE6LKqCt"
"UbLD2DACiB"
"JeLHTTg8vw"
```
nâng cao. In ra số lượng goroutine đã khởi tạo.
hoàn thiện để tối ưu, thu hồi channel và goroutine đã cấp phát.

- Nâng cao 1: Tạo 1 struct `Line` có trường gồm có: `số dòng hiện tại`, `giá trị` của dòng đó.
In ra màn hình cú pháp `${line_number} giá trị là: ${data}`.
- Nâng cao 2: Khi kết thúc chương trình đã cho đóng những vòng lặp vô hạn của các goroutine lại. Viết chương trình đó.
Giợi ý sử dụng biến `make([]chan bool, n)`
