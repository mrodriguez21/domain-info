package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	username string = "marshvee"
	dbHost   string = "localhost"
	dbPort   int    = 26257
	dbName   string = "domain_info"
)

// DBConnection stores the reference to the connection to the DB
type DBConnection struct {
	db *gorm.DB
}

// NewDBConnection creates a new instance of a DBConnection.
func NewDBConnection() *DBConnection {
	db := setupDB()
	return &DBConnection{db: db}
}

func setupDB() *gorm.DB {
	// Connect to CockroachDB
	addr := fmt.Sprintf("postgresql://%v@%v:%v/%v?sslmode=disable", username, dbHost, dbPort, dbName)
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Set to `true` and GORM will print out all DB queries.
	db.LogMode(false)

	// Automatically create the "servers" and "domains" table.
	db.AutoMigrate(&Domain{}, &Server{})

	return db
}

func (connection *DBConnection) createServer(server Server) error {
	return connection.db.Create(&server).Error
}

func (connection *DBConnection) updateServer(server Server) error {
	return connection.db.Model(&server).Where("ADDRESS = ?", server.Address).Update(server).Error
}

func (connection *DBConnection) deleteServer(serverAddress string) error {
	req := connection.db.Delete(Server{}, "ADDRESS = ?", serverAddress)
	if req.RowsAffected == 0 {
		fmt.Println("Server was not found.")
	}
	return req.Error
}

func (connection *DBConnection) getDomains() ([]Domain, error) {
	var domains []Domain
	err := connection.db.Preload("Servers").Find(&domains).Error
	return domains, err
}

func (connection *DBConnection) getDomain(domainName string) (Domain, error) {
	var domain Domain
	err := connection.db.Preload("Servers").Where("NAME = ?", domainName).Find(&domain).Error
	return domain, err
}

func (connection *DBConnection) createDomain(domain Domain) error {
	for _, server := range domain.Servers {
		connection.createServer(server)
	}
	return connection.db.Create(&domain).Error
}

func (connection *DBConnection) updateDomain(domain Domain) error {
	return connection.db.Model(&domain).Where("NAME = ?", domain.Name).Update(domain).Error
}
