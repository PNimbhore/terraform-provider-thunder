provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_fw_service_group_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_fw_service_group_tmpl_trigger_stats_inc" {

  name                        = "test"
  server_selection_fail_reset = 1
}