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
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Jwt struct {
	SecretKey  string `mapstructure:"secretKey"`
	Expiration int64 `mapstructure:"expiration"`
	Issuer     string `mapstructure:"issuer"`
}

type Log struct {
	LogSavePath string `mapstructure:"logSavePath"`
	LogSaveName string `mapstructure:"logSaveName"`
	LogFileExt  string `mapstructure:"logFileExt"`
	TimeFormat  string `mapstructure:"timeFormat"`
}
