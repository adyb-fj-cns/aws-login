build:
	go mod vendor
	docker build -f Dockerfile -t aws-login . --rm
	docker create --name aws-login aws-login
	#docker cp aws-login:/aws-login.windows-amd64.exe ./aws-login.windows-amd64.exe
	#docker cp aws-login:/aws-login.linux-amd64 ./aws-login.linux-amd64
	#docker cp aws-login:/aws-login.darwin-amd64 ./aws-login.darwin-amd64

	docker rm aws-login

	
