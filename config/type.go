package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Jwt      JwtConfig
}

type AppConfig struct {
	Port string `mapstructure:"APP_PORT"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

type RedisConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	Port     string `mapstructure:"REDIS_PORT"`
}

type JwtConfig struct {
	Secret string `mapstructure:"JWT_SECRET_KEY"`
}
