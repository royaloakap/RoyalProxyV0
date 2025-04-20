package utils

import (
	"encoding/json"
	"fmt"
	"goProxy/core/domains"
	"io/ioutil"
	"strings"
)

func AddDomain() {
	fmt.Println("[ " + PrimaryColor("Aucune configuration de domaine trouvée") + " ]")
	fmt.Println("[ " + PrimaryColor("Configurer de nouveaux domaines dans le Config.json") + " ]")
	fmt.Println("")
	gDomain := domains.Domain{
		Name:        AskString("Quel est le nom de votre domaine (eg. \"example.com\")", "example.com"),
		Backend:     AskString("Quel est le serveur/serveur sur lequel le proxy devrait proxy ??", "1.1.1.1"),
		Scheme:      strings.ToLower(AskString("Quel schéma le proxy doit-il utiliser pour communiquer avec votre backend? (http/https)", "http")),
		Certificate: AskString("Quel est le chemin d’accès au certificat SSL pour votre domaine ? (Laissez vide si vous utilisez le proxy derrière Cloudflare)", ""),
		Key:         AskString("Quel est le chemin d’accès à la clé SSL pour votre domaine ? (Laissez vide si vous utilisez le proxy derrière Cloudflare)", ""),
		Webhook: domains.WebhookSettings{
			URL:            AskString("Quelle est l’URL de votre webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			Name:           AskString("Quel est le nom de votre webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			Avatar:         AskString("Quelle est l’URL de votre avatar Webhook Discord ? (Laissez vide si vous n'en voulez pas)", ""),
			AttackStartMsg: AskString("Quel est le message que votre webhook devrait envoyer lorsque votre site Web est attaqué?", ""),
			AttackStopMsg:  AskString("Quel message votre webhook devrait-il envoyer lorsque votre site Web n'est plus attaqué?", ""),
		},
		FirewallRules:       []domains.JsonRule{},
		BypassStage1:        AskInt("À combien de requêtes de contournement par seconde souhaitez-vous activer l'étape 2?", 75),
		BypassStage2:        AskInt("À combien de requêtes de contournement par seconde souhaitez-vous activer l'étape 3?", 250),
		DisableBypassStage3: AskInt("Combien de requêtes de contournement par seconde sont suffisamment faibles pour désactiver l'étape 3?", 100),
		DisableRawStage3:    AskInt("Combien de requêtes par seconde sont suffisamment faibles pour désactiver l’étape 3 ? (Les demandes de contournement doivent encore être suffisamment faibles)", 250),
		DisableBypassStage2: AskInt("Combien de requêtes de contournement par seconde sont suffisamment faibles pour désactiver l’étape 2 ?", 50),
		DisableRawStage2:    AskInt("Combien de requêtes par seconde sont suffisamment faibles pour désactiver l’étape 2 ? (Les demandes de contournement doivent encore être suffisamment faibles)", 75),
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
