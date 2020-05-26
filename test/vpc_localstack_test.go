// +build localstack

package test

import (
	"fmt"
	"path"
	"regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCApplyEnabled_basic(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "foo.bar.baz",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_vpc",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 1)
	ValidateElasticIps(t, terraformOptions, 1)
}

func TestVPCApplyEnabled_twoAvailabilityZones(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "foo.bar.baz",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_vpc",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 1)
	ValidateElasticIps(t, terraformOptions, 1)
}

func TestVPCApplyEnabled_differentSubnetConfigurations(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "foo.bar.baz",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_vpc",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 1,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 1)
	ValidateElasticIps(t, terraformOptions, 1)
}

func TestVPCApplyEnabled_noPublicSubdomain(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_vpc",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 1,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 1)
	ValidateElasticIps(t, terraformOptions, 1)
}

func TestVPCApplyEnabled_natPerAZ(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_availability_zone",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 1,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 3)
	ValidateElasticIps(t, terraformOptions, 3)
}

func TestVPCApplyEnabled_natPerAZInTwoAZ(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             true,
		"vpc_name":           vpc_name,
		"subdomain":          "",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_availability_zone",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 1,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ValidateTerraformModuleOutputs(t, terraformOptions)
	ValidateNATGateways(t, terraformOptions, 2)
	ValidateElasticIps(t, terraformOptions, 2)
}

func TestVPCApplyDisabled(t *testing.T) {
	t.Parallel()

	vpc_name := fmt.Sprintf("vpc_enabled-%s", random.UniqueId())

	terraformModuleVars := map[string]interface{}{
		"enable":             false,
		"vpc_name":           vpc_name,
		"subdomain":          "foo.bar.baz",
		"cidr":               "10.10.0.0/16",
		"availability_zones": []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		"tags": map[string]string{
			"Name": vpc_name,
		},
		"nat_gateway": map[string]string{
			"behavior": "one_nat_per_vpc",
		},
		"enable_dns_support":               true,
		"assign_generated_ipv6_cidr_block": true,
		"private_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     0,
		},
		"public_subnets": map[string]int{
			"number_of_subnets": 3,
			"newbits":           8,
			"netnum_offset":     100,
		},
	}

	terraformOptions := SetupTestCase(t, terraformModuleVars)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	terraformApplyOutput := terraform.InitAndApply(t, terraformOptions)
	assert.Contains(t, terraformApplyOutput, "Apply complete! Resources: 0 added, 0 changed, 0 destroyed.")
}

func SetupTestCase(t *testing.T, terraformModuleVars map[string]interface{}) *terraform.Options {
	testRunFolder, err := files.CopyTerraformFolderToTemp("../", t.Name())
	assert.NoError(t, err)
	t.Logf("Copied files to test folder: %s", testRunFolder)

	localstackCongigDestination := path.Join(testRunFolder, "localstack.tf")
	files.CopyFile("fixtures/localstack.tf", localstackCongigDestination)
	t.Logf("Copied localstack file to: %s", localstackCongigDestination)

	terraformOptions := &terraform.Options{
		TerraformDir: testRunFolder,
		Vars:         terraformModuleVars,
	}
	return terraformOptions
}

func ValidateTerraformModuleOutputs(t *testing.T, terraformOptions *terraform.Options) {
	ValidateVPC(t, terraformOptions)

	ValidateVPCDefaultSecurityGroup(t, terraformOptions)

	ValidateVPCRoute53ZoneID(t, terraformOptions)
	ValidateVPCRoute53ZoneName(t, terraformOptions)

	ValidateVPCRoutingTables(t, terraformOptions)

	ValidateVPCSubnets(t, terraformOptions)
	ValidateDependId(t, terraformOptions)
}

func ValidateNATGateways(t *testing.T, terraformOptions *terraform.Options, expectedNumberOfResources int) {
	nat_gateway_ids := terraform.OutputList(t, terraformOptions, "nat_gateway_ids")
	assert.Len(t, nat_gateway_ids, expectedNumberOfResources)
	ValidateEachElementInArray(t, nat_gateway_ids, "nat-*")
}

