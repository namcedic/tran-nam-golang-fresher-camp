Vì sao hệ thống lại cần pubsub và queue?

- Các module giao tiếp với nhau sẽ gọi chồng chéo lên nhau, Pubsub và queue được tạo ra để tránh việc này.
- Pubsub và queue là nơi trung gian để các module giao tiếp với nhau.
- Việc sử dụng Pubsub và queue làm cho API chính không bị ảnh hưởng do việc gửi message được xử lý ở một luồng khác không phải luồng chính, và message sẽ được gửi đến khi nào subcriber hoặc consumer nhận mới thôi.
