provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_ldap_inst_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_server_ldap_inst_tmpl_trigger_stats_rate" {

  name                  = "test"
  admin_bind_failure    = 1
  authorize_failure     = 1
  bind_failure          = 1
  duration              = 60
  other_error           = 1
  pw_change_failure     = 1
  search_failure        = 1
  ssl_session_failure   = 1
  threshold_exceeded_by = 5
  timeout_error         = 1
}