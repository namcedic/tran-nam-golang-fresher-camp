# tran-nam-golang-fresher-camp

Index trong table DB có công dụng gì?

- Chỉ mục (INDEX) trong SQL là bảng tra cứu đặc biệt mà công cụ tìm kiếm cơ sở dữ liệu có thể sử dụng để tăng nhanh thời gian và hiệu suất truy xuất dữ liệu tương tự như mục lục của một cuốn sách.
- Index giúp tăng tốc đáng kể tốc độ đọc dữ liệu, trong khi làm chậm dữ liệu ghi vào ít hơn.
- Index tìm kiếm theo kiểu nhị phân nên tăng tốc độ tìm kiếm lên nhiều so với tìm kiếm không index;
- Không phải quét toàn bộ dữ liệu của bảng nếu dùng với index có khóa chính là cột đầu tiên trong bảng

Nếu không có index thì khác biệt trong truy vấn là gì?
- Truy vấn sẽ quét toàn bộ dữ liệu của bảng để tìm kiếm dữ liệu, trong khi với index sẽ tìm kiếm nhị phân và sau đó tìm tới khóa chính để tìm ra được dữ liệu cần thiết.
