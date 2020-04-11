# Domain Info

Domain Info is a web app for consulting information about a certain domain and its servers. For each domain we can see the following:

* Status - Up or down
* Logo - taken from the HTML's head
* Title - also taken from the HTML's head
* SSL Grade - the lowest SSL grade of all its servers
* Previous SSL Grade - the grade it had on the last request made (an hour or more ago)
* Servers* 
  * Address - server's IP or host
  * SSL Grade - given by [SSL Labs](https://www.ssllabs.com/)
  * Country - get using  `whois <ip>`
  * Owner - name of the organization that owns the IP, get using  `whois <ip>`

*We get this information using the [SSL Labs API](https://www.ssllabs.com/projects/ssllabs-apis/index.html).

You can also see a list with all the past requests that have been made.

## Requirements

You need to have the following installed in your computer:

* NodeJS - install it [here](https://nodejs.org/en/download/)
* Go - install it [here](https://golang.org/doc/install)
* CockroachDB - install it [here](https://www.cockroachlabs.com/docs/dev/install-cockroachdb-windows.html)

## Deployment

To run this project locally, download or `git clone` the repository. Then

### Database setup

1. Create a new directory for storing your database files.

2. Inside that directory run:

   ```bash
   cockroach start-single-node --insecure # Starts a single-node cluster
   cockroach sql --insecure # Connects to the cluster via SQL shell
   CREATE DATABASE domain_info; # Creates a new database
   ```

**It will deploy in:**

```
http://localhost:26257/
```

### Back end

1. Go to `backend` directory
2. Run the following commands:

```bash
 make # Installs needed dependencies and builds project

.\backend # Runs the project
```

**It will deploy in:**

```
http://localhost:8090/
```

You can deploy it in other port by running:

```bash
.\backend -port=":9090" 
```

> *Note:* The URL for connecting to the database is set by default to `"postgresql://root@localhost:26257/domain_info?sslmode=disable"`. You can change it using the `-addr` flag.

### Front end

1. Go to `frontend` directory
2. Run the following commands:

```bash
npm install  # Installs needed dependencies 
npm run serve # Compiles and hot-reloads
```

**It will deploy in:**

```
http://localhost:8081/
```

## Author

- [Mariana Rodríguez](https://mrodriguez21.github.io/) 👩‍💻

## MIT License

This project is licensed by the MIT [License](https://github.com/mrodriguez21/domain-info/blob/master/LICENSE).

