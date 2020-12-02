package server

type ServerList struct {
	Grpc Server // main gRPC config for this service
	App  App
}

type Server struct {
	TLS     bool   `yaml:"tls"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

type App struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}
