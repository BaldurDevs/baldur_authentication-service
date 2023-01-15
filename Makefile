AUTH_APP=authApp

build:
	@echo "Building auth binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_APP} ./cmd
	@echo "Done!"

remove_binary:
	@echo "Removing binary file"
	rm ${AUTH_APP} || true
