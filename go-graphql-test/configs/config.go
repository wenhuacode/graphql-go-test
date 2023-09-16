package configs

import "time"

// Config object
type Config struct {
	Env       string        `mapstructure:"env" json:"env"`
	Pepper    string        `mapstructure:"pepper" json:"pepper"`
	HMACKey   string        `mapstructure:"hmackey" json:"hmackey"`
	Mysql     MySQLConfig   `mapstructure:"mysql" json:"mysql"`
	Mailgun   MailgunConfig `mapstructure:"mailgun" json:"mailgun"`
	JWTSecret string        `mapstructure:"jwtsecret" json:"jwtsecret"`
	Host      string        `mapstructure:"host" json:"host"`
	Port      string        `mapstructure:"port" json:"port"`
	FromEmail string        `mapstructure:"fromemail" json:"fromemail"`
}

type MailgunConfig struct {
	APIKey string `mapstructure:"apikey" json:"apikey"`
	// PublicAPIKey string `env:"MAILGUN_PUBLIC_KEY"`
	Domain string `mapstructure:"domain" json:"domain"`
}

type MySQLConfig struct {
	Host                  string        `mapstructure:"host" json:"host"`
	Port                  int           `mapstructure:"port" json:"port"`
	Username              string        `mapstructure:"username" json:"username"`
	Password              string        `mapstructure:"password" json:"password"`
	Database              string        `mapstructure:"database" json:"database"`
	MaxIdleConnections    int           `mapstructure:"maxidleconnections" json:"maxidleconnections"`
	MaxOpenConnections    int           `mapstructure:"maxopenconnections" json:"maxopenconnections"`
	MaxConnectionLifetime time.Duration `mapstructure:"maxconnectionlifetime" json:"maxconnectionlifetime"`
	LogLevel              int           `mapstructure:"loglevel" json:"loglevel"`
}
