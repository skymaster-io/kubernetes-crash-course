FROM golang:1.18 as builder
COPY /src /src
WORKDIR /src
RUN go build -o hello-world main.go

FROM debian
COPY --from=builder /src/hello-world /bin/hello-world
CMD [ "/bin/hello-world" ]