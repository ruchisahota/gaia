init: codegen
default: codegen

codegen:
	@echo '* Generating Midgard models'
	monogen -f specs -L elemental
	monogen -f specs -L html
	rm -f ./golang/*.go && cp codegen/elemental/1.0/*.go ./golang
	rm -rf ./apidoc/* && cp -a codegen/html/* ./apidoc
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
