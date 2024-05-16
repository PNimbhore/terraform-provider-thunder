provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_smpp_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_smpp_trigger_stats_inc" {

  name                             = "test"
  msg_proxy_client_fail            = 1
  msg_proxy_fail_start_server_conn = 1
  msg_proxy_server_fail            = 1
  select_client_fail               = 1
  select_server_fail               = 1
}