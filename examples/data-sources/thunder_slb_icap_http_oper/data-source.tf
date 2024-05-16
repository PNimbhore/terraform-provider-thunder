provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_icap_http_oper" "thunder_slb_icap_http_oper" {
}
output "get_slb_icap_http_oper" {
  value = ["${data.thunder_slb_icap_http_oper.thunder_slb_icap_http_oper}"]
}