init: codegen
default: codegen

codegen: codegen_v1

codegen_v1:
	@echo '* Generating Midgard models'
	@rm -rf codegen
	@elegen folder -d v1/specs
	@rm -f ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	@rm -rf codegen

publish:
	@it pull
	@ake codegen
	@it commit -am "codegen"
	@it push
