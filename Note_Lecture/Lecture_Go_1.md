# Khoá học Golang restful API

> > > > Sẽ lựa chọn GIN đi theo xuyên suốt cái dự án của chúng ta

>

- Ví dụ như POST thì người ta sẽ có là c.ShouldBindJSON() để mà thao tác với cái vấn đề đó thôi

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

  - Còn đi sâu hơn thì chúng ta sẽ có nhiều cái response, ví dụ như là trả về một cái list nó có limit, skip, offset, ... nó có phân trang thì nó sẽ khác -=> mà trả về cái đối tượng tương tự như vậy thì chúng ta có nhiều cái cách trả về của các phương thức khác nhau nữa nên là khi đi sâu vào dự án chúng ta sẽ làm dần dần.

  - Thì chúng ta sẽ có 2 cái trạng thái một là `ErrorResponse` 2 là `SuccessResponse`, thì là còn nhiều lỗi lắm giống như là chúng ta làm từng bước thì nó sẽ có nhiều cái event xảy ra -> Giống như là token bị lỗi , token hết hạn, refresh-token hết hạn,... Thì lúc đó chúng ta cần phải khai báo nhiều cái ở trong `msg` để mà thống nhất

> > > > Go (6) GIN vs Logs Handler

- Một BE mà không có log là một sai lầm, ghi log là một phần kh ông thể thiếu ở trong môi trường lập trình ứng dụng

- Thì ở trong Golang này chúng ta sẽ sử dụng thằng `Zap` một trong những mã nguồn mở của `Uber` ghi log rất là mạnh mẽ -> Cung cấp ghi log phân cấp có cấu trúc cực kì nhanh và phân bổ tài nguyên tối thiểu -> Thì cho thấy thằng `Zap` có hiệu suất vượt trội hầu hết hơn hẳn các thằng khác so với các thư viện ghi Log có cấu trúc tương đương ở trong golang

- Thằng Zap nó cung cấp 2 loại log đó chính là `Sugar` và `Logger`, và thằng `Infof` thì nó là dạng giống như là `Printf` ở trong thằng golang này vậy

  - Thì thằng Logger nó cung cấp kiểu `Key` và `Value`

  - Thì thằng Zap này nó cung cấp cho chúng ta 3 cái cấu hình để mà chúng ta lựa chọn ghi log trong nhật kí -> Đó là Example, nhưng mà thật ra thì nó 2 cái cấu hình rất là cơ bản đó là `development` và `production`

  - Chúng ta sẽ có 3 cái level khác nhau, 3 cái cấp độ nhật kí khác nhau

    - Thằng Development thì nó đưa ra INFO log ra thời gian của nó, còn thằng Production thì nó tự customize lại cái log của nó -> Cho nên chúng ta sẽ làm một cái riêng cho dự án `Go Ecommerce` BE của mình và tích hợp nó vào dự án để mà sử dụng xuyên suốt trong cả quá trình phát triển luôn -> Sẽ áp dụng vào `NewProduction` và `NewDevelopment` làm sao để mà cho nó có hiệu quả

  - Việc thực hiện `customize` thì nó sẽ như thế nào

    - Đầu tiên thì chúng ta sẽ viết định dạng của log, zapcore là một tính năng đóng gói của zap

    - Bây giờ chúng ta muốn định dạng về cái thời gian của chúng ta đi, thời gian và ngày tháng ở trong production thì nó đang là `timestamp` bây giờ chúng ta muốn định dạng nó về định dạng thời gian mà chúng ta hay thấy nhất `2025-01-29 10:00:00` thì nó sẽ như thế nào

      - Thi đầu tiên cần phải khai báo thêm một cái `encodeConfig.EncodeTime`

      - Về thằng write thì cái thằng zap nguyên mẫu thì nó không cho phép chúng ta phân đoạn từng file một -> Do nó phải sử dụng một cái gói mà thằng Uber nó khuyến nghị đó là `Lumberjack`

      - 3 cái tham số này khi mà mở một file trong nodejs hay golang đều phải có 3 cái tham số này `file, _ := os.OpenFile(name, flag, perm)` -> Đó là `name, flag, perm`

        - Thứ nhất là tập tin - tên của file
        - THứ hai là flag(Đây là kiểu số nguyên INT) - giúp cho chúng ta đại diện cho chế độ mở tập tin, ví dụ như read-only, write-only, read-write, ... hoặc là append, ...
        - Thứ ba là perm - độ truy cập của file, ví dụ như là 066, hay 0755 là quyền đọc và ghi thực thi cho `Owner`

      - Ví dụ như đại diện đọc cho một cái file ở bên ngoài thì chúng ta sẽ làm như thế nào

  - Do là demo nên là cần phải viết trong một cái file main, còn khi mà viết thực thì cần phải viết vào bên trong cái folder `logger` ở bên trong `pkg` của dự án thì nó mới đúng được

  - Và cái nữa là cái log.txt ở trong thằng `zap` thì nó không hỗ trợ việc phân đoạn ở trong nhật ký -> Nên nếu chúng ta muốn là 30MB nó tự xoá hoặc là maximum mấy ngày thì nó tự xoá, nâng cao như vậy thì chúng ta nên sử dụng package nào cho nó đúng -> Thì Lumberjack tích hợp như thế nào thì video sau sẽ rõ

