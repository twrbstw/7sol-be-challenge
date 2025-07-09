package models

type Configs struct {
	DbConfig DbConfig
}

type DbConfig struct {
	Uri  string
	Name string
}
