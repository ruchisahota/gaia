init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen:
	monogen -f specs -L elemental
	monogen -f specs -L pyelemental
	rm -f go/*.go && rm -rf python/*.py && cp codegen/elemental/1.0/*.go go && cp codegen/pyelemental/1.0/*.py python
	rm -rf codegen
