// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBigqueryReservationReservation_bigqueryReservationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckBigqueryReservationReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationReservation_bigqueryReservationBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_reservation.reservation",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testAccBigqueryReservationReservation_bigqueryReservationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_reservation" "reservation" {
	name           = "tf-test-my-reservation%{random_suffix}"
	location       = "asia-northeast1"
	// Set to 0 for testing purposes
	// In reality this would be larger than zero
	slot_capacity     = 0
	ignore_idle_slots = false
	concurrency       = 0
}
`, context)
}

func testAccCheckBigqueryReservationReservationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_reservation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("BigqueryReservationReservation still exists at %s", url)
			}
		}

		return nil
	}
}
