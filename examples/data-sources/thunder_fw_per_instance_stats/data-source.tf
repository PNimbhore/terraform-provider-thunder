provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_per_instance_stats" "thunder_fw_per_instance_stats" {

}
output "get_fw_per_instance_stats" {
  value = ["${data.thunder_fw_per_instance_stats.thunder_fw_per_instance_stats}"]
}