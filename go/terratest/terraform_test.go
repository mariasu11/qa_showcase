package test

import (
  "testing"
  "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformExampleModule(t *testing.T) {
  tfOpts := &terraform.Options{
    TerraformDir: "../terraform/modules/example_module",
    Vars: map[string]interface{}{
      "instance_count": 2,
    },
  }
  defer terraform.Destroy(t, tfOpts)
  terraform.InitAndApply(t, tfOpts)
}
