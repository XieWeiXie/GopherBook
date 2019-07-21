BINARY=votes
VERSION=1.0.0
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Env=production -s -w"
DEV_LDFLAGS=-ldflags "-X main.Env=dev"
TEST_LDFLAGS=-ldflags "-X main.Env=test"

default:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

install:
	go mod vendor

run:
	go run main.go

prod:
	go build -o ${BINARY} -v ${LDFLAGS} -tags=jsoniter

beta:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  default install run prod beta clean
