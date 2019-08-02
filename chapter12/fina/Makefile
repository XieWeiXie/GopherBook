PROJECT="fifa"




default:
	go build -o ${PROJECT} main.go
	@chmod +x ${PROJECT}


run: default
	@./${PROJECT}

sync: default
	@./${PROJECT} sync

import: default
	@./${PROJECT} import


.PHONY: default run sync import
	
