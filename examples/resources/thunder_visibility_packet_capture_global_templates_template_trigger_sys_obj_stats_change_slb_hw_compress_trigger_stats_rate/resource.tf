provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_hw_compress_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_hw_compress_trigger_stats_rate" {

  name                          = "test"
  duration                      = 60
  failure_code                  = 1
  failure_count                 = 1
  max_outstanding_request_count = 1
  max_outstanding_submit_count  = 1
  ring_full_count               = 1
  threshold_exceeded_by         = 5
}