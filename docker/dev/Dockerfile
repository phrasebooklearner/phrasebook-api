FROM phrasebooklearner/phrasebook-base-image:latest
MAINTAINER oliosinter <olio.sinter@gmail.com>

# Installing go-junit-report (Converts go test output to an xml report, suitable for applications that expect junit xml reports (e.g. Jenkins))
RUN go get -u github.com/jstemmer/go-junit-report

# Setting cache directory (though glide still doesn't support it)
ENV XDG_CACHE_HOME=/cache

RUN mkdir /reports

# Add test entrypoint script
COPY scripts/test.sh /usr/local/bin/test.sh
RUN chmod +x /usr/local/bin/test.sh

# Add test command script
COPY scripts/test-run.sh /usr/local/bin/test-run.sh
RUN chmod +x /usr/local/bin/test-run.sh

# Add build entrypoint script
COPY scripts/build.sh /usr/local/bin/build.sh
RUN chmod +x /usr/local/bin/build.sh

COPY ./glide.yaml /go/src/phrasebook-api/
COPY ./glide.lock /go/src/phrasebook-api/
COPY ./main.go /go/src/phrasebook-api/
COPY ./src /go/src/phrasebook-api/src

WORKDIR /go/src/phrasebook-api