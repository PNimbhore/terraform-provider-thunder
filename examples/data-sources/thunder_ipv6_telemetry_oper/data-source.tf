provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_telemetry_oper" "thunder_ipv6_telemetry_oper" {


}
output "get_ipv6_telemetry_oper" {
  value = ["${data.thunder_ipv6_telemetry_oper.thunder_ipv6_telemetry_oper}"]
}