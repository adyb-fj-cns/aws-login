FROM golang:latest as build

RUN mkdir -p /temp
WORKDIR /temp
ADD . .

RUN GOOS=windows GOARCH=amd64 go build -o aws-login.windows-amd64.exe  ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o aws-login.linux-amd64  ./main.go
RUN GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o aws-login.darwin-amd64  ./main.go

FROM alpine:latest

COPY --from=build /temp/aws-login.windows-amd64.exe .
COPY --from=build /temp/aws-login.linux-amd64 .
COPY --from=build /temp/aws-login.darwin-amd64 .

RUN chmod +x ./aws-login.linux-amd64
RUN chmod +x ./aws-login.darwin-amd64
