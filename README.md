# FourLeafPortalApi
This project was generated with Go cli.

## MySQL Server
Make sure you've installed mysql server and docker on your machine and have set up a username and password. Open the databases/user.go file and replace username 'root' and password 'password' with your local mysql credentials.
Run `docker build -t mysql_db:mysql_db .` to build the docker image with the mysql instance
Then Run `docker run -p 4400:3306 mysql_db` to run the image created previously

## Development Server
Run `go run .` to start server, running on http://127.0.0.1:4300

## Sources
This Golang project is adapted from the [following walkthrough](https://codesource.io/how-to-setup-golang-authentication-with-jwt-token/).