ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine AS build

WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

ARG VERSION
ENV CGO_ENABLED=0

RUN go build -ldflags "-s -X otus-notification/cmd.Version=$VERSION" -v -a -o /bin/otus-notification main.go

FROM scratch
COPY --from=build /bin/otus-notification /bin/otus-notification

ENTRYPOINT ["/bin/otus-notification"]