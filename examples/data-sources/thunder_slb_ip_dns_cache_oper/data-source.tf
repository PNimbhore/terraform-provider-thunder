provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_ip_dns_cache_oper" "thunder_slb_ip_dns_cache_oper" {
}
output "get_slb_ip_dns_cache_oper" {
  value = ["${data.thunder_slb_ip_dns_cache_oper.thunder_slb_ip_dns_cache_oper}"]
}