provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_lif_stats" "thunder_interface_lif_stats" {

  ifname = "test"
}
output "get_interface_lif_stats" {
  value = ["${data.thunder_interface_lif_stats.thunder_interface_lif_stats}"]
}