zack_folder := "zack"
vince_folder := "vince"
midgard_folder := "midgard"

init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen: codegen_squall codegen_zack codegen_vince codegen_midgard

codegen_squall:
	@echo '* Generating Squall models'
	monogen -f specs -L elemental
	# monogen -f specs -L pyelemental
	monogen -f specs -L html
	rm -f golang/*.go && cp codegen/elemental/1.0/*.go golang
	# rm -rf python/*.py python/requirements.txt MANIFEST.in && cp codegen/pyelemental/gaia/*.py python/gaia && cp codegen/pyelemental/requirements.txt python && cp codegen/pyelemental/MANIFEST.in python && cp codegen/pyelemental/setup.py python
	rm -rf apidoc/* && cp -a codegen/html/* apidoc
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push

codegen_zack:
	cd $(zack_folder) && make codegen
	# cd $(zack_folder)/golang && go build

codegen_vince:
	cd $(vince_folder) && make codegen
	# cd $(vince_folder)/golang && go build

codegen_midgard:
	cd $(midgard_folder) && make codegen
	cd $(midgard_folder)/golang && go build
