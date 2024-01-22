provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_geo_location_file_oper" "thunder_ddos_geo_location_file_oper" {
}
output "get_ddos_geo_location_file_oper" {
  value = ["${data.thunder_ddos_geo_location_file_oper.thunder_ddos_geo_location_file_oper}"]
}