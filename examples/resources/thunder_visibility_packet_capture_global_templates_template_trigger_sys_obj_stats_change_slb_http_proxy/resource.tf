provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http_proxy" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http_proxy" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    parsereq_fail         = 1
    svrsel_fail           = 1
    fwdreq_fail           = 1
    fwdreqdata_fail       = 1
    snat_fail             = 1
    req_over_limit        = 1
    req_rate_over_limit   = 1
  }
}