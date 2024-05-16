provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_logging_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_logging_trigger_stats_rate" {

  name                        = "test"
  duration                    = 60
  http_logging_invalid_format = 1
  log_dropped                 = 1
  session_limit_exceeded      = 1
  threshold_exceeded_by       = 5
}