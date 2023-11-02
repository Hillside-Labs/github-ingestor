SRC := $(shell find -name '*.go')

build:
	docker build --no-cache -t hillsidelabs/gh-ingestor .


.PHONY: gosmee
gosmee:
	gosmee client $(WEBHOOK_URL) http://localhost:3000 

gh-ingestor: $(SRC)
	go build -o gh-ingestor ./cmd/

.PHONY: docker-run
docker-run:
	@if [ -z "$(WEBHOOK_URL)" ]; then \
        echo "please set environment variable WEBHOOK_URL."; \
        exit 1; \
	fi
	docker run -e WEBHOOK_URL=$(WEBHOOK_URL) -p 3000:3000 hillsidelabs/gh-ingestor
