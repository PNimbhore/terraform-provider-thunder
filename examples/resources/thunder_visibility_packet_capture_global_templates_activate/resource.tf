provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_activate" "thunder_visibility_packet_capture_global_templates_activate" {

  template = "global_temp"
}