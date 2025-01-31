package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
)

// Nó có nhiệm vụ Run trước tất cả các file config
func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
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