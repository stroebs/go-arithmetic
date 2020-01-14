// Google cloud provider
// Credentials stored in environment
provider "google" {}

// GKE cluster
resource "google_container_cluster" "primary" {
  name               = "go-arithmetic-cluster"
  location           = "europe-west1"
  min_master_version = "1.14.9-gke.2"

  remove_default_node_pool = true
  initial_node_count       = 1
}

// GKE node pool
resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "go-arithmetic-node-pool"
  location   = "europe-west1"
  cluster    = "${google_container_cluster.primary.name}"
  node_count = 1
  version    = "1.14.9-gke.2"

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/pubsub",
      "https://www.googleapis.com/auth/service.management.readonly",
      "https://www.googleapis.com/auth/servicecontrol",
      "https://www.googleapis.com/auth/trace.append",
    ]
  }
}
