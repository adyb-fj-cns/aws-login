FROM golang:latest as build

RUN mkdir -p /go/src/github.com/adyb-fj-cns/adybuxton/aws-login/
WORKDIR /go/src/github.com/adyb-fj-cns/adybuxton/aws-login
ADD . .

RUN GOOS=windows GOARCH=amd64 go build -o aws-login.exe  ./main.go
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o aws-login.linux-amd64  ./main.go
RUN GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o aws-login.darwin-amd64  ./main.go

COPY /go/src/github.com/adyb-fj-cns/aws-login/aws-login.windows-amd64.exe /aws-login.windows-amd64.exe
#COPY /go/src/github.com/adyb-fj-cns/aws-login/aws-login.linux-amd64 /aws-login.linux-amd64
COPY /go/src/github.com/adyb-fj-cns/aws-login/aws-login.darwin-amd64 /aws-login.darwin-amd64

FROM alpine:latest
COPY --from=build /go/src/github.com/adyb-fj-cns/aws-login/aws-login.windows-amd64.exe .
#COPY --from=build /go/src/github.com/adyb-fj-cns/aws-login/aws-login.linux-amd64 .
COPY --from=build /go/src/github.com/adyb-fj-cns/aws-login/aws-login.darwin-amd64 .

#RUN chmod +x ./aws-login.linux-amd64
RUN chmod +x ./aws-login.darwin-amd64
