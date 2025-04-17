package settings

type ZapLoggerSetting struct {
	Level      string `mapstructure:"level"`
	LogPath    string `mapstructure:"logPath"`
	MaxSize    int    `mapstructure:"maxSize"`    // megabytes
	MaxBackups int    `mapstructure:"maxBackups"` // days
	MaxAge     int    `mapstructure:"maxAge"`     // days
	Compress   bool   `mapstructure:"compress"`   // disabled by default
	LocalTime  bool   `mapstructure:"localTime"`  // disabled by default
}
