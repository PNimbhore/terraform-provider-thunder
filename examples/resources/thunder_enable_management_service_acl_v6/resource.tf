provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_management_service_acl_v6" "thunder_enable_management_service_acl_v6" {

  acl_name      = "testing4"
  all_data_intf = 1
  management    = 1
  user_tag      = "55"

}