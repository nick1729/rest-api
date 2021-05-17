package types

// DbConfig contains DB config file struct
type DbConfig struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DbName string
}
