# tran-nam-golang-fresher-camp

Vì sao trong khoá học này các bạn được khuyên không nên dùng khoá ngoại (FK), điểm yếu của khoá ngoại là gì?

- Nhiều service lớn phân tách ra thành nhiều service nhỏ mỗi service nắm giữ những phần database riêng biệt, giao tiếp với nhau qua API nên dùng FK không khả thi.
- Khi dữ liệu lớn lên sẽ tác động vào việc insert, edit, delete các dữ liệu trong DB, FK làm cho dữ liệu nhất quán nhưng cũng ràng buộc đến các bảng dữ liệu khác. Trước khi thao tác với dữ liệu sẽ phải kiểm dữ liệu các bảng con tham chiếu đến nó dẫn đến làm chậm tốc độ đọc ghi dữ liệu.
- Gây lãng phí tài nguyên đối với những DB không yêu cầu quá cao về tính toàn vẹn dữ liệu
- Khóa ngoại chỉ nên dùng ở những DB yêu cầu sự chặt chẽ và nhất quán về dữ liệu như các hệ thống ngân hàng, thanh toán tài chính...
