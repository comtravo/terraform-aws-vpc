// +build localstack

package test

import (
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCApplyEnabled(t *testing.T) {
	t.Parallel()

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           "vpc_enabled",
		"subdomain":          "foo.bar.baz",
		"cidr":               "10.10.0.0/16",
		"azs":                []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"nat_az_number":      1,
		"environment":        "vpc_enabled",
		"replication_factor": 3,
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform_apply_output := terraform.InitAndApply(t, terraformOptions)
	assert.Contains(t, terraform_apply_output, "Apply complete! Resources: 25 added, 0 changed, 0 destroyed.")
}

func TestVPCApplyDisabled(t *testing.T) {
	t.Parallel()

	terraformModuleVars := map[string]interface{}{
		"enable":             false,
		"vpc_name":           "vpc_disabled",
		"subdomain":          "foo.bar.bazz",
		"cidr":               "10.11.0.0/16",
		"azs":                []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"nat_az_number":      1,
		"environment":        "vpc_disabled",
		"replication_factor": 3,
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform_apply_output := terraform.InitAndApply(t, terraformOptions)
	assert.Contains(t, terraform_apply_output, "Apply complete! Resources: 0 added, 0 changed, 0 destroyed.")
}

func SetupTestCase(t *testing.T, terraformModuleVars map[string]interface{}) *terraform.Options {
	testRunFolder, err := files.CopyTerraformFolderToTemp("../", t.Name())
	assert.NoError(t, err)
	t.Logf("Copied files to test folder: %s", testRunFolder)

	localstackCongigDestination := path.Join(testRunFolder, "localstack.tf")
	files.CopyFile("localstack.tf", localstackCongigDestination)
	t.Logf("Copied localstack file to: %s", localstackCongigDestination)

	terraformOptions := &terraform.Options{
		TerraformDir: testRunFolder,
		Vars:         terraformModuleVars,
	}
	return terraformOptions
}
