bin:
	sh -c "./build_local.sh"

deploy: bin
	cp -rf ./bin/* ~/work/gamex_dev/trunk/toolkits_bin/bin/
	cp -rf ./bin/* ~/work/gamex_dev/trunk/toolkits_bin/ops/bin/

all: deploy

.PHONY: bin deploy all

