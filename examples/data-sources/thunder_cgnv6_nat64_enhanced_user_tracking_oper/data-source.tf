provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_nat64_enhanced_user_tracking_oper" "thunder_cgnv6_nat64_enhanced_user_tracking_oper" {
}
output "get_cgnv6_nat64_enhanced_user_tracking_oper" {
  value = ["${data.thunder_cgnv6_nat64_enhanced_user_tracking_oper.thunder_cgnv6_nat64_enhanced_user_tracking_oper}"]
}