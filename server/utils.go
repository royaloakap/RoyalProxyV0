package server

import (
	"fmt"
	"net"
)

func Print(conn net.Conn, str ...interface{}) {
	fmt.Fprint(conn, fmt.Sprint(str...))
}

func GetRank(rank int) string {
	var ret string
	switch rank {
	case 0:
		ret = "banned"
		break
	case 1:
		ret = "user"
		break
	case 2:
		ret = "Admin"
		break
	}
	return ret
}
