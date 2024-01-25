provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_bandwidth_stats" "thunder_system_bandwidth_stats" {
}
output "get_system_bandwidth_stats" {
  value = ["${data.thunder_system_bandwidth_stats.thunder_system_bandwidth_stats}"]
}