build:
	$(MAKE) -C client
	mkdir -p server/src/gohome/static
	cp -r client/build/* server/src/gohome/static
.PHONY: build
