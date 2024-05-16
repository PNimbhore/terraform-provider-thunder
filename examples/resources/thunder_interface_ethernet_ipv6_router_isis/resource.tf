provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6_router_isis" "thunder_interface_ethernet_ipv6_router_isis" {
  ifnum = 2
  tag   = "108"
}