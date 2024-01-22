provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_pptp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_pptp" {

  name = "test"
  trigger_stats_inc {
    no_gre_session_match            = 1
    call_req_pns_call_id_mismatch   = 1
    call_reply_pns_call_id_mismatch = 1
  }
}