package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"goProxy/core/domains"
	"goProxy/core/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func Generate() {

	fmt.Println("[ " + utils.PrimaryColor("Aucun fichier de configuration trouvé") + " ]")
	fmt.Println("[ " + utils.PrimaryColor("Configurer Royal proxy maintenant") + " ]")
	fmt.Println("")

	gConfig := domains.Configuration{
		Proxy: domains.Proxy{
			Cloudflare:  utils.AskBool("Utiliser ce proxy avec Cloudflare ? (y/N)", false),
			Network:     "tcp",
			AdminSecret: utils.RandomString(25),
			APISecret:   utils.RandomString(30),
			Timeout: domains.TimeoutSettings{
				Idle:       utils.AskInt("Combien de secondes une connexion indle doit-elle rester ouverte?", 3),
				Read:       utils.AskInt("Combien de secondes une connexion de lecture doit-elle rester ouverte?", 5),
				Write:      utils.AskInt("Combien de secondes une connexion d'écriture doit-elle rester ouverte?", 5),
				ReadHeader: utils.AskInt("Combien de secondes faut-il accorder pour lire un en-tête de connexion?", 5),
			},
			Secrets: map[string]string{
				"cookie":     utils.RandomString(20),
				"javascript": utils.RandomString(20),
				"captcha":    utils.RandomString(20),
			},
			Ratelimits: map[string]int{
				"requests":           utils.AskInt("Après combien de requêtes d’une IP dans les 2 minutes doit-elle être bloquée ?", 1000),
				"unknownFingerprint": utils.AskInt("Après combien de demandes d'une empreinte digitale inconnue dans les 2 minutes doit-elle être bloquée?", 150),
				"challengeFailures":  utils.AskInt("Après combien de tentatives infructueuses pour résoudre un défi provenant d'une adresse IP dans les 2 minutes, celle-ci devrait-elle être bloquée?", 40),
				"noRequestsSent":     utils.AskInt("Après combien de tentatives de connexion TCP sans envoyer de requête HTTP depuis une IP dans les 2 minutes doit-elle être bloquée?", 10),
			},
		},
		Domains: []domains.Domain{},
	}

	domains.Config = &gConfig

	jsonConfig, err := json.Marshal(gConfig)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("config.json", jsonConfig, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	AddDomain()
}

func AddDomain() {
	fmt.Println("[ " + utils.PrimaryColor("Aucune configuration de domaine trouvée") + " ]")
	fmt.Println("[ " + utils.PrimaryColor("Configurer de nouveaux domaines dans le Config.json") + " ]")
	fmt.Println("")
	gDomain := domains.Domain{
		Name:        utils.AskString("Quel est le nom de votre domaine (eg. \"example.com\")", "example.com"),
		Backend:     utils.AskString("Quel est la backend/serveur sur lequel le proxy devrait proxy ??", "1.1.1.1"),
		Scheme:      strings.ToLower(utils.AskString("Quel schéma le proxy doit-il utiliser pour communiquer avec votre backend ? (http/https)", "http")),
		Certificate: utils.AskString("Quel est le chemin d’accès au certificat SSL pour votre domaine ? (Laissez vide si vous utilisez le proxy derrière Cloudflare)", ""),
		Key:         utils.AskString("Quel est le chemin d’accès à la clé SSL pour votre domaine ? (Laissez vide si vous utilisez le proxy derrière Cloudflare)", ""),
		Webhook: domains.WebhookSettings{
			URL:            utils.AskString("Quelle est l’URL de votre webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			Name:           utils.AskString("Quel est le nom de votre webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			Avatar:         utils.AskString("Quelle est l’URL de votre avatar Webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			AttackStartMsg: utils.AskString("Quel est le message que votre webhook devrait envoyer lorsque votre site Web est attaqué?", ""),
			AttackStopMsg:  utils.AskString("Quel message votre webhook devrait-il envoyer lorsque votre site Web n'est plus attaqué?", ""),
		},
		FirewallRules:       []domains.JsonRule{},
		BypassStage1:        utils.AskInt("À combien de requêtes de contournement par seconde souhaitez-vous activer l'étape 2?", 75),
		BypassStage2:        utils.AskInt("À combien de requêtes de contournement par seconde souhaitez-vous activer l'étape 3?", 250),
		DisableBypassStage3: utils.AskInt("Combien de requêtes de contournement par seconde sont suffisamment faibles pour désactiver l'étape3?", 100),
		DisableRawStage3:    utils.AskInt("Combien de requêtes par seconde sont suffisamment faibles pour désactiver l’étape 3 ? (Les demandes de contournement doivent encore être suffisamment faibles)", 250),
		DisableBypassStage2: utils.AskInt("Combien de requêtes de contournement par seconde sont suffisamment faibles pour désactiver l'étape2?", 50),
		DisableRawStage2:    utils.AskInt("Combien de requêtes par seconde sont suffisamment faibles pour désactiver l’étape 2 ? (Les demandes de contournement doivent encore être suffisamment faibles)", 75),
	}

	domains.Config.Domains = append(domains.Config.Domains, gDomain)

	jsonConfig, err := json.Marshal(domains.Config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("config.json", jsonConfig, 0644)
	if err != nil {
		panic(err)
	}
}

func GetFingerprints(url string, target *map[string]string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("impossible de récupérer les empreintes digitales: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("impossible de récupérer les empreintes digitales: " + err.Error())
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return errors.New("impossible de récupérer les empreintes digitales: " + err.Error())
	}
	return nil
}
