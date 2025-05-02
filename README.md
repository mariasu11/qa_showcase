# QA Showcase

## Summary
This repo demonstrates cross-stack testing and automation:

- **Go**  
  - `api_health_test.go`: simple HTTP health checks  
  - `integration_test.go`: API integration tests  
  - `terratest/terraform_test.go`: Terratest for Terraform modules  

- **TypeScript**  
  - `api.spec.ts`: Playwright health checks  
  - `integration.spec.ts`: Playwright integration tests  

- **Bash**  
  - `ci_pipeline.sh`: orchestrates all tests in one command  

- **Terraform**  
  - `main.tf` + `modules/example_module`: sample infra  

- **Performance**  
  - `perf/k6_load_test.js`: k6 load test for any `/health` endpoint  

## Prerequisites
- Go 1.18+  
- Terraform 1.x  
- Node.js 16+  
- k6 (https://k6.io/)  
- AWS creds (for Terraform module tests)  

## Running Everything

# 1) Clone the repo
git clone https://github.com/youruser/qa-showcase.git
cd qa-showcase

# 2) (Optional) Point at your local API
export API_ADDR=http://127.0.0.1:8080

# 3) Run the CI script (all tests)
bash bash/ci_pipeline.sh

# 4) Run performance test
k6 run perf/k6_load_test.js
