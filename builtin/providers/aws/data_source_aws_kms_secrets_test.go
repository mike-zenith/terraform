package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAwsKmsSecrets_setsOutput(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceAwsKmsSecretsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.aws_kms_secret.foosecret", "barone"),
					resource.TestCheckResourceAttrSet("data.aws_kms_secret.foosecret", "bartwo"),
				),
			},
		},
	})
}

const testAccDataSourceAwsKmsSecretsConfig = `
data "aws_kms_secrets" "foosecret" {
	secret {
		name = "barone"
		payload = "payload1"
	}
	secret {
		name = "bartwo"
		payload = "payload2"
	}
}
`
