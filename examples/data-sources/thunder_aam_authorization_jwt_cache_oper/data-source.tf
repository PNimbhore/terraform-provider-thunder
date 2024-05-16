provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authorization_jwt_cache_oper" "thunder_aam_authorization_jwt_cache_oper" {
}
output "get_aam_authorization_jwt_cache_oper" {
  value = ["${data.thunder_aam_authorization_jwt_cache_oper.thunder_aam_authorization_jwt_cache_oper}"]
}