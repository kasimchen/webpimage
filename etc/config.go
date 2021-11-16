package etc


var ConfigVar *Config
type Config struct {

	Mysql Mysql
	Redis Redis
	RabbitMq RabbitMq


}

type Mysql struct {

	DB string `yaml:"db"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Redis struct {
	Host string `yaml:"host"`
	Password string `yaml:"password"`
	Port string `yaml:"port"`
	MaxIdle int `yaml:"maxidle"`
	MaxActive int `yaml:"maxactive"`
}

type RabbitMq struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}