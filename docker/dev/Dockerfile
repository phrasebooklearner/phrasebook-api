FROM phrasebooklearner/phrasebook-base-image:latest
MAINTAINER oliosinter <olio.sinter@gmail.com>

# Setting cache directory (though glide still doesn't support it)
ENV XDG_CACHE_HOME=/cache

# OUTPUT: Build artefacts are output here
VOLUME /go/bin

# OUTPUT: Downloaded dependencies are output here
VOLUME /dependencies

# OUTPUT: Test reports are output here
VOLUME /reports

# Add test entrypoint script
COPY scripts/test.sh /usr/local/bin/test.sh
RUN chmod +x /usr/local/bin/test.sh

# Add build entrypoint script
COPY scripts/build.sh /usr/local/bin/build.sh
RUN chmod +x /usr/local/bin/build.sh

COPY ./glide.yaml /go/src/phrasebook-api/
COPY ./glide.lock /go/src/phrasebook-api/
COPY ./main.go /go/src/phrasebook-api/
COPY ./src /go/src/phrasebook-api/src

WORKDIR /go/src/phrasebook-api