package db

type DatabaseList struct {
	Product_DB struct {
		MySql Database
	}
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Adapter  string `yaml:"adapter"`
}