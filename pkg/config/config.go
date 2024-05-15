package config

type Config struct {
	Env            string `mapstructure:"ENV"`
	Debug          string `mapstructure:"DEBUG"`
	Host           string `mapstructure:"HOST"`
	PgDbmsName     string `mapstructure:"PG_DBMS_NAME"`
	PgDriverName   string `mapstructure:"PG_DRIVER_NAME"`
	PgHost         string `mapstructure:"PG_HOST"`
	PgPort         string `mapstructure:"PG_PORT"`
	PgUsername     string `mapstructure:"PG_USERNAME"`
	PgPassword     string `mapstructure:"PG_PASSWORD"`
	PgDatabaseName string `mapstructure:"PG_DATABASE_NAME"`
	PgSSLMode      string `mapstructure:"PG_SSL_MODE"`
}
