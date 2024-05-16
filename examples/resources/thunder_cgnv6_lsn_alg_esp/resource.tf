provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_esp" "thunder_cgnv6_lsn_alg_esp" {
  esp_value = "enable"
  sampling_enable {
    counters1 = "all"
  }
}