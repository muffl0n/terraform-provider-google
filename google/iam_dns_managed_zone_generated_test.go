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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDNSManagedZoneIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZoneIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dns_managed_zone_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/managedZones/%s roles/viewer", GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-zone-googlecloudexample%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDNSManagedZoneIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dns_managed_zone_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/managedZones/%s roles/viewer", GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-zone-googlecloudexample%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDNSManagedZoneIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDNSManagedZoneIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dns_managed_zone_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/managedZones/%s roles/viewer user:admin@hashicorptest.com", GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-zone-googlecloudexample%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDNSManagedZoneIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZoneIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_dns_managed_zone_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/managedZones/%s", GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-zone-googlecloudexample%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDNSManagedZoneIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dns_managed_zone_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/managedZones/%s", GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-zone-googlecloudexample%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZoneIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.net."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}

resource "google_dns_managed_zone_iam_member" "foo" {
  project = google_dns_managed_zone.default.project
  managed_zone = google_dns_managed_zone.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDNSManagedZoneIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.net."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dns_managed_zone_iam_policy" "foo" {
  project = google_dns_managed_zone.default.project
  managed_zone = google_dns_managed_zone.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDNSManagedZoneIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.net."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}

data "google_iam_policy" "foo" {
}

resource "google_dns_managed_zone_iam_policy" "foo" {
  project = google_dns_managed_zone.default.project
  managed_zone = google_dns_managed_zone.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDNSManagedZoneIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.net."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}

resource "google_dns_managed_zone_iam_binding" "foo" {
  project = google_dns_managed_zone.default.project
  managed_zone = google_dns_managed_zone.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDNSManagedZoneIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.net."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}

resource "google_dns_managed_zone_iam_binding" "foo" {
  project = google_dns_managed_zone.default.project
  managed_zone = google_dns_managed_zone.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
