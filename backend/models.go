package main

import "time"

// Server is a model in the "servers" table.
type Server struct {
	Address  string `json:"address" gorm:"primary_key"`
	SSLGrade string `json:"ssl_grade" gorm:"type:varchar(3)"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

// Domain is a model in the "domains" table.
type Domain struct {
	ServersChanged   bool      `json:"servers_changed"`
	SSLGrade         string    `json:"ssl_grade" gorm:"type:varchar(3)"`
	PreviousSSLGrade string    `json:"previous_ssl_grade" gorm:"type:varchar(3)"`
	Logo             string    `json:"logo"`
	Title            string    `json:"title"`
	IsDown           bool      `json:"is_down"`
	LastChecked      time.Time `json:"-" gorm:"index:lastChecked"`
	Name             string    `json:"-" gorm:"primary_key"`

	Servers []Server `json:"servers" gorm:"many2many:domain_servers"`
}

type DomainResponse struct {
	Name string `json:"domain"`
	Info Domain `json:"info"`
}

type DomainsResponse struct {
	Items []DomainResponse `json:"items"`
}

type HostAPI struct {
	Status  string      `json:"status"`
	Servers []ServerAPI `json:"endpoints"`
}

type ServerAPI struct {
	IPAddress string `json:"ipAddress"`
	Name      string `json:"serverName"`
	Grade     string `json:"grade"`
}
