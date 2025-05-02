#!/usr/bin/env bash
set -euo pipefail

echo "🔍 Terraform fmt & validate"
terraform fmt -check
terraform validate

echo "🧪 Go unit & integration tests"
go test ./go -timeout 30s -v

echo "🔧 Terratest"
go test ./go/terratest -timeout 5m -v

echo "📦 Install Node modules"
npm ci

echo "🚀 Playwright tests"
npx playwright test --reporter=list

echo "✅ All tests passed!"
