FROM golang:1.13 as builder

COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d ./... && \
	go test ./... && \
	go build -a -installsuffix cgo -o /assets/server /src/server && \
	go build -a -installsuffix cgo -o /assets/client /src/client && \
	chmod +x /assets/server && \
	chmod +x /assets/client

FROM scratch as arithmetic

COPY --from=builder /assets/server /
COPY --from=builder /assets/client /
ENTRYPOINT ["/server"]
