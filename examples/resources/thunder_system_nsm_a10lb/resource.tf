provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_nsm_a10lb" "thunder_system_nsm_a10lb" {
  kill = 0
}