.PHONY: test test-verbose test-coverage clean init help

# 初始化项目
init:
	@echo "初始化测试项目..."
	@go mod tidy
	@cp .env.example .env
	@echo "请编辑 .env 文件，填写正确的配置"
	@echo "初始化完成"

# 运行所有测试
test:
	@echo "运行API测试..."
	@go test -v ./cases/...

# 运行测试（详细输出）
test-verbose:
	@echo "运行API测试（详细模式）..."
	@go test -v -count=1 ./cases/...

# 生成覆盖率报告
test-coverage:
	@echo "生成测试覆盖率报告..."
	@mkdir -p reports
	@go test -coverprofile=reports/coverage.out ./cases/...
	@go tool cover -html=reports/coverage.out -o reports/coverage.html
	@echo "覆盖率报告已生成: reports/coverage.html"

# 查看覆盖率
test-coverage-view: test-coverage
	@start reports/coverage.html

# 清理测试报告
clean:
	@echo "清理测试报告..."
	@rm -rf reports/*

# 运行指定模块测试
test-module:
	@echo "运行指定模块测试: $(module)"
	@go test -v ./cases/$(module)_test.go

# 帮助
help:
	@echo "可用命令:"
	@echo "  make init              - 初始化项目"
	@echo "  make test              - 运行所有测试"
	@echo "  make test-verbose      - 运行测试（详细模式）"
	@echo "  make test-coverage     - 生成覆盖率报告"
	@echo "  make test-coverage-view- 生成并查看覆盖率报告"
	@echo "  make test-module=xxx   - 运行指定模块测试"
	@echo "  make clean             - 清理测试报告"
