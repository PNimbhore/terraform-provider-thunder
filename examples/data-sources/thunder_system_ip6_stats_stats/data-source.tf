provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_ip6_stats_stats" "thunder_system_ip6_stats_stats" {
}
output "get_system_ip6_stats_stats" {
  value = ["${data.thunder_system_ip6_stats_stats.thunder_system_ip6_stats_stats}"]
}