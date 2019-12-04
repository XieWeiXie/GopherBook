BINARY=BeeQuick

VERSION=1.0.0

BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Env=production -s -w"

DEV_LDFLAGS=-ldflags "-X main.Env=dev"

TEST_LDFLAGS=-ldflags "-X main.Env=test"

default:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

production:
	go build -o ${BINARY} -v ${LDFLAGS} -tags=jsoniter

dev:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

test:
	go build -o ${BINARY} -v ${TEST_LDFLAGS} -tags=jsoniter

run:
	go run -v ${DEV_LDFLAGS} -tags=jsoniter main.go

.PHONY: default production dev test run
