all: build package

build:
	go mod vendor
	docker build -f Dockerfile -t aws-login-builder . --rm

package:
	docker create --name aws-login-builder aws-login-builder
	docker cp aws-login-builder:/aws-login.windows-amd64.exe ./aws-login.windows-amd64.exe
	docker cp aws-login-builder:/aws-login.linux-amd64 ./aws-login.linux-amd64
	docker cp aws-login-builder:/aws-login.darwin-amd64 ./aws-login.darwin-amd64
	docker rm aws-login-builder

package-windows:
	docker create --name aws-login-builder aws-login-builder
	docker cp aws-login-builder:/aws-login.windows-amd64.exe ./aws-login.exe
	docker rm aws-login-builder

windows: build package-windows



	

	
