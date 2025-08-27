package types

type Config struct {
	Port             int    `env:"PORT" default:"8080"`
	DBHost           string `env:"DB_HOST" default:"localhost"`
	DBName           string `env:"DB_NAME" required:"true"`
	DBPort           string `env:"DB_PORT" required:"true"`
	DBUsername       string `env:"DB_USERNAME" required:"true"`
	DBPassword       string `env:"DB_PASSWORD" required:"true"`
	DBMinConnections int    `env:"DB_MIN_CONNECTIONS" default:"0"`
	DBMaxConnections int    `env:"DB_MAX_CONNECTIONS" default:"10"`
	DBLogger         bool   `env:"DB_LOGGER" default:"true"`
	ApiPrefix        string `env:"API_PREFIX" default:"api"`
	UseSQLConnector  bool   `env:"USE_SQL_CONNECTOR" default:"false"`
	Environment      string `env:"GO_ENV" default:"production"`
	JwtSecret        string `env:"JWT_SECRET" required:"true"`
	DBSSLMode        string `env:"DB_SSL_MODE" default:"disable"`
}
