provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_app_protocol_port_udp_port_add_app_name" "thunder_ip_app_protocol_port_udp_port_add_app_name" {

  name     = 62
  port     = 54736
  user_tag = "46"
  interface {
    management = 0
    ve_cfg {
      ve_start = 252
      ve_end   = 252
    }
    eth_cfg {
      ethernet_start = 2
      ethernet_end   = 2
    }
  }
}