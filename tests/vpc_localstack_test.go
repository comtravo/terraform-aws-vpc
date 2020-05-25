// +build localstack

package test

import (
	"fmt"
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

	private_subnets := terraform.OutputList(t, terraformOptions, "private_subnets")
	public_subnets := terraform.OutputList(t, terraformOptions, "public_subnets")
	assert.Len(t, private_subnets, 3)
	assert.Len(t, public_subnets, 3)
	assert.NotEqual(t, public_subnets, private_subnets)

	vpc_id := terraform.Output(t, terraformOptions, "vpc_id")
	assert.Regexp(t, "vpc-*", vpc_id)

	vpc_default_sg := terraform.Output(t, terraformOptions, "vpc_default_sg")
	assert.Regexp(t, "sg-*", vpc_default_sg)

	net0ps_zone_id := terraform.Output(t, terraformOptions, "net0ps_zone_id")
	private_zone_id := terraform.Output(t, terraformOptions, "private_zone_id")
	subdomain_zone_id := terraform.Output(t, terraformOptions, "subdomain_zone_id")
	public_subdomain_zone_id := terraform.Output(t, terraformOptions, "public_subdomain_zone_id")
	assert.NotEqual(t, "", net0ps_zone_id)
	assert.NotEqual(t, "", private_zone_id)
	assert.Equal(t, net0ps_zone_id, private_zone_id)
	assert.NotEqual(t, "", subdomain_zone_id)
	assert.NotEqual(t, "", public_subdomain_zone_id)
	assert.Equal(t, subdomain_zone_id, public_subdomain_zone_id)
	assert.NotEqual(t, private_zone_id, public_subdomain_zone_id)

	public_subdomain := terraform.Output(t, terraformOptions, "public_subdomain")
	assert.Equal(t, terraformModuleVars["subdomain"], public_subdomain)

	private_subdomain := terraform.Output(t, terraformOptions, "private_subdomain")
	assert.Equal(t, fmt.Sprintf("%s-net0ps.com.", terraformModuleVars["vpc_name"]), private_subdomain)

	vpc_private_routing_table_id := terraform.Output(t, terraformOptions, "vpc_private_routing_table_id")
	vpc_public_routing_table_id := terraform.Output(t, terraformOptions, "vpc_public_routing_table_id")
	assert.Regexp(t, "rtb-*", vpc_private_routing_table_id)
	assert.Regexp(t, "rtb-*", vpc_public_routing_table_id)
	assert.NotEqual(t, vpc_private_routing_table_id, vpc_public_routing_table_id)

	depends_id := terraform.Output(t, terraformOptions, "depends_id")
	assert.NotEqual(t, "", depends_id)

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
