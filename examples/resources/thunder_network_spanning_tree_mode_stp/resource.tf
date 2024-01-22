provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_spanning_tree_mode_stp" "thunder_network_spanning_tree_mode_stp" {
  action     = 1
  priority   = 32768
  vlan_start = 426
}