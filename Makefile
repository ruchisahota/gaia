init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'

codegen:
	monogen -f specs -L elemental
	rm -f *.go && cp codegen/elemental/1.0/*.go .
	rm -rf codegen
