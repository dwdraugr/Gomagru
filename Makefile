all:
	if [ ! -f config.json ]; then cp config.json.example config.json; fi
	go get
	go build
	@echo "All done! Start with './gomagru'"