package config

import "os"

// * global env
func GinMode() string {
	env := os.Getenv("GIN_MODE")
	if env == "" {
		panic("can't load env gin mode")
	}
	return env
}
func JwtSecreet() string {
	env := os.Getenv("JWT_SECRET")
	if env == "" {
		panic("can't load env jwt secreet")
	}
	return env
}
func AppPort() string {
	env := os.Getenv("APP_PORT")
	if env == "" {
		panic("can't load env app port")
	}
	return env
}

// * database env
func DbHost() string {
	env := os.Getenv("DB_HOST")
	if env == "" {
		panic("can't load env db host")
	}
	return env
}
func DbUsername() string {
	env := os.Getenv("DB_USER")
	if env == "" {
		panic("can't load env db username")
	}
	return env
}
func DbPassword() string {
	env := os.Getenv("DB_PASS")
	if env == "" {
		panic("can't load env db password")
	}
	return env
}
func DbName() string {
	env := os.Getenv("DB_NAME")
	if env == "" {
		panic("can't load env db name")
	}
	return env
}
func DbPort() string {
	env := os.Getenv("DB_PORT")
	if env == "" {
		panic("can't load env db port")
	}
	return env
}
