
addlicense:
	# go install github.com/google/addlicense@v1.0.0
	find . -name "*.go" | xargs addlicense -c 'Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230' -l apache
