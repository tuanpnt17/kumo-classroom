package settings

type AppConfig struct {
	Server    ServerSetting    `mapstructure:"server"`
	Postgres  DatabaseSetting  `mapstructure:"postgres"`
	MySql     DatabaseSetting  `mapstructure:"mysql"`
	ZapLogger ZapLoggerSetting `mapstructure:"zap"`
}
