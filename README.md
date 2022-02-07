# Gosnel Skeleton Application

## .env Configuration Example

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
