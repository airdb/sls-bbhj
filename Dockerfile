# Stage 1: Builder
FROM golang
MAINTAINER info@airdb.com

ENV GITHUB=github.com/airdb/mina-api
ENV BUILDDIR=/go/src/${GITHUB}

WORKDIR ${BUILDDIR}

ADD . ${BUILDDIR}

RUN go get -u github.com/swaggo/swag/cmd/swag && \
	GO111MODULE=off swag init -g web/router.go && \
	go build -o main main.go


# Stage 2: Release the binary from the builder stage
FROM golang

ENV GITHUB=github.com/airdb/mina-api
ENV BUILDDIR=/go/src/${GITHUB}

COPY --from=0 ${BUILDDIR}/ /srv

EXPOSE 8080

WORKDIR /srv
CMD ["/srv/main"]
