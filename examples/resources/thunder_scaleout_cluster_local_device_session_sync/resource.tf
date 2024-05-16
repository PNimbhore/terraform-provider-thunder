provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device_session_sync" "thunder_scaleout_cluster_local_device_session_sync" {

  cluster_id    = 2
  follow_shared = 0
  reachability_options {
    skip_default_route = 1
  }
}