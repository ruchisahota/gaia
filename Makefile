midgard_folder := midgardmodels
squall_folder := squallmodels
vince_folder := vincemodels
zack_folder := zackmodels
rufus_folder := rufusmodels

init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen: codegen_squall codegen_zack codegen_vince codegen_midgard codegen_squall codegen_rufus

codegen_zack:
	cd $(zack_folder) && make codegen
	# cd $(zack_folder)/golang && go build

codegen_vince:
	cd $(vince_folder) && make codegen
	# cd $(vince_folder)/golang && go build

codegen_midgard:
	cd $(midgard_folder) && make codegen
	# cd $(midgard_folder)/golang && go build

codegen_squall:
	cd $(squall_folder) && make codegen
	# cd $(squall_folder)/golang && go build

codegen_rufus:
	cd $(rufus_folder) && make codegen
	# cd $(rufus_folder)/golang && go build

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push

upload:
	gsutil -m rsync -d -R vincemodels/v1/apidoc gs://doc.aporeto.com/api/v1/vince
	gsutil -m rsync -d -R midgardmodels/v1/apidoc gs://doc.aporeto.com/api/v1/midgard
	gsutil -m rsync -d -R squallmodels/v1/apidoc gs://doc.aporeto.com/api/v1/squall
	gsutil -m rsync -d -R zackmodels/v1/apidoc gs://doc.aporeto.com/api/v1/zack
