PROJECT="fina"

default:
	go build -o ${PROJECT} main.go
	@chmod +x ${PROJECT}

run: default
	@./${PROJECT}

sync: default
	@./${PROJECT} sync

import: default
	@./${PROJECT} import

install:
	go mod download

.PHONY: default run sync import install
	
