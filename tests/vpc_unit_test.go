// +build unit

package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVPCPlanVPCEnabled(t *testing.T) {
	t.Parallel()

	terratestOptions := &terraform.Options{
		TerraformDir: "../examples/localstack/vpc_enabled/",
		Vars:         map[string]interface{}{},
	}

	plan_output := terraform.InitAndPlan(t, terratestOptions)

	assert.Contains(t, plan_output, "25 to add, 0 to change, 0 to destroy.")
}

func TestVPCPlanVPCDisabled(t *testing.T) {
	t.Parallel()

	terratestOptions := &terraform.Options{
		TerraformDir: "../examples/localstack/vpc_disabled/",
		Vars:         map[string]interface{}{},
	}

	plan_output := terraform.InitAndPlan(t, terratestOptions)

	assert.Contains(t, plan_output, "No changes. Infrastructure is up-to-date.")
}
