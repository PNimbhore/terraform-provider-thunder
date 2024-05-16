provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_nat_list_oper" "thunder_scaleout_debug_nat_list_oper" {
}
output "get_scaleout_debug_nat_list_oper" {
  value = ["${data.thunder_scaleout_debug_nat_list_oper.thunder_scaleout_debug_nat_list_oper}"]
}