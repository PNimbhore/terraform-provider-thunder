provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_control_cpu_stats" "thunder_system_control_cpu_stats" {
}
output "get_system_control_cpu_stats" {
  value = ["${data.thunder_system_control_cpu_stats.thunder_system_control_cpu_stats}"]
}