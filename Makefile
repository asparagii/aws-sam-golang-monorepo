COMMON_FLAGS += --no-print-directory
MAKE_FLAGS += $(COMMON_FLAGS)
CLEAN_FLAGS += $(COMMON_FLAGS)
BUILD=make build $(MAKE_FLAGS)
CLEAN=make clean $(CLEAN_FLAGS)

.PHONY: clean build storageService webService deploy

build: storageService webService

storageService:
	@cd services/storageService; echo "> storageService"; $(BUILD)

webService:
	@cd services/webService; echo "> webService"; $(BUILD)

clean:
	@cd services/storageService; $(CLEAN)
	@cd services/webService; $(CLEAN)

deploy: build
	bin/deploy.sh
