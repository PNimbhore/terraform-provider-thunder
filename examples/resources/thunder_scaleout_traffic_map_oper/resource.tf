provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_traffic_map_oper" "thunder_scaleout_traffic_map_oper" {
}
output "get_scaleout_traffic_map_oper" {
  value = ["${data.thunder_scaleout_traffic_map_oper.thunder_scaleout_traffic_map_oper}"]
}