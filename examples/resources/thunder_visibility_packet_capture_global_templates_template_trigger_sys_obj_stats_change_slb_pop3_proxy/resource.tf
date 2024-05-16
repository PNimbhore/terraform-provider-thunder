provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_pop3_proxy" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    svrsel_fail           = 1
    no_route              = 1
    snat_fail             = 1
    line_too_long         = 1
    invalid_start_line    = 1
    unsupported_command   = 1
    bad_sequence          = 1
    rsv_persist_conn_fail = 1
    smp_v6_fail           = 1
    smp_v4_fail           = 1
    insert_tuple_fail     = 1
    cl_est_err            = 1
    ser_connecting_err    = 1
    server_response_err   = 1
  }
}