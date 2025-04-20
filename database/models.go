package database

type User struct {
	Id       int
	Username string
	ApiKey   string
	Conns    int
	MaxTime  int
	Rank     int
	Expire   int
}

type APIS struct {
	Id     int
	Link   string
	Tag    string
	Active int
}

type Mirais struct {
	Id       int
	Host     string
	Port     int
	Username string
	Password string
	Tag      string
}

type Qbots struct {
	Id       int
	Host     string
	Port     int
	Username string
	Password string
	Tag      string
}

type SSH_SERVERS struct {
	Id       int
	Host     string
	Port     int
	Username string
	Password string
	Tag      string
}

type Attacks struct {
	Id       int
	Username string
	Host     string
	Port     int
	Duration int
	Method   string
	TimeSent int64
}

type Methods struct {
	Id      int
	Methods string
	Active  int
}
