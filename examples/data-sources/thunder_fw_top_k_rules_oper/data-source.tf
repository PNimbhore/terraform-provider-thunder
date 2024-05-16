provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_top_k_rules_oper" "thunder_fw_top_k_rules_oper" {

}
output "get_fw_top_k_rules_oper" {
  value = ["${data.thunder_fw_top_k_rules_oper.thunder_fw_top_k_rules_oper}"]
}