func ValidateElasticIps(t *testing.T, terraformOptions *terraform.Options, expectedNumberOfResources int) {
	elastic_ips := terraform.OutputList(t, terraformOptions, "elastic_ips")
	assert.Len(t, elastic_ips, expectedNumberOfResources)
	ValidateEachElementInArray(t, elastic_ips, "eip-*")
}

func ValidateEachElementInArray(t *testing.T, array []string, regularExpression string) {
	for _, element := range array {
		assert.Regexp(t, regularExpression, element)
	}
}

func ValidateVPCSubnets(t *testing.T, terraformOptions *terraform.Options) {
	private_subnets := terraform.OutputList(t, terraformOptions, "private_subnets")
	public_subnets := terraform.OutputList(t, terraformOptions, "public_subnets")

	assert.Len(t, private_subnets, terraformOptions.Vars["private_subnets"].(map[string]int)["number_of_subnets"])
	assert.Len(t, public_subnets, terraformOptions.Vars["public_subnets"].(map[string]int)["number_of_subnets"])
	assert.NotEqual(t, public_subnets, private_subnets)
	ValidateEachElementInArray(t, private_subnets, "subnet-*")
	ValidateEachElementInArray(t, public_subnets, "subnet-*")
}

func ValidateVPC(t *testing.T, terraformOptions *terraform.Options) {
	vpc_id := terraform.Output(t, terraformOptions, "vpc_id")
	assert.Regexp(t, "vpc-*", vpc_id)
}

func ValidateVPCDefaultSecurityGroup(t *testing.T, terraformOptions *terraform.Options) {
	vpc_default_sg := terraform.Output(t, terraformOptions, "vpc_default_sg")
	assert.Regexp(t, "sg-*", vpc_default_sg)
}

func ValidateVPCRoute53ZoneID(t *testing.T, terraformOptions *terraform.Options) {
	net0ps_zone_id := terraform.Output(t, terraformOptions, "net0ps_zone_id")
	private_zone_id := terraform.Output(t, terraformOptions, "private_zone_id")

	subdomain_zone_id := terraform.Output(t, terraformOptions, "subdomain_zone_id")
	public_subdomain_zone_id := terraform.Output(t, terraformOptions, "public_subdomain_zone_id")

	assert.NotEqual(t, "", net0ps_zone_id)
	assert.NotEqual(t, "", private_zone_id)
	assert.Equal(t, net0ps_zone_id, private_zone_id)

	publicSubdomainRegex := regexp.MustCompile("^[A-Za-z0-9,-_.\\s]+$")

	if terraformOptions.Vars["subdomain"] == "" {
		publicSubdomainRegex = regexp.MustCompile("^$")
	}

	assert.Regexp(t, publicSubdomainRegex, subdomain_zone_id)
	assert.Regexp(t, publicSubdomainRegex, public_subdomain_zone_id)
	assert.Equal(t, subdomain_zone_id, public_subdomain_zone_id)

	assert.NotEqual(t, private_zone_id, public_subdomain_zone_id)
}

func ValidateVPCRoute53ZoneName(t *testing.T, terraformOptions *terraform.Options) {
	public_subdomain := terraform.Output(t, terraformOptions, "public_subdomain")
	private_subdomain := terraform.Output(t, terraformOptions, "private_subdomain")

	assert.Equal(t, terraformOptions.Vars["subdomain"], public_subdomain)
	assert.Equal(t, fmt.Sprintf("%s-net0ps.com.", terraformOptions.Vars["vpc_name"]), private_subdomain)
}

func ValidateVPCRoutingTables(t *testing.T, terraformOptions *terraform.Options) {
	vpc_private_routing_table_id := terraform.Output(t, terraformOptions, "vpc_private_routing_table_id")
	vpc_public_routing_table_id := terraform.Output(t, terraformOptions, "vpc_public_routing_table_id")

	assert.Regexp(t, "rtb-*", vpc_private_routing_table_id)
	assert.Regexp(t, "rtb-*", vpc_public_routing_table_id)
	assert.NotEqual(t, vpc_private_routing_table_id, vpc_public_routing_table_id)
}

func ValidateDependId(t *testing.T, terraformOptions *terraform.Options) {
	depends_id := terraform.Output(t, terraformOptions, "depends_id")
	assert.NotEqual(t, "", depends_id)
}