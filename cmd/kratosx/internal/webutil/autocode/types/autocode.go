package types

type ConnectDatabaseRequest struct {
	Drive    string `json:"drive"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbName"`
}
