provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_hw_compress_stats" "thunder_slb_hw_compress_stats" {
}
output "get_slb_hw_compress_stats" {
  value = ["${data.thunder_slb_hw_compress_stats.thunder_slb_hw_compress_stats}"]
}