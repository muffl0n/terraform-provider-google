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

func TestAccComputeRoute_routeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckComputeRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRoute_routeBasicExample(context),
			},
			{
				ResourceName:            "google_compute_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "next_hop_instance", "next_hop_vpn_tunnel"},
			},
		},
	})
}

func testAccComputeRoute_routeBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_route" "default" {
  name        = "tf-test-network-route%{random_suffix}"
  dest_range  = "15.0.0.0/24"
  network     = google_compute_network.default.name
  next_hop_ip = "10.132.1.5"
  priority    = 100
}

resource "google_compute_network" "default" {
  name = "tf-test-compute-network%{random_suffix}"
}
`, context)
}

func TestAccComputeRoute_routeIlbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckComputeRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRoute_routeIlbExample(context),
			},
			{
				ResourceName:            "google_compute_route.route-ilb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "next_hop_instance", "next_hop_vpn_tunnel"},
			},
		},
	})
}

func testAccComputeRoute_routeIlbExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  name                    = "tf-test-compute-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-compute-subnet%{random_suffix}"
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

resource "google_compute_health_check" "hc" {
  name               = "tf-test-proxy-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service" "backend" {
  name          = "tf-test-compute-backend%{random_suffix}"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_forwarding_rule" "default" {
  name     = "tf-test-compute-forwarding-rule%{random_suffix}"
  region   = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_route" "route-ilb" {
  name         = "tf-test-route-ilb%{random_suffix}"
  dest_range   = "0.0.0.0/0"
  network      = google_compute_network.default.name
  next_hop_ilb = google_compute_forwarding_rule.default.id
  priority     = 2000
}
`, context)
}

func testAccCheckComputeRouteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_route" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/routes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil, isPeeringOperationInProgress)
			if err == nil {
				return fmt.Errorf("ComputeRoute still exists at %s", url)
			}
		}

		return nil
	}
}
