GOPATH=${HOME}/go:/usr/share/gocode
# Safer if not lowercase. Use CamelCase. (watch out in make clean)
APP=adbGoAutomator
# Safer if not empty. (watch out in make clean)
RELEASEDIR=releases/

all: build
build:
	make build-one GOOS=linux GOARCH=amd64
build-one:
	go build
	make -C sample-apps/demo/ clean init build-one
init:
	go mod init ${APP}/v2
	go mod tidy || true
	# go mod download
	go get
onetime:
	go run .
build-release: build
	make build-one GOOS=darwin GOARCH=amd64
	make build-one GOOS=linux GOARCH=386
	make build-one GOOS=windows GOARCH=386 dotEXE=.exe
	make build-one GOOS=windows GOARCH=amd64 dotEXE=.exe
build-more:
	make build-one GOOS=android GOARCH=arm || true
	make build-one GOOS=darwin GOARCH=386 || true
	make build-one GOOS=darwin GOARCH=arm || true
	make build-one GOOS=darwin GOARCH=arm64 || true
	make build-one GOOS=plan9 GOARCH=386 || true
	make build-one GOOS=plan9 GOARCH=amd64 || true
	make build-one GOOS=solaris GOARCH=amd64 || true
	make build-one GOOS=dragonfly GOARCH=amd64 || true
	make build-one GOOS=freebsd GOARCH=386 || true
	make build-one GOOS=freebsd GOARCH=amd64 || true
	make build-one GOOS=freebsd GOARCH=arm || true
	make build-one GOOS=linux GOARCH=arm || true
	make build-one GOOS=linux GOARCH=arm64 || true
	make build-one GOOS=linux GOARCH=ppc64 || true
	make build-one GOOS=linux GOARCH=ppc64le || true
	make build-one GOOS=linux GOARCH=mips || true
	make build-one GOOS=linux GOARCH=mipsle || true
	make build-one GOOS=linux GOARCH=mips64 || true
	make build-one GOOS=linux GOARCH=mips64le || true
	make build-one GOOS=netbsd GOARCH=386 || true
	make build-one GOOS=netbsd GOARCH=amd64 || true
	make build-one GOOS=netbsd GOARCH=arm || true
	make build-one GOOS=openbsd GOARCH=386 || true
	make build-one GOOS=openbsd GOARCH=amd64 || true
	make build-one GOOS=openbsd GOARCH=arm || true
clean:
	rm -f ${RELEASEDIR}${APP}-* go.sum go.mod
	go clean
	make -C sample-apps/demo/ clean
# on dev env:
run:
	./${RELEASEDIR}${APP}-linux-amd64
