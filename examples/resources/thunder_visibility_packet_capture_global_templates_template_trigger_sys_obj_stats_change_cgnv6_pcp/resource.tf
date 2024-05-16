provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_pcp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_pcp" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by   = 5
    duration                = 60
    pkt_not_request_drop    = 1
    pkt_too_short_drop      = 1
    noroute_drop            = 1
    unsupported_version     = 1
    not_authorized          = 1
    malform_request         = 1
    unsupp_opcode           = 1
    unsupp_option           = 1
    malform_option          = 1
    no_resources            = 1
    unsupp_protocol         = 1
    cannot_provide_suggest  = 1
    address_mismatch        = 1
    excessive_remote_peers  = 1
    pkt_not_from_nat_inside = 1
    l4_process_error        = 1
    internal_error_drop     = 1
    unsol_ance_sent_fail    = 1
  }
}