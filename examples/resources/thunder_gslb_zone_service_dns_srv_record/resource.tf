provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_srv_record" "thunder_gslb_zone_service_dns_srv_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  port         = 62159
  priority     = 41546
  sampling_enable {
    counters1 = "all"
  }
  srv_name = "110"
  ttl      = 0
  weight   = 10
}