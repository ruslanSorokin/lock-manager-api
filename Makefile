.SILENT: tools.download
tools.download:
	@go mod download

.SILENT: tools.install
tools.install:
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
