provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_l7session_stats" "thunder_slb_l7session_stats" {
}
output "get_slb_l7session_stats" {
  value = ["${data.thunder_slb_l7session_stats.thunder_slb_l7session_stats}"]
}