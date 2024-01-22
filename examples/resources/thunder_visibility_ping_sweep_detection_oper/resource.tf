provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_ping_sweep_detection_oper" "thunder_visibility_ping_sweep_detection_oper" {

}
output "get_visibility_ping_sweep_detection_oper" {
  value = ["${data.thunder_visibility_ping_sweep_detection_oper.thunder_visibility_ping_sweep_detection_oper}"]
}