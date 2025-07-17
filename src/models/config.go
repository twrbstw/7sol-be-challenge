package models

type Configs struct {
	DbConfig     DbConfig
	LoggerConfig LoggerConfig
	AppConfig    AppConfig
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

type AppConfig struct {
	TokenTimeout string //in minute(s)
	SecretKey    string
}
