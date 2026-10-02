package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App      AppConfig      `envPrefix:"APP_"`
	CORS     CORSConfig     `envPrefix:"CORS_"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	Redis    RedisConfig    `envPrefix:"REDIS_"`
	Logger   LoggerConfig   `envPrefix:"LOGGER_"`
	JWT      JWTConfig      `envPrefix:"JWT_"`
	Cookie   CookieConfig   `envPrefix:"COOKIE_"`
}

type AppConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"8080"`
	Mode string `env:"MODE" envDefault:"dev"` // "dev" or "prod"
}

type CORSConfig struct {
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:";" envDefault:"*"`
}

type PostgresConfig struct {
	Host     string `env:"HOST,required,notEmpty"`
	Port     int    `env:"PORT,required,notEmpty"`
	User     string `env:"USER,required,notEmpty"`
	Password string `env:"PASSWORD,required,notEmpty"`
	DB       string `env:"DB,required,notEmpty"`
}

type RedisConfig struct {
	Host     string `env:"HOST,required,notEmpty"`
	Port     int    `env:"PORT,required,notEmpty"`
	Password string `env:"PASSWORD,required,notEmpty"`
	DB       int    `env:"DB,required,notEmpty"`
}

type LoggerConfig struct {
	Level      string `env:"LEVEL" envDefault:"Debug"`
	SkipFrames int    `env:"SKIP_FRAMES" envDefault:"0"`
	Format     string `env:"FORMAT" envDefault:"json"` // can be "json" or "console"
	Out        string `env:"OUT" envDefault:"std"`     // can be "std" or "file"; if the "file" field is selected, the file name is logs.txt.
}

type JWTConfig struct {
	Secret        string        `env:"SECRET,required,notEmpty"`
	AccessExpiry  time.Duration `env:"ACCESS_EXPIRY" envDefault:"15m"`
	RefreshExpiry time.Duration `env:"REFRESH_EXPIRY" envDefault:"720h"`
}

type CookieConfig struct {
	Domain   string `env:"DOMAIN"`
	Secure   bool   `env:"SECURE" envDefault:"False"`  // true for HTTPS
	SameSite string `env:"SAME_SITE" envDefault:"Lax"` // "Strict", "Lax", or "None"
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
