# Khoá học Golang restful API

> > > > Sẽ lựa chọn GIN đi theo xuyên suốt cái dự án của chúng ta

>

- Ví dụ như POST thì người ta sẽ có là c.ShouldBindJSON() để mà thao tác với cái vấn đề đó thôi
  \

> > > > Go (4) GIN và MVC

- Sẽ thực hiện các thao tác với một controller ở trong , nên là đặt tên cho cái folder controller của mỗi tính năng là rất quan trọng ở trong golang -> Thì cái hàm đó được gọi là Free Func vì nó không thuộc về một đối tượng nào, không thuộc về một cái xuất nào mà nó dùng trực tiếp GIN context để mà trả về trực tiếp một cái `JSON response`

  - Cái khó khăn ở đây là việc mở rộng và bảo trì cho team khi mà phát triển dự án,

  - Nó không tận dụng được cái `struct` của thằng golang, thế thì `struct` là gì, thì chúng ta sẽ có một video nói về struct thật là sâu, sẽ nói sơ qua cái struct ở bên trong golang

  - Và cái thứ 2 sẽ là `NewUserController` để mà khởi tạo một cái đối tượng instance ở bên trong `user` (vì trong user có thể có nhiều service thì sao, nên là sau này chúng ta sẽ nói về cái vấn đề đó)

  - Thì khi mà viết như vậy thì cho phép chúng ta tổ chức mã nguồn thật là tốt, cái thứ 2 là dễ dàng trong việc quản lí trạng thái dữ liệu liên quan của người dùng thông qua cái trường hợp của `struct`(struct cực kì là mạnh mẽ) -> Những bài sau chúng ta sẽ đi sau về `struct` `con trỏ` và `địa chỉ con trỏ`, Ngoài ra còn có Go `Routine` và `Channel` là lập trình đồng bộ cho phép chúng ta lập trình nhiều `Routine` thì cái vấn đề đó chúng ta sẽ bàn về sau ở trong các chương kế tiếp

- Thì bây giờ chúng ta sẽ quay lại cái `Pong` để mà chúng ta cảm nhận được struct ở trong golang nó tuyệt vời như thế nào

- Khi mà nó chạy bình thường như vậy là không có cái vấn đề gì rồi nhưng mà chúng ta muốn `handle` cái `error` cho toàn bộ dự án thì như thế nào -> Cách handle error trả về thống nhất cho dữ liệu người dùng như thế nào thì video sau sẽ thực hiện cái vấn đề đó

> > > > Go (5) GIN vs Error Handler

- Những cái dự án gần đây thường viết một cái method thông qua một cái `struct`

- Nếu như có lỗi thì cần khai báo chung chung cho ng dùng mà thôi, không được cho biết là hệ thống đang bị lỗi gì nếu mà như vậy thì `hacker` nó sẽ dò ra được cái mật khẩu của người dùng.

-Thì thằng `response` nó sẽ sử dụng 2 phương thức đó là

- Cái phương thức đầu tiên đó chính là, bây giờ thì nó sẽ được thay thể ở bên trong -> Thì bây giờ chúng ta sẽ viết một cái `response` để mà thay thế cái thằng `gin.H` ở trong từng thằng `controller`

  - Bây giờ chúng ta cần phải định dạng thống nhất cho ae

  - `func SuccessResponse(c \*gin.Context, code int, data interface{}) {}`, gin.Context{} còn phải có ở trong mỗi response

> > > > Go (6) GIN vs Logs Handler

- Video 5

> > > > Go (7) GIN vs Viper Config

- Video 5

> > > > Go (8) GIN vs Middleware

- Video 5

> > > > Go (9) GIN vs Unit test Go

- Video 5

> > > > Go (9) GIN vs Unit test Go

- Video 5

> > > > Go (10) - Fix func main để tạo nhóm kết nối chung(Initialize)

- Video 5
