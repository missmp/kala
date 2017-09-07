FROM golang

RUN go get github.com/missmp/kala
ENTRYPOINT kala run
EXPOSE 8000
