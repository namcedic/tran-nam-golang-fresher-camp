Khi nào cần tạo các cột số đếm ngay trên table dữ liệu (VD: liked_count trên restaurants)?

- Khi API lấy dữ liệu của bảng chịu tải read cao và thường xuyên, khi đó nếu cột số đếm mà phải tổng hợp từ bảng khác thì tốn thêm thời gian, cũng như database phải tính toán tổng hợp lại số liệu làm cho hệ thống DB phải chịu thêm tải.
- Ví dụ cột liked_count nếu được lưu trên bảng restaurants, do bảng restaurants là bảng thường xuyên chịu tải đọc cao, việc lưu cột liked_count sẽ khiến API lấy được dữ liệu dễ dàng mà ko phải gọi thêm API để tính toán và tổng hợp từ bảng restaurant_likes.
