SHELL = bash
default: plugins

.PHONY: clean-plugins
clean-plugins:
	@echo "==> Cleaning plugins..."
	@rm -rf ./bin/plugins/
	@echo "==> Done"

.PHONY: clean
clean: clean-plugins
	@echo "==> Cleaning build artifacts..."
	@rm -f ./bin/nomad-autoscaler
	@echo "==> Done"

.PHONY: bin/plugins/cron
bin/plugins/cron:
	@echo "==> Building $@..."
	@mkdir -p $$(dirname $@)
	@cd ./plugins/strategy/cron && go build -o ../../../$@
	@echo "==> Done"

.PHONY: plugins
plugins: \
	bin/plugins/cron