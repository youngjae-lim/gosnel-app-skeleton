# Gosnel Skeleton Application

## Overview

## Features

- Migrations
- Authentication
- Models
- Handlers
- Sessions
- Mailing
- Remote File Systems
  - S3
  - Minio
  - SFTP
  - WEBDAV

- Maintenance Mode
- Screen Capture of testing webpage

## Configurations

### Postgres database

1. Configure ```/config/database.yml```

```yml
development:
  dialect: postgres
  database: gosnel
  user: postgres
  password: password
  host: localhost
  port: 5432
  pool: 5
```

2. Configure  .env

```env
# database config - postgres or mysql
# 💁 make sure you stop postgres service on your local machine
# otherwise, it will try to connect to postgre db on your local machine, not on docker container
DATABASE_TYPE=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASS=password
DATABASE_NAME=gosnel
DATABASE_SSL_MODE=disable
```

## CLI Application

### Migrations

### Authentication

To make authentication functionality, run the follwoing in the project root:

```shell
$gosnel make auth
```

Then open up ```models.go``` and update Model struct and New() function as follows:

```go
// models.go
package data

import (
 "database/sql"
 "fmt"
 "os"

 db2 "github.com/upper/db/v4"
 "github.com/upper/db/v4/adapter/mysql"
 "github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

// Models is the wrapper for all database models.
type Models struct {
 // any models inserted here (and in the New function)
 // are easily accessible throughout the entire application

 // For auth, these models are needed
 RememberToken RememberToken
 Tokens        Token
 Users         User
}

// New initializes the models package for use.
func New(databasePool *sql.DB) Models {
 db = databasePool

 switch os.Getenv("DATABASE_TYPE") {
 case "mysql", "mariadb":
  upper, _ = mysql.New(databasePool)
 case "postgres", "postgresql":
  upper, _ = postgresql.New(databasePool)
 default:
  // do nothing
 }

 return Models{
  RememberToken: RememberToken{},
  Users:         User{},
  Tokens:        Token{},
 }
}

// getInsertID returns the integer value of a newly inserted id (using upper).
func getInsertID(i db2.ID) int {
 idType := fmt.Sprintf("%T", i)
 if idType == "int64" {
  return int(i.(int64))
 }

 return i.(int)
}
```

Then open up ```routes.go``` file and add the routes as following:

```go
// routes.go
package main

import (
 "net/http"

 "github.com/go-chi/chi/v5"
 "github.com/youngjae-lim/gosnel"
)

func (a *application) routes() *chi.Mux {
 // middleware must come before any routes

 // add routes here
 a.get("/", a.Handlers.Home)

 // add these routes after running 'gosnel make auth'
 a.get("/users/login", a.Handlers.UserLogin)
 a.post("/users/login", a.Handlers.PostUserLogin)
 a.get("/users/logout", a.Handlers.Logout)

 a.get("/auth/{provider}", a.Handlers.SocialLogin)
 a.get("/auth/{provider}/callback", a.Handlers.SocialMediaCallBack)

 // static files
 fileServer := http.FileServer(http.Dir("./public"))
 a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

 // routes from gosnel
 a.App.Routes.Mount("/gosnel", gosnel.Routes())
 a.App.Routes.Mount("/api", a.ApiRoutes())

 return a.App.Routes
}


```

### Models

### Handlers

### Sessions

### Mailing

### Maintenance Mode

To go into maintenance mode,

```shell
$gosnel down
```

To go back into live mode,

```shell
$gosnel up
```

### Example of .env configuration

```env
# Give your application a unique name (no spaces)
APP_NAME=myapp
APP_URL=http://localhost:4000

# false for production, true for development
DEBUG=true

# the port should we listen on
PORT=4000
RPC_PORT=12345
ALLOWED_URLS="/login,/admin"

# the server name, e.g., www.mysite.com
SERVER_NAME=localhost

# should we use HTTPS?
SECURE=false

# database config - postgres or mysql
# 💁 make sure you stop postgres service on your local machine
# otherwise, it will try to connect to postgre db on your local machine, 
# not on docker container
DATABASE_TYPE=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASS=password
DATABASE_NAME=gosnel
DATABASE_SSL_MODE=disable

# redis config
REDIS_HOST=localhost:6379
REDIS_PASSWORD=
REDIS_PREFIX=myapp

# cache: Redis or badger
CACHE=redis

# cooking seetings
COOKIE_NAME=myapp
COOKIE_LIFETIME=1440
COOKIE_PERSIST=true
COOKIE_SECURE=false
COOKIE_DOMAIN=localhost

# session store: cookie, redis, mysql, or postgres
SESSION_TYPE=redis

# mail settings
SMTP_HOST=
SMTP_USERNAME=
SMTP_PASSWORD=
SMTP_PORT=
SMTP_ENCRYPTION=
# domain is not required for sendgrid
MAIL_DOMAIN=
FROM_NAME=
FROM_ADDRESS=

# mail settings for API services
# For MAILER_API, please type smtp, sendgrid, mailgun, or sparkhost
MAILER_API=
MAILER_KEY=
# url is not required for SendGrid
MAILER_URL=

# template engine: go or jet
RENDERER=jet

# the encryption key; must be exactly 32 characters long
KEY=l7HrN539xnhnNvjFfVv5NF9ThfLdl5rX

# remote file system variables for s3, minio, sftp, and webdav

# S3 object storage for linode
S3_SECRET=<your_s3_secret>
S3_KEY=<your_s3_key>
S3_REGION=us-east-1
S3_ENDPOINT=us-east-1.linodeobjects.com
S3_BUCKET=gosnel

# Minio
MINIO_ENDPOINT=127.0.0.1:9000
MINIO_KEY=root
MINIO_SECRET=password
MINIO_USESSL=false
MINIO_REGION=us-east-1
MINIO_BUCKET=testbucket

# SFTP
SFTP_HOST=localhost
SFTP_USER=sftp
SFTP_PASS=password
SFTP_PORT=2022

# WEBDAV
WEBDAV_HOST=http://127.0.0.1:10080
WEBDAV_USER=sftp
WEBDAV_PASS=password

# permitted upload types
ALLOWED_MIMETYPES="image/gif,image/jpeg,image/png,application/pdf"
# max upload size 10mb
MAX_UPLOAD_SIZE=1048576000

# social login for github
GITHUB_KEY=<your_github_key>
GITHUB_SECRET=<your_github_secret>
GITHUB_CALLBACK=http://localhost:4000/auth/github/callback

# social login for google
GOOGLE_KEY=<your_google_key>
GOOGLE_SECRET=<your_google_secret>
GOOGLE_CALLBACK=http://localhost:4000/auth/google/callback
```
