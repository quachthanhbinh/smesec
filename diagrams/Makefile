.PHONY: diagrams clean build

OUT := go-diagrams/out

# Generate all .dot files and render to PNG
diagrams: build
	go run cmd/logical/main.go
	go run cmd/deployment/main.go
	go run cmd/integrations/main.go
	@echo ""
	@echo "Rendering PNGs (requires Graphviz)..."
	dot -Tpng $(OUT)/logical-architecture.dot     -o $(OUT)/logical-architecture.png
	dot -Tpng $(OUT)/deployment-architecture.dot  -o $(OUT)/deployment-architecture.png
	dot -Tpng $(OUT)/integration-touchpoints.dot  -o $(OUT)/integration-touchpoints.png
	@echo "Done. Output in $(OUT)/"

# Generate .dot only (no Graphviz required)
dot:
	mkdir -p $(OUT)
	go run cmd/logical/main.go
	go run cmd/deployment/main.go
	go run cmd/integrations/main.go
	@echo "Done. .dot files in $(OUT)/"

# Build check (compile only)
build:
	go build ./...

clean:
	rm -rf $(OUT)/*.dot $(OUT)/*.png
