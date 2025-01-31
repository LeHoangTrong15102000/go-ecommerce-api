package main

import "go-ecommerce-backend-api/internal/initialize"

// File main này bây giờ chúng ta sẽ không làm cái gì ở trong đây nữa
func main() {
	// r := routers.NewRouter()
	// Init MySQL
	// Init Redis
	// Init Kafka
	// Các phương thức init này nó phải được nhóm - Group lại vào trong một cái func khác ở trong một cái package khác, còn thằng main chỉ là nơi chỉ chạy phương thức Run() mà thôi

	// r.Run()

	initialize.Run()
}
