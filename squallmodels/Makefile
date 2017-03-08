init: codegen
default: codegen

codegen:
	@echo '* Generating Squall models'
	rm -rf codegen
	monogen -f specs -L elemental
	monogen -f specs -L html
	rm -f ./golang/*.go && cp codegen/elemental/*.go ./golang
	rm -rf ./apidoc/* && cp -a codegen/html/* ./apidoc
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
