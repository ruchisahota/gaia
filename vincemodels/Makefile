init: codegen
default: codegen

codegen: codegen_v1

codegen_v1:
	@echo '* Generating Vince models v1'
	@rm -rf codegen
	@elegen folder -d v1/specs
	@rm -f ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	@rm -rf codegen

publish:
	@git pull
	@make codegen
	@git commit -am "codegen"
	@git push
