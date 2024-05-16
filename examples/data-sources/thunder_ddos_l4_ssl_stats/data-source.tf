provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_l4_ssl_stats" "thunder_ddos_l4_ssl_stats" {
}
output "get_ddos_l4_ssl_stats" {
  value = ["${data.thunder_ddos_l4_ssl_stats.thunder_ddos_l4_ssl_stats}"]
}