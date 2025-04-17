package settings

type DatabaseSetting struct {
	User                     string `mapstructure:"user"`
	Password                 string `mapstructure:"password"`
	Host                     string `mapstructure:"host"`
	Port                     int    `mapstructure:"port"`
	DbName                   string `mapstructure:"dbName"`
	MaxIdleConns             int    `mapstructure:"maxIdleConns"`
	MaxOpenConns             int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetimeInSeconds int    `mapstructure:"connMaxLifetimeInSeconds"`
}
