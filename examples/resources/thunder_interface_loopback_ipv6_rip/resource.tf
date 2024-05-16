provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ipv6_rip" "thunder_interface_loopback_ipv6_rip" {

  ifnum = 4
  split_horizon_cfg {
    state = "disable"
  }
}