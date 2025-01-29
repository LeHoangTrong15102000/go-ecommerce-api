package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()

	// //  Sugar
	// sugar.Infof("Hello, World!")

	// // Logger
	// logger := zap.NewExample()
	// logger.Info("Hello, World!", zap.String("name", "Tips Golang"))


 // 2.
//  logger := zap.NewExample()
//  logger.Info("Hello wworld")


// 	// Development
//  logger, _ = zap.NewDevelopment()
//  logger.Info("Hello Development")

//  // Production
//  logger, _ = zap.NewProduction()
//  logger.Info("Hello Production")


 // 3. 
 encoder := getEncoderLog()
 sync := getLogWriterSync()
 core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

 logger := zap.New(core, zap.AddCaller())

 logger.Info("Info log", zap.Int("line", 1))
 logger.Error("Error log", zap.Int("line", 2))
}


// Format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig() // khởi tạo một cái giá trị mới

	// Nó sẽ biến đổi cái timestamp thành định dạng quen thuộc cho chúng ta
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Gán lại bằng giá trị khác
	// ts -> timestamp
	encodeConfig.TimeKey = "time"
	// info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Caller: "cli/main.log.go:24" -> cái caller này cho chúng ta biết được là cái file này nó được ghi ở file nào và dòng nào
	// Thì cái thằng này nó phải áp dụng ở cái điều kiện đó là
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // Thì cái này là ở file nào và mở file nào


	return zapcore.NewJSONEncoder(encodeConfig)
}


// Ghi logs vào vị trí nào thì cần phải khai báo thêm một hàm nữa
// Về thằng write thì cái thằng zap nguyên mẫu thì nó không cho phép chúng ta phân đoạn từng file một
func getLogWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file) // Thêm cái file vào 


	// Sẽ sử dụng một cái gói để mà đại diện `standardError` cho cái luồng xuất lỗi tiêu chuẩn -  Thì sẽ thêm một `Stderr error``
	syncConsole := zapcore.AddSync(os.Stderr) // Đây là luồng xuất lỗi tiêu chuẩn

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}