package models

type Configs struct {
	DbConfig     DbConfig
	LoggerConfig LoggerConfig
}

type DbConfig struct {
	Uri  string
	Name string
}

type LoggerConfig struct {
	Format     string
	TimeFormat string
	TimeZone   string
}
