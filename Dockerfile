FROM golang:1.10 as builder

RUN go get -u github.com/Masterminds/glide

WORKDIR /go/src/github.com/mansm/restpruts
COPY *.go  ./
COPY glide* ./


RUN glide install && \
    go build -o app


FROM scratch
COPY --from=builder /go/src/github.com/mansm/restpruts/app .
CMD /app