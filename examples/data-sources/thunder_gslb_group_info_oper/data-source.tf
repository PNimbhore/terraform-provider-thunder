provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_group_info_oper" "thunder_gslb_group_info_oper" {
}
output "get_gslb_group_info_oper" {
  value = ["${data.thunder_gslb_group_info_oper.thunder_gslb_group_info_oper}"]
}