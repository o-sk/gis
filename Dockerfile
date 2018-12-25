FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/o-sk/gis
COPY . .
RUN go mod download
WORKDIR /go/src/github.com/o-sk/gis/web/gis
RUN go build main.go

FROM alpine
COPY --from=builder /go/src/github.com/o-sk/gis/web/gis /app
RUN apk add ca-certificates

CMD /app/main $PORT