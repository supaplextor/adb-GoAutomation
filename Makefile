GOPATH=${HOME}/go:/usr/share/gocode
# Safer if not lowercase. Use CamelCase. (watch out in make clean)
APP=adbGoAutomator
# Safer if not empty. (watch out in make clean)
RELEASEDIR=releases/

all: build run
build:
	make build-one GOOS=linux GOARCH=amd64
init:
	go mod init ${APP}/v2
	go mod tidy || true
	# go mod download
	go get
onetime:
	go run .
build-one:
	go build -o ${RELEASEDIR}${APP}-${GOOS}-${GOARCH}${dotEXE}
	test "windows" = "${GOOS}" && make release-windows || echo release-windows ${GOOS}
	test "windows" != "${GOOS}" && make release-most || echo release-most ${GOOS}
release-most:
	tar -C ${RELEASEDIR} -cvf ${RELEASEDIR}${APP}-${GOOS}-${GOARCH}.tar.bz2 ${APP}-${GOOS}-${GOARCH} ../README.md ../netmetering.sql ../vendor.env
release-windows:
	zip -j ${RELEASEDIR}${APP}-${GOOS}-${GOARCH}.zip ${RELEASEDIR}${APP}-${GOOS}-${GOARCH}${dotEXE} README.md *.sql vendor.env
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
# on dev env:
run:
	./${RELEASEDIR}${APP}-linux-amd64
