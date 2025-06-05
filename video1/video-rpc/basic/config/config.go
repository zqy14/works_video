package config

type Config struct {
	System System `mapstructure:"system" json:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql"`
}

type Mysql struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     uint64 `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
}

type System struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     uint64 `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
}
