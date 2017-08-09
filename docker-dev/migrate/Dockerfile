FROM golang:1.8

# Installing github.com/mattes/migrate/cli
RUN go get -u -d github.com/mattes/migrate/cli github.com/go-sql-driver/mysql
RUN go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

# Add entrypoint script
COPY migrate/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Set defaults for entrypoint
ENTRYPOINT ["entrypoint.sh"]