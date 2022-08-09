package main

import "flag"

func main() {
}

func newConfigWithFlag() Config {
	dbHost := flag.String("host", "localhost", "Database Host Server")
	dbPort := flag.String("port", "5432", "Database Server Port")
	dbUser := flag.String("user", "postgres", "Database User Name")
	dbPassword := flag.String("password", "", "Database User Password")
	dbName := flag.String("db", "enigma", "Database Name")

	flag.Parse()
	return config
}

func newConfigWithFlagShortLong() Config {
	var dbHost string
	var dbPort string
	var dbName string
	var dbUser string
	var dbPassword string

	flag.StringVar(&dbHost, "host", "localhost", "Database Host Server")
	flag.StringVar(&dbPort, "port", "5432", "Database Server Port")
	flag.StringVar(&dbUser, "user", "postgres", "Database User Name")
	flag.StringVar(&dbPassword, "password", "", "Database User Password")
	flag.StringVar(&dbName, "db", "enigma", "Database Name")
}
