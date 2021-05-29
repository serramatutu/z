githooks:
	sh scripts/githooks.sh

launch_debug:
	dlv debug ./cmd/z -l 127.0.0.1:2345 --headless -- $(ARGS)
