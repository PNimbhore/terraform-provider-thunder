provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy_l4_type_src_dst_overflow" "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy_l4_type_src_dst_overflow" {
  dst_entry_name        = "test"
  src_based_policy_name = "test"
  class_list_name       = "test"
  dummy_name            = "configuration"
  deny                  = 1
  glid                  = "3"
  protocol              = "tcp"
  user_tag              = "test"

}