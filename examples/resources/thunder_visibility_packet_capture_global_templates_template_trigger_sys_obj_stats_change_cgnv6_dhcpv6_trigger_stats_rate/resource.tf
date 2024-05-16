provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6_trigger_stats_rate" {

  duration                  = 60
  packets_dropped           = 1
  pkts_dropped_during_clear = 1
  rcv_not_supported_msg     = 1
  threshold_exceeded_by     = 5
}