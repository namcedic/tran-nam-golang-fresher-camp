- Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud?
+ Nó sẽ ảnh hưởng tới việc tải các service chính khác vì upload là một service tiêu tốn cpu và băng thông khá nhiều, nên tách ra thành một service riêng, tránh ảnh hưởng đến service chính. 
+ Lưu file upload vào ngay bên trong service sẽ khiến service tốn dung lượng để lưu trữ, chứa các file vật lý.
+ Dùng Cloud giúp tiết kiệm chi phí cho các thiết bị lưu trữ vật lý.
+ Dùng Cloud sẽ tăng cường bảo mật, giảm nguy cơ mất dữ liệu, có thể truy cập mọi lúc mọi nơi với nhiều loại thiết bị khác nhau và dễ mở rộng.


-Vì sao không chứa ảnh binary vào DB?
+ DB sẽ yêu cầu nhiều dung lượng để lưu trữ hơn và sẽ làm ảnh hưởng đến việc back-up, truy xuất dữ liệu.
+ Lưu trữ binary data lớn trong DB sẽ ảnh hưởng đến hiệu năng của DB.
+ Nếu có lỗi trong binary data sẽ khó phát hiện.
+ Làm tăng tính bảo mật cho dữ liệu
