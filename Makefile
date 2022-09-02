default: generate-docs

clean-docs:
	rm -rf docs/

generate-docs: clean-docs
	mkdir -p docs
	go run ./ docs ./docs