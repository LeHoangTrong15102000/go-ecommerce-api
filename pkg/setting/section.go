package setting

// Chúng ta tham chiếu cái global Config này để mà gọi ở những cái file tiếp theo
type Config struct{
	Mysql MySQLSetting `mapstructure:"mysql"` // Trỏ tới file local của chúng ta
}


// Chứa tất cả tổng thể của một cái setting, chúng ta tham chiếu cái global config này	
type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}