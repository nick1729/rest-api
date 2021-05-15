package types

// Config file struct
type Config struct {
	Login string `json:"db_login"`
	Pass  string `json:"db_password"`
	Ip    string `json:"db_ip"`
	Port  int    `json:"db_port"`
	Table string `json:"table_name"`
}
