GOPATH=${HOME}/usr/src::${HOME}/go:/usr/share/gocode
# Safer if not lowercase. Use CamelCase. (watch out in make clean)
APP=demo-adbGoAutomation
# Safer if not empty. (watch out in make clean)
RELEASEDIR=releases/

all: build
build:
	make build-one 	GOOS=linux GOARCH=amd64
build-one:
	go build -o ${RELEASEDIR}${APP}-${GOOS}-${GOARCH}
init:
	go mod init ${APP}/v2
	go mod tidy || true
	# go mod download
	go get
onetime:
	go run .
clean:
	rm -f releases/${APP}-* go.sum go.mod
	go clean
