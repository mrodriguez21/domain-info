package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/likexian/whois-go"
)

// Regular expression taken from: https://www.socketloop.com/tutorials/golang-use-regular-expression-to-validate-domain-name
func validateDomainName(domain string) bool {
	regExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
 ]{2,3})$`)

	return regExp.MatchString(domain)
}

func validateDomainExists(domain string) bool {
	_, err := net.LookupIP(domain)
	return err == nil
}

func generateURLFromDomain(domain string) string {
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = domain
	return customURL.String()
}

func validImage(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	return strings.Contains(response.Header.Get("Content-Type"), "image")
}

func getInfoFromHTML(url string) (string, string, error) {
	var logo, title string
	// Make HTTP request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request.", err)
		return "", "", err
	}
	defer response.Body.Close()
	defer client.CloseIdleConnections()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Printf("Error loading HTTP response body: %v\n", err)
	}

	head := document.Find("head")

	getIcon := func(i int, elem *goquery.Selection) bool {
		rel, exists := elem.Attr("rel")
		if exists && strings.Contains(rel, "icon") {
			logo = elem.AttrOr("href", "")
		}
		if logo != "" && validImage(logo) {
			return false
		} else if logo != "" && validImage(url+logo) {
			logo = url + logo
			return false
		}
		logo = ""
		return true
	}
	head.Find("link").EachWithBreak(getIcon)

	title, _ = head.Find("title").Html()

	if title == "" {
		getTitle := func(i int, elem *goquery.Selection) bool {
			name, exists := elem.Attr("name")
			if exists && name == "title" {
				title = elem.AttrOr("content", "")
			}
			return title == ""
		}
		head.Find("meta").EachWithBreak(getTitle)
	}

	return logo, title, nil
}

func getWhoIs(ip string) (string, string) {
	var country, owner string

	whoisRaw, err := whois.Whois(ip)
	if err != nil {
		fmt.Println("Error executing whois command: ", err)
	}

	lines, _ := StringToLines(whoisRaw)
	for _, line := range lines {
		if strings.HasPrefix(line, "OrgName:") {
			owner = SplitBySpace(line, 2)[1]
		} else if strings.HasPrefix(line, "Country:") || strings.HasPrefix(line, "country:") {
			country = SplitBySpace(line, 2)[1]
		}
	}

	return country, owner
}

// Returns 0 if g1 == g2, -n if g1 < g2, n if g1 > g2
func compareSSLGrades(g1, g2 string) int {
	if g1 == g2 {
		return 0
	}
	grade1, grade2 := string(g1[0]), string(g2[0])
	if grade1 < grade2 {
		return 1
	} else if grade1 > grade2 {
		return -1
	} else {
		var degree1, degree2 string
		if len(g1) > 1 {
			degree1 = string(g1[1])
		}
		if len(g2) > 1 {
			degree2 = string(g2[1])
		}

		getDegreeValue := func(s string) int {
			switch s {
			case "+":
				return 2
			case "":
				return 1
			case "-":
				return 0
			default:
				return -1
			}
		}
		return getDegreeValue(degree1) - getDegreeValue(degree2)
	}
}

func getSSLInfo(domain string) ([]Server, bool, bool, string, string, time.Time, error) {
	var host HostAPI
	var serversData []Server
	var isDown, serversChanged bool
	var sslGrade, prevSSLGrade string
	var lastChecked time.Time

	// Make HTTP request until ready
	ready := false
	for !ready {
		var curHost HostAPI
		response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
		if err != nil {
			fmt.Printf("Error making HTTP request: %v", err)
			return serversData, isDown, serversChanged, sslGrade, prevSSLGrade, lastChecked, err
		}
		json.NewDecoder(response.Body).Decode(&curHost)

		if curHost.Status == "READY" || curHost.Status == "ERROR" {
			ready = true
			host = curHost
			isDown = curHost.Status == "ERROR"
			lastChecked = time.Now()
			defer response.Body.Close()
		} else {
			response.Body.Close()
			time.Sleep(5 * time.Second)
		}
	}

	for _, server := range host.Servers {
		var serverData Server
		serverData.Address = server.IPAddress
		serverData.SSLGrade = server.Grade
		if sslGrade == "" {
			sslGrade = server.Grade
		} else if server.Grade != "" {
			if compareSSLGrades(server.Grade, sslGrade) < 0 {
				sslGrade = server.Grade
			}
		}
		serverData.Country, serverData.Owner = getWhoIs(serverData.Address)
		if serverData.Owner == "" && server.Name != "" {
			serverData.Country, serverData.Owner = getWhoIs(serverData.Address)
			if serverData.Country != "" || serverData.Owner != "" {
				serverData.Address = server.Name
			}
		}
		serversData = append(serversData, serverData)
	}

	return serversData, isDown, serversChanged, sslGrade, prevSSLGrade, lastChecked, nil
}

func getDomainData(domainName string) (domain Domain, err error) {
	domain.Name = domainName

	// Get SSL and server info
	domain.Servers, domain.IsDown, domain.ServersChanged, domain.SSLGrade, domain.PreviousSSLGrade, domain.LastChecked, err = getSSLInfo(domainName)
	if err != nil {
		fmt.Printf("API is down")
	}

	// Get info from the domain's HTML (logo, title)
	if !domain.IsDown {
		domain.Logo, domain.Title, err = getInfoFromHTML(generateURLFromDomain(domainName))
		if err != nil {
			fmt.Printf("Couldn't get HTML: %v\n", err)
		}
	}

	return
}