> > > > Go (7) GIN vs Viper Config

- Việc quản lí cấu hình là nhiệm cơ bản nhưng mà nó rất rất là quan trọng, nó liên quan đến việc quản lí hiệu quả cái biến config ở bên trong cái cấu hình của dự án của mình, trong cái application của mình, chẳng hạn như là thông tin kết nối database, các cái API của third-party, ... -> Nên là chúng ta cần phải có một cái công cụ và một cái thư viện quản lí tốt giúp cho các lập trình viên quản lí các cấu hình này dễ dàng hơn, ngoài ra còn cải thiện tính bảo mật và tính linh hoạt của dự án

- Nếu bây giờ lỡ có nhiều `database` ở trong dự án thì cần phải làm như thế nào,

  - Nếu như mà chúng ta thực hiện giống như trong cái file local.yaml thì cái file config của chúng tôi nó chứa cái đối tượng rất là phức tạp -> Cho nên chúng ta ưu tiên sử dụng UnMarshal để mà thực hiện việc liên kết các config trực tiếp và đưa vào một cái cấu trức `slide`

  - Đầu tiên chúng ta cần xác định cái `file` ở trong cấu hình go phù hợp với cấu trúc của một cái `struct`

  - Thì Database cũng là một cái `slice struct`

- Bình thường thì cũng có sử dụng struct chứ không phải là không, nhưng mà khi đã ra dự án thực tế rồi thì nó quản lí nhiều nguồn tham chiếu nên cần phải có struct để mà quản lí(Thì khi đó struct nó mới phát huy tối đa hiệu quả được)

> > > > Go (8) GIN vs Middleware

- Middleware thì cần phải return về `handlerFunc`

- Có 2 cách để mà chúng ta viết middleware đó là:

  - Cách 1 là chúng ta sẽ return về `handlerFunc`

  - Cách 2 là chúng ta sẽ trỏ cái `*gin.Context` luôn là `gin.HandlerFunc`

- Đã có các midđleware rồi thì bây giờ chúng ta sẽ khai báo nó như thế nào

  - Khi mà chạy thì nó sẽ in cái `Before` trước, sau đó sẽ tới `MyHandler`, rồi cuối cùng thì nó sẽ trả về cái `Alter`

  - 5 cái middleware chúng ta thường hay sử dụng trong cái ứng dụng của chúng ta đó là: auth, cors, logger, ratelimit, errorhandler

  - Thì chúng ta sẽ làm thử cái middleware của `Auth` khi mà người dùng đăng nhập thì cần phải check Authentication của người dùng

> > > > Go (9) GIN vs Unit test Go

- Sẽ đi qua các cái ví dụ thực tế để mà hiểu rõ hơn gói testing ở trong `Golang`, -> Sẽ h iểu được cách viết test đơn giản, hiểu được cái luồng đi, cách kiểm tra một cái lỗi, cách sử dụng các công cụ(tool) hỗ trợ như thế nào

- Về viết test thì cần phải đảm bảo chất lượng và độ tin cậy của ứng dụng này

