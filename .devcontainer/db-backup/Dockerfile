FROM golang:alpine
WORKDIR /app
COPY . /app
RUN go build -o db-backup
RUN apk add --no-cache postgresql-client
RUN apk add --no-cache mariadb-client
RUN apk add --no-cache mariadb-connector-c
RUN apk --no-cache add mongodb-tools
RUN go install
# VOLUME /data
CMD ["/bin/sh"]