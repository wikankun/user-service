package database

func InitDB(config Config) {
	connectionString := GetConnectionString(config)
	err := Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
