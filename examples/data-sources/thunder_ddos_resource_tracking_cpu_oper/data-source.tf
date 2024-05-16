provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_resource_tracking_cpu_oper" "thunder_ddos_resource_tracking_cpu_oper" {
}
output "get_ddos_resource_tracking_cpu_oper" {
  value = ["${data.thunder_ddos_resource_tracking_cpu_oper.thunder_ddos_resource_tracking_cpu_oper}"]
}