provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device_traffic_redirection_interfaces" "thunder_scaleout_cluster_local_device_traffic_redirection_interfaces" {

  cluster_id = 2
  loopback_cfg {
    loopback = 5
  }
}