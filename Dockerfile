FROM golang:1.13 as builder

COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d ./... && \
	go test ./... && \
	go build -a -installsuffix cgo -o /assets/server /src/server && \
	chmod +x /assets/server

FROM scratch as server

COPY --from=builder /assets/server /opt/server/
CMD ["/opt/server/server"]
