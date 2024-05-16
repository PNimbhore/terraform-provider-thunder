provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_service_group_stats" "thunder_fw_service_group_stats" {

  name = "sg1"
}
output "get_fw_service_group_stats" {
  value = ["${data.thunder_fw_service_group_stats.thunder_fw_service_group_stats}"]
}