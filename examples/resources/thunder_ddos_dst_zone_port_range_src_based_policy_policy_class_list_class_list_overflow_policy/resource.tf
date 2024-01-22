provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy" "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy" {
  zone_name             = "test"
  port_range_start      = 20
  port_range_end        = 22
  protocol              = "tcp"
  src_based_policy_name = "test"
  class_list_name       = "test"
  dummy_name            = "configuration"
  glid                  = "3"
  action                = "deny"
  log_enable            = 1
  log_periodic          = 1
  user_tag              = "test"

}