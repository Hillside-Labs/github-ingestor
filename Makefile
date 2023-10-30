.PHONY: link-gosmee
link-gosmee:
	gosmee client $(WEBHOOK_URL) https://localhost:8080 
