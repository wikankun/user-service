package database

func InitDB() {
	connectionString := GetConnectionString()
	err := Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
