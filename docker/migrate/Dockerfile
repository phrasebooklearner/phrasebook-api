FROM phrasebooklearner/phrasebook-base-image:latest
MAINTAINER oliosinter <olio.sinter@gmail.com>

# Installing github.com/mattes/migrate/cli
RUN go get -u -d github.com/mattes/migrate/cli github.com/go-sql-driver/mysql
RUN go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

# Add entrypoint script
COPY scripts/migrate.sh /usr/local/bin/migrate.sh
RUN chmod +x /usr/local/bin/migrate.sh

# Set defaults for entrypoint
ENTRYPOINT ["migrate.sh"]
# Set default command
CMD ["-path", "/migrations", "up"]

COPY ./migrations /migrations