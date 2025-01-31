package initialize

// Nó có nhiệm vụ Run trước tất cả các file config
func Run() {
	// Load configuration
	LoadConfig()
	// Khi mà loadconfig xong thì phải ghi log liền
  InitLogger()
	// Connect Mysql
	InitMysql()
	// Connect Redis
	InitRedis()
	// Connect Router
	r := InitRouter()

	r.Run(":8002")
}