provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_ocsp" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_ocsp" {

  name = "test"
  trigger_stats_inc {
    stapling_request_dropped  = 1
    stapling_response_failure = 1
    stapling_response_error   = 1
    stapling_response_timeout = 1
    request_dropped           = 1
    response_failure          = 1
    response_error            = 1
    response_timeout          = 1
    job_start_error           = 1
    polling_control_error     = 1
  }
}