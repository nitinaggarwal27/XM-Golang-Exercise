# XM-Golang-Exercise - v21.0.0

In this exercise you need to build a REST API microservice to handle Companies.
Company is an entity defined by the following attributes:

* Name
* Code
* Country
* Website
* Phone

All four CRUD operations are required.
For read operation, fetching one or many companies should be available.
Each Company’s attribute should be available as filtering in the CRUD operations.

### Option 1
Creation and deletion operations must be allowed only for requests received from users located
in Cyprus. The location must be retrieved based on the user’s IP address via the service
https://ipapi.co/.

### Option 2
Creation and deletion operations must be allowed only for authenticated requests.
JWT or any other authentication/authorization standard can be used.

### Optional
On each mutating operation, a JSON formatted event must be produced to a service bus (Kafka,
RabbitMQ etc.). In a production environment, those events could be used to notify other
microservices of a data change and conduct some business logic, i.e. sending emails.
## Requirements ##

1. SQLite/Postgres/MySQL Database


## Environment Variables => ##

```
//Service based
export METASECURE_AUTH_PORT="8000"
export BUILD_IMAGE="v27.0.0"
export ENVIRONMENT="production"


// database
export DB_ENGINE="sqlite"
export DB_NAME="xm_exercise"
// pass for postgres or mysql
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASS=""

//jwt Configuration
export PRIVATE_KEY="same among all the metasecure backend services"

//admin credentials
export ADMIN_EMAIL="admin@xm.com"
export ADMIN_PASS="*****"
```

## How to run the app ##

### 1. Configuration using environment variables ###

```
i.    Export above all environment variables
ii.   Build the app or binary -> command -> `$ go install`
iii.  Run the app or binary -> command -> `$GOPATH/bin/auth-service --conf=environment`
```

### 2. Configuration using TOML file ###

```
i.    Create a configuration toml file by taking reference from example.toml file
ii.   Build the app or binary -> command -> `$ go install`
iii.  Run the app or binary -> command -> `$GOPATH/bin/auth-service --conf=toml --file=<path of toml file`
```

`Note :- for any help regarding flags, run this command '$GOPATH/bin/auth-service --help'`
