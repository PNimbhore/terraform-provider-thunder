provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_fragmentation_outbound" "thunder_cgnv6_nat64_fragmentation_outbound" {

  frag_action = "drop"

}