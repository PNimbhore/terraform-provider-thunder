provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_ddos_protection_l4_entries_oper" "thunder_cgnv6_ddos_protection_l4_entries_oper" {
}
output "get_cgnv6_ddos_protection_l4_entries_oper" {
  value = ["${data.thunder_cgnv6_ddos_protection_l4_entries_oper.thunder_cgnv6_ddos_protection_l4_entries_oper}"]
}