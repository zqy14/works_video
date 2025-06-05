package config

type Config struct {
	System System `mapstructure:"system" json:"system"`
}

type NaCos struct {
	NamespaceId string `mapstructure:"namespaceId" json:"namespaceId"`
	IpAddr      string `mapstructure:"ipAddr" json:"ipAddr"`
	Port        uint64 `mapstructure:"port" json:"port"`
	DataId      string `mapstructure:"dataId" json:"dataId"`
	Group       string `mapstructure:"group" json:"group"`
}

type System struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     uint64 `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
}
