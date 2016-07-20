include domingo.mk

PROJECT_NAME := gaia


## Domingo

ci: domingo_contained_build

init: install_monolithe install_monolithe_plugins codegen domingo_init
test: domingo_test
release:


## Custom Tasks

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'

install_gobindata:
	go get -u github.com/jteeuwen/go-bindata/...

codegen:
	monogen -f specs -L elemental
	rm -f *.go && cp codegen/elemental/1.0/*.go .
	rm -rf codegen
