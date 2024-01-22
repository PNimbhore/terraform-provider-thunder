provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_alg_ftp" "thunder_cgnv6_lsn_alg_ftp" {
  ftp_value = "disable"
  sampling_enable {
    counters1 = "all"
  }
}