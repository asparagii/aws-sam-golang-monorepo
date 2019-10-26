.PHONY: clean all storageService webService deploy

build: storageService webService

storageService:
	(cd services/storageService; make build)

webService:
	(cd services/webService; make build)

clean:
	(cd services/storageService; make clean)
	(cd services/webService; make clean)

deploy: build
	bin/deploy.sh
