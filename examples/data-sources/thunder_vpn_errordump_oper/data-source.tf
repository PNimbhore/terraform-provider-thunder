provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_errordump_oper" "thunder_vpn_errordump_oper" {


}
output "get_vpn_errordump_oper" {
  value = ["${data.thunder_vpn_errordump_oper.thunder_vpn_errordump_oper}"]
}