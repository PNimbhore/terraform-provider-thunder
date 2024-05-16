provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_resource_usage_oper" "thunder_fw_resource_usage_oper" {

}
output "get_fw_resource_usage_oper" {
  value = ["${data.thunder_fw_resource_usage_oper.thunder_fw_resource_usage_oper}"]
}