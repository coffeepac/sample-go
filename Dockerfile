FROM golang

ADD sample-go.go /

RUN go build -o /sample-go /sample-go.go

EXPOSE 8080
ENTRYPOINT ["/sample-go"]