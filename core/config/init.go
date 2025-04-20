package config

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"goProxy/core/db"
	"goProxy/core/domains"
	"goProxy/core/firewall"
	"goProxy/core/proxy"
	"goProxy/core/server"
	"goProxy/core/utils"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kor44/gofilter"
)

func Load() {

	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			Generate()
		} else {
			panic(err)
		}
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&domains.Config)

	proxy.Cloudflare = domains.Config.Proxy.Cloudflare

	if domains.Config.Proxy.Network != "tcp" && domains.Config.Proxy.Network != "tcp4" && domains.Config.Proxy.Network != "tcp6" {
		domains.Config.Proxy.Network = "tcp"
	}

	proxy.CookieSecret = domains.Config.Proxy.Secrets["cookie"]
	if strings.Contains(proxy.CookieSecret, "CHANGE_ME") {
		panic("[ " + utils.PrimaryColor("!") + " ] [ cookie Secret contient « CHANGE_ME », refusant de se charger ]")
	}

	proxy.JSSecret = domains.Config.Proxy.Secrets["javascript"]
	if strings.Contains(proxy.JSSecret, "CHANGE_ME") {
		panic("[ " + utils.PrimaryColor("!") + " ] [ JS Secret contient 'CHANGE_ME', refusant de charger ]")
	}

	proxy.CaptchaSecret = domains.Config.Proxy.Secrets["captcha"]
	if strings.Contains(proxy.CaptchaSecret, "CHANGE_ME") {
		panic("[ " + utils.PrimaryColor("!") + " ] [ Captcha Secret contient « CHANGE_ME », refusant de charger ]")
	}

	proxy.AdminSecret = domains.Config.Proxy.AdminSecret
	if strings.Contains(proxy.AdminSecret, "CHANGE_ME") {
		panic("[ " + utils.PrimaryColor("!") + " ] [ Le secret d'administration contient « CHANGE_ME », refusant de charger ]")
	}

	proxy.APISecret = domains.Config.Proxy.APISecret
	if strings.Contains(proxy.APISecret, "CHANGE_ME") {
		panic("[ " + utils.PrimaryColor("!") + " ] [ Le secret API contient « CHANGE_ME ». Refuser de charger ]")
	}

	// Check if the Proxy Timeout Config has been set otherwise use default values

	if domains.Config.Proxy.Timeout.Idle != 0 {
		proxy.IdleTimeout = domains.Config.Proxy.Timeout.Idle
		proxy.IdleTimeoutDuration = time.Duration(proxy.IdleTimeout).Abs() * time.Second
	}

	if domains.Config.Proxy.Timeout.Read != 0 {
		proxy.ReadTimeout = domains.Config.Proxy.Timeout.Read
		proxy.ReadTimeoutDuration = time.Duration(proxy.ReadTimeout).Abs() * time.Second
	}

	if domains.Config.Proxy.Timeout.ReadHeader != 0 {
		proxy.ReadHeaderTimeout = domains.Config.Proxy.Timeout.ReadHeader
		proxy.ReadHeaderTimeoutDuration = time.Duration(proxy.ReadHeaderTimeout).Abs() * time.Second
	}

	if domains.Config.Proxy.Timeout.Write != 0 {
		proxy.WriteTimeout = domains.Config.Proxy.Timeout.Write
		proxy.WriteTimeoutDuration = time.Duration(proxy.WriteTimeout).Abs() * time.Second
	}

	// Didn't think anyone would actually read through this mess
	if len(domains.Config.Proxy.Colors) != 0 {
		utils.SetColor(domains.Config.Proxy.Colors)
	}

	if domains.Config.Proxy.RatelimitWindow != 0 {
		proxy.RatelimitWindow = domains.Config.Proxy.RatelimitWindow
	}

	proxy.IPRatelimit = domains.Config.Proxy.Ratelimits["requests"]
	proxy.FPRatelimit = domains.Config.Proxy.Ratelimits["unknownFingerprint"]
	proxy.FailChallengeRatelimit = domains.Config.Proxy.Ratelimits["challengeFailures"]
	proxy.FailRequestRatelimit = domains.Config.Proxy.Ratelimits["noRequestsSent"]

	GetFingerprints("https://raw.githubusercontent.com/royaloakap/Version/main/known_fingerprints.json", &firewall.KnownFingerprints)
	GetFingerprints("https://raw.githubusercontent.com/royaloakap/Version/main/bot_fingerprints.json", &firewall.BotFingerprints)
	GetFingerprints("https://raw.githubusercontent.com/royaloakap/Version/main/malicious_fingerprints.json", &firewall.ForbiddenFingerprints)

	for i, domain := range domains.Config.Domains {
		domains.Domains = append(domains.Domains, domain.Name)

		ipInfo := false
		firewallRules := []domains.Rule{}
		rawFirewallRules := domains.Config.Domains[i].FirewallRules
		for _, fwRule := range domains.Config.Domains[i].FirewallRules {

			if strings.Contains(fwRule.Expression, "ip.country") || strings.Contains(fwRule.Expression, "ip.asn") {
				ipInfo = true
			}
			rule, err := gofilter.NewFilter(fwRule.Expression)
			if err != nil {
				panic("[ " + utils.PrimaryColor("!") + " ] [ Erreur lors du chargement des règles de pare-feu personnalisées: " + utils.PrimaryColor(err.Error()) + " ]")
			}

			firewallRules = append(firewallRules, domains.Rule{
				Filter: rule,
				Action: fwRule.Action,
			})
		}

		dProxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: domain.Scheme,
			Host:   domain.Backend,
		})
		dProxy.Transport = &server.RoundTripper{}

		dProxyHandler := adaptor.HTTPHandler(dProxy)

		var cert tls.Certificate = tls.Certificate{}
		if !proxy.Cloudflare {
			var certErr error
			cert, certErr = tls.LoadX509KeyPair(domain.Certificate, domain.Key)
			if certErr != nil {
				panic("[ " + utils.PrimaryColor("!") + " ] [ " + utils.PrimaryColor("Erreur lors du chargement des certificats: "+certErr.Error()) + " ]")
			}
		}

		domains.DomainsMap.Store(domain.Name, domains.DomainSettings{
			Name: domain.Name,

			CustomRules:    firewallRules,
			IPInfo:         ipInfo,
			RawCustomRules: rawFirewallRules,

			DomainProxy:        dProxyHandler,
			DomainCertificates: cert,
			DomainWebhooks: domains.WebhookSettings{
				URL:            domain.Webhook.URL,
				Name:           domain.Webhook.Name,
				Avatar:         domain.Webhook.Avatar,
				AttackStartMsg: domain.Webhook.AttackStartMsg,
				AttackStopMsg:  domain.Webhook.AttackStopMsg,
			},

			BypassStage1:        domain.BypassStage1,
			BypassStage2:        domain.BypassStage2,
			DisableBypassStage3: domain.DisableBypassStage3,
			DisableRawStage3:    domain.DisableRawStage3,
			DisableBypassStage2: domain.DisableBypassStage2,
			DisableRawStage2:    domain.DisableRawStage2,
		})

		firewall.Mutex.Lock()
		domains.DomainsData[domain.Name] = domains.DomainData{
			Name:             domain.Name,
			Stage:            1,
			StageManuallySet: false,
			RawAttack:        false,
			BypassAttack:     false,
			LastLogs:         []string{},

			TotalRequests:    0,
			BypassedRequests: 0,

			PrevRequests: 0,
			PrevBypassed: 0,

			RequestsPerSecond:             0,
			RequestsBypassedPerSecond:     0,
			PeakRequestsPerSecond:         0,
			PeakRequestsBypassedPerSecond: 0,
			RequestLogger:                 []domains.RequestLog{},
		}
		firewall.Mutex.Unlock()
	}

	domains.DomainsMap.Store("debug", domains.DomainSettings{
		Name: "debug",
	})

	firewall.Mutex.Lock()
	domains.DomainsData["debug"] = domains.DomainData{
		Name:             "debug",
		Stage:            0,
		StageManuallySet: false,
		RawAttack:        false,
		BypassAttack:     false,
		BufferCooldown:   0,
		LastLogs:         []string{},

		TotalRequests:    0,
		BypassedRequests: 0,

		PrevRequests: 0,
		PrevBypassed: 0,

		RequestsPerSecond:             0,
		RequestsBypassedPerSecond:     0,
		PeakRequestsPerSecond:         0,
		PeakRequestsBypassedPerSecond: 0,
		RequestLogger:                 []domains.RequestLog{},
	}

	firewall.Mutex.Unlock()

	vcErr := VersionCheck()
	if vcErr != nil {
		panic("[ " + utils.PrimaryColor("!") + " ] [ " + vcErr.Error() + " ]")
	}

	if len(domains.Domains) == 0 {
		AddDomain()
		Load()
	} else {
		proxy.WatchedDomain = domains.Domains[0]
		db.Connect()
	}
}

func VersionCheck() error {
	resp, err := http.Get("https://raw.githubusercontent.com/royaloakap/Version/main/version.json")
	if err != nil {
		return errors.New("Failed to check for proxy version: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Failed to check for proxy version: " + err.Error())
	}

	var proxyVersions GLOBAL_PROXY_VERSIONS
	err = json.Unmarshal(body, &proxyVersions)
	if err != nil {
		return errors.New("Failed to check for proxy version: " + err.Error())
	}

	if proxy.ProxyVersion < proxyVersions.StableVersion {
		fmt.Println("[ " + utils.PrimaryColor("!") + " ] [ Nouvelle Version de Royal Proxy " + fmt.Sprint(proxyVersions.StableVersion) + " Trouvée. Vous utilisez " + fmt.Sprint(proxy.ProxyVersion) + ". Contactez @Royaloakap ou Pensez à télécharger la nouvelle version depuis Github " + proxyVersions.Download + " ]")

		fmt.Println("[ " + utils.PrimaryColor("!") + " ] [ Sortie en raison de la version obsolète de Royal Proxy / t.me/ROYAL_FAQ ]")
		os.Exit(1)

	}

	return nil
}
