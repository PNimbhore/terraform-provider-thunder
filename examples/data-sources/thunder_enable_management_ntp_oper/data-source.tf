provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_enable_management_ntp_oper" "thunder_enable_management_ntp_oper" {
}
output "get_enable_management_ntp_oper" {
  value = ["${data.thunder_enable_management_ntp_oper.thunder_enable_management_ntp_oper}"]
}