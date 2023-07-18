# Jackbox Backend
This project was generated with Go cli.

## MySQL Server
- Make sure you've installed mysql server and docker on your machine and have set up a password for the root user for mysql. 
- Open the /databases/user.go file and replace password 'password' with your local mysql root password. 
- Open up the Dockerfile and replace `password` with your mysql root user's password. 
- This is likely not necessary, but if you have errors running the docker container, you may need to run the following `docker pull mcr.microsoft.com/mssql/server:2022-latest`.
- Run `docker build -t mysql_db .` to build the docker image with the mysql instance. 
- Run `docker run -p 4400:3306 -d mysql_db` to run the image created previously.

## Development Server
Run `go run .` to start the server, running on http://127.0.0.1:4300

## Sources
This Golang project is built up from [a previous project of mine](https://github.com/weldonla/OneCauseGolangLogin), which was originally adapted from the [following walkthrough](https://codesource.io/how-to-setup-golang-authentication-with-jwt-token/).