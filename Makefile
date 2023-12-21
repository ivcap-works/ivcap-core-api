ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.phony: all addlicense

all: addlicense
	@echo "done"

addlicense:
	# go install github.com/google/addlicense@v1.0.0
	find ${ROOT_DIR} -name "*.go" | xargs addlicense -c 'Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230' -l apache
