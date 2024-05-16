provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_one_to_one_shared_pool_group_members_oper" "thunder_cgnv6_one_to_one_shared_pool_group_members_oper" {
}
output "get_cgnv6_one_to_one_shared_pool_group_members_oper" {
  value = ["${data.thunder_cgnv6_one_to_one_shared_pool_group_members_oper.thunder_cgnv6_one_to_one_shared_pool_group_members_oper}"]
}