provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc" {

  name                     = "test"
  aflex_authz_fail         = 1
  authn_failure            = 1
  authz_failure            = 1
  connect_failed           = 1
  create_timer_failed      = 1
  dns_resolve_failed       = 1
  get_socket_option_failed = 1
  misses                   = 1
  open_socket_failed       = 1
}