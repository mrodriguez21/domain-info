package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	addr string = "postgresql://marshvee@localhost:26257/domain_info?sslmode=disable"
)

// DBConnection stores the reference to the connection to the DB
type DBConnection struct {
	db *sql.DB
}

// NewDBConnection creates a new instance of a DBConnection.
func NewDBConnection() *DBConnection {
	db := setupDB()
	return &DBConnection{db: db}
}

func setupDB() *sql.DB {
	// Connect to CockroachDB
	db, err := sql.Open("postgres", addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Create the "servers" and "domains" table.
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS domains (
		servers_changed BOOL NULL,
		ssl_grade VARCHAR(3) NULL,
		previous_ssl_grade VARCHAR(3) NULL,
		logo STRING NULL,
		title STRING NULL,
		is_down BOOL NULL,
		last_checked TIMESTAMPTZ NULL,
		name STRING NOT NULL,
		PRIMARY KEY (name),
		INDEX lastchecked (last_checked DESC))
	`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS servers (
		address STRING NOT NULL,
		ssl_grade VARCHAR(3) NULL,
		country STRING NULL,
		owner STRING NULL,
		domain_name STRING NOT NULL,
		FOREIGN KEY (domain_name) REFERENCES domains(name))
	`); err != nil {
		log.Fatal(err)
	}

	return db
}

func (connection *DBConnection) createServer(server Server, domainName string) error {
	_, err := connection.db.Exec(
		`INSERT INTO servers (address, ssl_grade, country, owner, domain_name)
		VALUES ($1, $2, $3, $4, $5)`,
		server.Address,
		server.SSLGrade,
		server.Country,
		server.Owner,
		domainName,
	)
	return err
}

func (connection *DBConnection) updateServer(server Server) error {
	_, err := connection.db.Exec(
		`UPDATE servers 
		SET ssl_grade = $1, country = $2, owner = $3
		WHERE address = $4`,
		server.SSLGrade,
		server.Country,
		server.Owner,
		server.Address,
	)

	return err
}

func (connection *DBConnection) deleteServer(serverAddress string) error {
	_, err := connection.db.Exec(
		`DELETE FROM servers 
		WHERE address = $1`,
		serverAddress,
	)

	return err
}

func (connection *DBConnection) getDomains() ([]Domain, error) {
	var domains []Domain
	rows, err := connection.db.Query("SELECT * FROM domains ORDER BY last_checked DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var domain Domain
		err = rows.Scan(
			&domain.ServersChanged, &domain.SSLGrade,
			&domain.PreviousSSLGrade, &domain.Logo, &domain.Title,
			&domain.IsDown, &domain.LastChecked, &domain.Name,
		)
		if err != nil {
			log.Fatal(err)
		}

		rows2, err := connection.db.Query(
			`SELECT address, ssl_grade, country, owner FROM servers WHERE domain_name = $1`,
			domain.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer rows2.Close()
		for rows2.Next() {
			var server Server
			err := rows2.Scan(&server.Address, &server.SSLGrade, &server.Country, &server.Owner)
			if err != nil {
				log.Fatal(err)
			}
			domain.Servers = append(domain.Servers, server)
		}
		domains = append(domains, domain)
	}
	return domains, err
}

func (connection *DBConnection) getDomain(domainName string) (Domain, error) {
	var domain Domain
	rows, err := connection.db.Query(`SELECT * FROM domains WHERE name = $1`, domainName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(
			&domain.ServersChanged, &domain.SSLGrade,
			&domain.PreviousSSLGrade, &domain.Logo, &domain.Title,
			&domain.IsDown, &domain.LastChecked, &domain.Name,
		)
		if err != nil {
			log.Fatal(err)
		}

		rows2, err := connection.db.Query(
			`SELECT address, ssl_grade, country, owner FROM servers WHERE domain_name = $1`,
			domain.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer rows2.Close()
		for rows2.Next() {
			var server Server
			err := rows2.Scan(&server.Address, &server.SSLGrade, &server.Country, &server.Owner)
			if err != nil {
				log.Fatal(err)
			}
			domain.Servers = append(domain.Servers, server)
		}
	}

	return domain, err
}

func (connection *DBConnection) createDomain(domain Domain) {
	connection.db.Exec(
		`INSERT INTO domains 
			(servers_changed, ssl_grade, previous_ssl_grade,
			logo, title, is_down, last_checked, name)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		domain.ServersChanged,
		domain.SSLGrade,
		domain.PreviousSSLGrade,
		domain.Logo,
		domain.Title,
		domain.IsDown,
		domain.LastChecked,
		domain.Name,
	)

	for _, server := range domain.Servers {
		connection.createServer(server, domain.Name)
	}
}

func (connection *DBConnection) updateDomain(domain Domain) error {
	_, err := connection.db.Exec(
		fmt.Sprintf(
			`UPDATE domains 
			SET servers_changed = %v, ssl_grade = '%v', previous_ssl_grade = '%v', 
			logo = '%v', title = '%v', is_down = %v, last_checked = '%v'
			WHERE name = '%v'`,
			domain.ServersChanged,
			domain.SSLGrade,
			domain.PreviousSSLGrade,
			domain.Logo,
			domain.Title,
			domain.IsDown,
			domain.LastChecked,
			domain.Name,
		),
	)

	return err
}
