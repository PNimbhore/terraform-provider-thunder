---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_global_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_global_stats: Statistics for the object global
  PLACEHOLDER
---

# thunder_aam_authentication_global_stats (Data Source)

`thunder_aam_authentication_global_stats`: Statistics for the object global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_global_stats" "thunder_aam_authentication_global_stats" {
}
output "get_aam_authentication_global_stats" {
  value = ["${data.thunder_aam_authentication_global_stats.thunder_aam_authentication_global_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `active_session` (Number) Total Active Auth-Sessions
- `active_user` (Number) Total Active Users
- `aflex_authz_fail` (Number) Total Authorization failure number in aFleX
- `aflex_authz_succ` (Number) Total Authorization success number in aFleX
- `auth_ctx_num` (Number) Total Auth Contexts
- `authn_failure` (Number) Total Authentication failure number
- `authn_success` (Number) Total Authentication success number
- `authz_failure` (Number) Total Authorization failure number
- `authz_success` (Number) Total Authorization success number
- `connect` (Number) Total AAM Connection
- `connect_failed` (Number) Total AAM Connect Failed
- `create_timer_failed` (Number) Total AAM Timer Creation Failed
- `created_timer` (Number) Total AAM Timer Created
- `dns_resolve_failed` (Number) Total AAM DNS resolve failed
- `domain_wlist_match` (Number) Total DOMAIN WHITELIST match number
- `domain_wlist_unmatch` (Number) Total DOMAIN WHITELIST unmatch number
- `get_socket_option_failed` (Number) Total AAM Get Socket Option Failed
- `misses` (Number) Total Authentication Request Missed
- `ocsp_stapling_requests_to_a10authd` (Number) Total OCSP Stapling Request
- `ocsp_stapling_responses_from_a10authd` (Number) Total OCSP Stapling Response
- `open_socket_failed` (Number) Total AAM Open Socket Failed
- `opened_socket` (Number) Total AAM Socket Opened
- `requests` (Number) Total Authentication Request
- `responses` (Number) Total Authentication Response
- `total_request` (Number) Total Request Received by A10 Auth Service

