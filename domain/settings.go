package domain

type Settings struct {
	Http HTTPSettings
	DB   DBSettings
}

type HTTPSettings struct {
	Host string
	Port uint16
}

type DBSettings struct {
	Host     string
	Port     uint16
	Name     string
	User     string
	Password string
	SSL      string
}
