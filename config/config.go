package config

type Config struct {
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Jwt    Jwt    `mapstructure:"jwt"`
	Log    Log    `mapstructure:"log"`
}

type App struct {
	DBType          string `mapstructure:"dbType"`
	RuntimeRootPath string `mapstructure:"runtimeRootPath"`
	ImageSavePath   string `mapstructure:"imageSavePath"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Redis struct {
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	MaxIdle     string `mapstructure:"maxIdle"`
	MaxActive   string `mapstructure:"maxActive"`
	IdleTimeout string `mapstructure:"idleTimeout"`
}

type Jwt struct {
	SecretKey string `mapstructure:"secretKey"`
}

type Log struct {
	LogSavePath string `mapstructure:"logSavePath"`
	LogSaveName string `mapstructure:"logSaveName"`
	LogFileExt  string `mapstructure:"logFileExt"`
	TimeFormat  string `mapstructure:"timeFormat"`
}
