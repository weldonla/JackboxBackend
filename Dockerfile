FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=password

COPY ./migrations/user-up.sql /docker-entrypoint-initdb.d/