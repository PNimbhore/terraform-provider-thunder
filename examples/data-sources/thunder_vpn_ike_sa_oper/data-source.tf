provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_ike_sa_oper" "thunder_vpn_ike_sa_oper" {


}
output "get_vpn_ike_sa_oper" {
  value = ["${data.thunder_vpn_ike_sa_oper.thunder_vpn_ike_sa_oper}"]
}