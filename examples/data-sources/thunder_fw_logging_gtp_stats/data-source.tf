provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_logging_gtp_stats" "thunder_fw_logging_gtp_stats" {

}
output "get_fw_logging_gtp_stats" {
  value = ["${data.thunder_fw_logging_gtp_stats.thunder_fw_logging_gtp_stats}"]
}