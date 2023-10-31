SRC := $(shell find -name '*.go')

.PHONY: link-gosmee
link-gosmee:
	gosmee client $(WEBHOOK_URL) http://localhost:3000 

gh-ingestor-server: $(SRC)
	go build .
