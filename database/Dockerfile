FROM postgres:16-alpine

RUN apk add --no-cache gettext

COPY ./files/docker-entry-point-initdb.d/ /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/*.sh

EXPOSE 5432