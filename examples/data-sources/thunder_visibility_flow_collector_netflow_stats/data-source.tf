provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_flow_collector_netflow_stats" "thunder_visibility_flow_collector_netflow_stats" {
}
output "get_visibility_flow_collector_netflow_stats" {
  value = ["${data.thunder_visibility_flow_collector_netflow_stats.thunder_visibility_flow_collector_netflow_stats}"]
}