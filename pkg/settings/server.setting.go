package settings

type ServerSetting struct {
	Port        int    `mapstructure:"port"`
	Environment string `mapstructure:"environment"`
}
