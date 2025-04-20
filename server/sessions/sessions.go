package sessions

import (
	"net"
	"sync"
	"webserver/src/database"
	"webserver/src/server/terminal"
)

var Sessions = make(map[int64]*Session)
var SessionMutex sync.Mutex

type Session struct {
	ID   int64
	User *database.User
	Conn net.Conn
	Term *terminal.Term
	Chat bool
}

func Count() int {
	return len(Sessions)
}
