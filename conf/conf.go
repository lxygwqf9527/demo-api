package conf

// 全局config实例对象
var config *Config

func C() *Config {
	if config == nil {
		panic("Load Config first")
	}

	return config
}

// 初始化一个有默认值的config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "123456",
		Database:    "demo",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
		MaxLifeTime: 60 * 60,
		MaxIdleTime: 60 * 60,
	}
}

// MySQL todo
type MySQL struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" env:"MYSQL_PORT"`
	UserName    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	// lock        sync.Mutex
}
