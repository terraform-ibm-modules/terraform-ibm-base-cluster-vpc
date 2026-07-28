// Tests in this file are run in the PR pipeline and the continuous testing pipeline
package test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testhelper"
)

// Use existing resource group
const resourceGroup = "geretain-test-resources"

// Ensure every example directory has a corresponding test
const openshiftExampleDir = "examples/openshift"
const kubernetesExampleDir = "examples/kubernetes"

const yamlLocation = "../common-dev-assets/common-go-assets/common-permanent-resources.yaml"

var permanentResources map[string]interface{}

func TestMain(m *testing.M) {

	var err error
	permanentResources, err = common.LoadMapFromYaml(yamlLocation)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func setupOptions(t *testing.T, prefix string, dir string) *testhelper.TestOptions {
	options := testhelper.TestOptionsDefaultWithVars(&testhelper.TestOptions{
		Testing:       t,
		TerraformDir:  dir,
		Prefix:        prefix,
		ResourceGroup: resourceGroup,
	})
	return options
}

func TestRunOCPExample(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("ocp-%s", strings.ToLower(random.UniqueID()))
	options := setupOptions(t, prefix, openshiftExampleDir)
	options.TerraformVars = map[string]interface{}{
		"cos_instance_crn": permanentResources["general_test_storage_cos_instance_crn"],
	}

	output, err := options.RunTestConsistency()
	assert.Nil(t, err, "This should not have errored")
	assert.NotNil(t, output, "Expected some output")
}

func TestRunIKSExample(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("iks-%s", strings.ToLower(random.UniqueID()))
	options := setupOptions(t, prefix, kubernetesExampleDir)

	output, err := options.RunTestConsistency()
	assert.Nil(t, err, "This should not have errored")
	assert.NotNil(t, output, "Expected some output")
}

func TestRunUpgradeOCPExample(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("ocp-upg-%s", strings.ToLower(random.UniqueID()))
	options := setupOptions(t, prefix, openshiftExampleDir)
	options.TerraformVars = map[string]interface{}{
		"cos_instance_crn": permanentResources["general_test_storage_cos_instance_crn"],
	}
	output, err := options.RunTestUpgrade()
	if !options.UpgradeTestSkipped {
		assert.Nil(t, err, "This should not have errored")
		assert.NotNil(t, output, "Expected some output")
	}
}

func TestRunUpgradeIKSExample(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("iks-upg-%s", strings.ToLower(random.UniqueID()))
	options := setupOptions(t, prefix, kubernetesExampleDir)

	output, err := options.RunTestUpgrade()
	if !options.UpgradeTestSkipped {
		assert.Nil(t, err, "This should not have errored")
		assert.NotNil(t, output, "Expected some output")
	}
}
