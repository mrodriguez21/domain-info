package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
)

// Service is an http server that handles REST requests.
type Service struct {
	db      *DBConnection
	Handler fasthttp.RequestHandler
}

// NewService creates a new instance of a Service.
func NewService(db *DBConnection) Service {
	service := Service{db: db}
	router := fasthttprouter.New()
	router.GET("/domains/:domain", service.getDomain)
	router.GET("/domains", service.getPreviousDomains)
	service.Handler = cors.Default().Handler(router.Handler)

	return service
}

func (s *Service) getDomain(ctx *fasthttp.RequestCtx) {
	domainName, ok := ctx.UserValue("domain").(string)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	}
	if isDomain := validateDomainName(domainName); !isDomain {
		fmt.Println("Invalid domain.")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	}

	domain, _ := s.db.getDomain(domainName)

	if domain.Name == "" {
		// Get the domain's info and save to DB
		domain, _ = getDomainData(domainName)
		s.db.createDomain(domain)
	} else if time.Since(domain.LastChecked) > time.Hour {
		// Get the current state of the domain's servers
		newDomain, _ := getDomainData(domainName)
		newDomain.PreviousSSLGrade = domain.SSLGrade

		// Check if servers have changed
		hasChanged := domain.IsDown != newDomain.IsDown
		if !hasChanged {
			oldServers := make(map[string]Server)
			newServers := make(map[string]Server)

			for _, server := range domain.Servers {
				oldServers[server.Address] = server
			}
			for _, server := range newDomain.Servers {
				newServers[server.Address] = server
			}

			fmt.Println(oldServers)
			fmt.Println(newServers)

			// Check if old server is same as current or no longer exists
			for _, server := range domain.Servers {
				newServer, ok := newServers[server.Address]
				if !ok {
					s.db.deleteServer(server.Address)
					hasChanged = true
				} else if newServer != server {
					s.db.updateServer(newServer)
					hasChanged = true
				}
			}

			// Check if new server didn't exist
			for _, server := range newDomain.Servers {
				_, ok := oldServers[server.Address]
				if !ok {
					s.db.createServer(server, domain.Name)
					hasChanged = true
				}
			}
		}

		newDomain.ServersChanged = hasChanged

		//Update domain's record in DB
		domain = newDomain
		s.db.updateDomain(domain)
	}

	json.NewEncoder(ctx).Encode(&domain)
	ctx.SetContentType("application/JSON")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Service) getPreviousDomains(ctx *fasthttp.RequestCtx) {
	var response DomainsResponse
	domains, _ := s.db.getDomains()
	for _, domain := range domains {
		var domainResponse DomainResponse
		domainResponse.Info = domain
		domainResponse.Name = domain.Name
		response.Items = append(response.Items, domainResponse)
	}
	json.NewEncoder(ctx).Encode(&response)
	ctx.SetContentType("application/JSON")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
