provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_h323" "thunder_debug_h323" {
  dumy = 0
}