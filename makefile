# Variables
PKG := ./...
REPORT_DIR := reports
TEST_OUT := $(REPORT_DIR)/test-report.out
COVERAGE_FILE := $(REPORT_DIR)/coverage.out
COVERAGE_HTML := $(REPORT_DIR)/coverage.html

# Create reports directory 
.PHONY: prepare
prepare:
	@mkdir -p $(REPORT_DIR)

# Run tests and show total coverage
.PHONY: go-test
go-test: prepare
	@echo "🧪 Running tests..."
	@go test -v -cover $(PKG) | tee $(TEST_OUT)
	@echo "✅ Tests completed"

# Generate coverage report
.PHONY: go-report
go-report: prepare
	@echo "📊 Generating coverage report..."
	@go test -coverprofile=$(COVERAGE_FILE) $(PKG)
	@go tool cover -func=$(COVERAGE_FILE) | tee $(REPORT_DIR)/coverage-summary.txt
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "✅ Report generated at $(COVERAGE_HTML)"
	@xdg-open $(COVERAGE_HTML) 2>/dev/null || open $(COVERAGE_HTML) 2>/dev/null || start $(COVERAGE_HTML) 2>/dev/null || echo "🔗 Open manually: $(COVERAGE_HTML)"