provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_waf" "thunder_debug_waf" {
  level = 1
}