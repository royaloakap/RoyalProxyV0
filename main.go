package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goProxy/core/config"
	"goProxy/core/pnc"
	"goProxy/core/proxy"
	"goProxy/core/server"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var Fingerprint string = "S3LF_BU1LD_0R_M0D1F13D" // 455b9300-0a6f-48f1-82ee-bb1f6cf43500

type LicenseResponse struct {
	StatusMsg      string `json:"status_msg"`
	StatusOverview string `json:"status_overview"`
	StatusCode     int    `json:"status_code"`
	StatusID       string `json:"status_id"`
	DiscordID      string `json:"discord_id"`
}

type Config struct {
	License string `json:"license"`
}

func main() {
	// Initialisation du proxy
	proxy.Fingerprint = Fingerprint

	// Ouverture du fichier de journalisation
	logFile, err := os.OpenFile("crash.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Initialisation des gestionnaires de panique et de signal
	pnc.InitHndl()
	defer pnc.PanicHndl()

	// Désactivation du journal d'erreurs (à utiliser si nécessaire)
	// log.SetOutput(logFile)

	// Début du programme
	fmt.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46mStarting Royal Proxy ...\u001B[38;5;230m")

	// Chargement de la configuration
	config.Load()
	fmt.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46m Loaded Config ...\u001B[38;5;230m")

	// Chargement de la clé de licence depuis config.json
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error reading configuration file\u001B[38;5;230m: %v", err)
	}

	var conf Config
	err = json.Unmarshal(configData, &conf)
	if err != nil {
		log.Fatalf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m  Error parsing configuration file\u001B[38;5;230m: %v", err)
	}

	licenseKey := conf.License
	if licenseKey == "" {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m  A license key is missing in the configuration file\u001B[38;5;230m.")
		os.Exit(1)
	}

	// Paramètres pour la requête HTTP
	product := "RoyalProxylite"
	apiKey := "LicenseKey"
	apiUrl := "https://your-api.net/api/client"

	requestData := map[string]string{
		"licensekey": licenseKey,
		"product":    product,
	}

	// Envoi de la requête HTTP
	requestDataJSON, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE] Error marshaling request data: %v", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestDataJSON))
	if err != nil {
		log.Fatalf("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE] Error creating HTTP request: %v", err)
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error sending HTTPS request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error reading response body:", err)
		os.Exit(1)
	}

	var licenseData LicenseResponse
	err = json.Unmarshal(body, &licenseData)
	if err != nil {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Error parsing response body:", err)
		os.Exit(1)
	}

	if licenseData.StatusOverview == "success" {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46m Royal CNC FREE VERSION is started and was created by \u001B[38;5;122m ~ Royaloakap ~ \u001B[38;5;46m Your licence\u001B[38;5;230m", licenseKey, "\u001B[38;5;46m is valid.\u001B[38;5;230m")
	} else {
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;196m Your licence\u001B[38;5;230m", licenseKey, "\u001B[38;5;196m is invalid or has reached a ceiling. Contact me on \u001B[38;5;122m discord.gg/RoyalC2 \u001B[38;5;230m or \u001B[38;5;122m @royaloakap.\u001B[38;5;230m")
		log.Println("\u001B[0m\u001B[107m\u001B[38;5;163m[LICENSE]\u001B[0m\u001B[38;5;46m Open a ticket on Discord.gg/RoyalC2\u001B[38;5;230m")
		os.Exit(1)
	}

	// Initialisation du serveur
	fmt.Println("Initialising ...")
	go server.Monitor()
	for !proxy.Initialised {
		time.Sleep(500 * time.Millisecond)
	}

	go server.Serve()

	// Maintien du serveur en cours d'exécution
	select {}
}
