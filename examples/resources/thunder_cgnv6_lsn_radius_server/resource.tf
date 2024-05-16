provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_radius_server" "thunder_cgnv6_lsn_radius_server" {

  accounting_interim_update = "ignore"
  accounting_on             = "ignore"
  accounting_start          = "append-entry"
  accounting_stop           = "delete-entry"
  attribute {
    attribute_value = "inside-ipv6-prefix"
    prefix_length   = "32"
    prefix_vendor   = 49464
    prefix_number   = 134
  }
  disable_reply = 1
  listen_port   = 65514
  sampling_enable {
    counters1 = "all"
  }
  secret        = 1
  secret_string = "121"

}