- `Test-Driven` là một phương pháp phát triển trong đó chúng ta sẽ viết test trước khi mà viết code -> Có thể là chúng ta làm xong và đưa cái file test đó cho `tester` để mà họ -> Để mà có thể giúp chúng ta rút ngắn cái thời gian xảy ra lỗi nhiều hơn -> Thì cách thực hiện và sử dụng nó như thế nào thì chúng ta sẽ đi tìm hiểu nó.

- Với những cái test-case cần phải biết được đầu vào và đầu ra của nó - input và output như thế nào thì mới có thể viết được

- Cái func ở trong file test thì đầu tiên là nó phải có chữ test thì mới được -> Và ở trong đó thì chúng ta cần phải add cái gói `testing` của go vào

  - Chúng ta cần phải khai báo 2 cái biến:

    - 1 là input đầu vào là gì
    - 2 là output đầu ra là gì

  - Và thực tế là sẽ nhận được cái gì

- Ngoài ra các lập trình viên còn có gói `testify` để mà viết test ở trong `golang`

- Cái hàm Require nó sẽ khác so với thằng `Assert` khi mà nó test mà nó bị failed thì các thằng đằng sau nó sẽ bị dừng lại, còn thằng Assert khi mà nó failed thì nó vẫn tiếp tục thực hiện phương thức ở phía dưới

  - Require nó sẽ chặn cái câu lệnh ở phía dưới nếu mà chạy câu test không đúng

- Thì ở trong golang thì cái testing cơ bản không thể nào mà dùng tại đây được -> Trong golang nó hỗ trợ cho chúng ta một cái tool gọi là `Coverage` -> Thì trong lập trình ngta gọi là `Code Cover` -> Thì nó là một cái thước đo rất là quan trọng để mà xác định tỉ lệ phần trăm cái `source code` của chúng ta nó đã được kiểm tra bởi cái `testing` hay chưa

  - Có thể tạo một cái rp để mà kiểm tra độ bao phủ code của chúng ta là bao nhiêu ở trong `Golang` để mà đọc quả HTML

  - Chạy một cái câu lệnh để mà chạy tất cả các test và sau đó lưu kết quả vào cái file HTML thì chỉ cần sử dụng tool là `go tool cover -html=coverage.out`

- Qua video sau chúng ta sẽ thực hiện sắp xếp lại thư mục, loại lại cái config, kết nối database, kết nối redis, và làm cái gì đó đầy đủ trong một file -> Bắt đầu chinh chiến Golang từ golang từ level 2 - 5

> > > > Go (10) - Fix func main để tạo nhóm kết nối chung(Initialize)

- Video 5

> > > > Go (11) - Sủ dụng Lumberjack quản lí FILE LOG lớn

> > > > Go (12) - Connect MySQL Pool hiệu quả và test BenchMark cho 2 trường

- Video

> > > > Go (13) - Connect Redis Pool hiệu quả và test BenchMark cho 2 trường hợp

- Video

> > > > Go (14) - Triển khai Router cho TEAM Lớn | Cuộc chiến bắt đầu

- Video

> > > > Go (15) - Kafka thực hành về mua bán cổ phía với các tình huống

- Video

> > > > Go (16) - Interface vì sao các dự án lớn tôi luôn đề nghị triển khai | NOCODE

- Video

> > > > Go (17) - Interface cách triển khai và bố trí CODE | THỰC HÀNH

- Video

> > > > Go (18) - Wire Dependency Injection

- Video

> > > > Go (19) - Migrating Schema With GORM to MySQL and DUMP Database

- Video

> > > > Go (20) - Chiến đấu DOCKER, sai lầm cách build này level 0,1

- Video

> > > > Go (21) - Chiến đấu với DOCKER LINK, docker compose build project Level 2 ,3 ,4

- Video

> > > > Go (22) - Advanced: Công ty đề nghị chuyển GORM sang SQLC như thế nào?

- Video

> > > > Go (23) - Advanced: Goose(A Database Migration Tool) hiệu suất cao của dân Backend

- Video

> > > > Go (24) - User Register Send OTP To Email Template

- Video

> > > > Go (25) - Loại đại ca GORM thay thế tân binh GORM vs SQLC hệ thống trở nên mạnh mẽ

- Video

> > > > Go (26) - Làm việc với Microservices, GO handler OTP, Java send OTP to AWS

- Video

> > > > Go (27) - Triển khai Kafka Microservices

- Video
