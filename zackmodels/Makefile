init: codegen
default: codegen

codegen:
	@make codegen_v1
	@make codegen_v2

codegen_v1:
	@echo '* Generating Zack models v1'
	@rm -rf codegen
	@elegen folder -d v1/specs
	@rm -f ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	@rm -rf codegen

codegen_v2:
	@echo '* Generating Zack models v2'
	@rm -rf codegen
	@elegen folder -d v2/specs
	@rm -f ./v2/golang/*.go && cp codegen/elemental/*.go ./v2/golang
	@rm -rf codegen

publish:
	@git pull
	@make codegen
	@git commit -am "codegen"
	@git push
