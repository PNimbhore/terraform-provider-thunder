provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_alg_esp" "thunder_cgnv6_fixed_nat_alg_esp" {
  sampling_enable {
    counters1 = "all"
  }
}