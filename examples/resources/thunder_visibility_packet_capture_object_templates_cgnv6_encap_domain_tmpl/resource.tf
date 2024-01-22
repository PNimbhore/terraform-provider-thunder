provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl" "thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by               = 5
    duration                            = 60
    inbound_addr_port_validation_failed = 1
    inbound_rev_lookup_failed           = 1
    inbound_dest_unreachable            = 1
    outbound_addr_validation_failed     = 1
    outbound_rev_lookup_failed          = 1
    outbound_dest_unreachable           = 1
    packet_mtu_exceeded                 = 1
    interface_not_configured            = 1
  }
  trigger_stats_severity {
    error          = 1
    error_alert    = 1
    error_warning  = 1
    error_critical = 1
    drop           = 1
    drop_alert     = 1
    drop_warning   = 1
    drop_critical  = 1
  }
  user_tag = "105"
}