# tran-nam-golang-fresher-camp
Khoá chính (PK) trong table DB có công dụng gì?


- Khóa chính (PK) được sử dụng để định danh duy nhất mỗi dòng(record) tại bảng của cơ sở dữ liệu.
- Không được phép có dòng dữ liệu nào khác nhau mà có khóa chính trùng lặp nhau.
- Khóa chính nằm trước sẽ được ưu tiên sắp xếp trước theo vị trí vật lý của dữ liệu, giúp tăng tốc truy vấn dữ liệu.
- Ngoài ra, nó còn sử dụng để thiết lập quan hệ 1-n (hay ràng buộc tham chiếu) giữa hai bảng tại cơ sở dữ liệu.

Vì sao dùng ID là số tự tăng?
- ID tự động tăng cho phép một số duy nhất được tạo tự động khi một dữ liệu mới được chèn vào bảng.
- Không cần kiểm tra khóa sự tồn tại của khóa trong bảng dữ liệu, giảm thiểu sự phức tạp.

Khi nào một table dùng khoá chính trên nhiều cột?
- Khi một bảng không mang ý nghĩa là một thực thể, mục đích tồn tại để tạo liên kết giữa hai bảng với nhau.
- Khi kết hợp từ hai hay nhiều cột trên bảng dữ liệu sẽ được một khóa không trùng lặp nhau về dữ liệu.
