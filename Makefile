init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen: codegen_squall codegen_zack codegen_vince codegen_midgard codegen_squall codegen_rufus codegen_yuffie codegen_barret codegen_highwind

codegen_zack:
	cd zackmodels && make codegen

codegen_vince:
	cd vincemodels && make codegen

codegen_midgard:
	cd midgardmodels && make codegen

codegen_squall:
	cd squallmodels && make codegen

codegen_rufus:
	cd rufusmodels && make codegen

codegen_yuffie:
	cd yuffiemodels && make codegen

codegen_barret:
	cd barretmodels && make codegen

codegen_highwind:
	cd highwindmodels && make codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
