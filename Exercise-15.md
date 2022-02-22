Trong trường hợp tạo cột đếm thì làm sao để update cột đó? Làm sao để API chính không bị block vì phải update số đếm?

- Sử dụng một API trong API chính để gọi hàm Update với SQL Expression để update cột.
- Để API chính không bị block vì phải update số đếm ta đặt hàm gọi update cột vào một goroutine riêng và dùng recover để tránh crash chương trình.
