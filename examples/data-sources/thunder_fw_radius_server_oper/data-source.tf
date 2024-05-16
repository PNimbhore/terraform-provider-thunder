provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_radius_server_oper" "thunder_fw_radius_server_oper" {

}
output "get_fw_radius_server_oper" {
  value = ["${data.thunder_fw_radius_server_oper.thunder_fw_radius_server_oper}"]
}