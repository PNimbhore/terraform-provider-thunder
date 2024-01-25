---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_saml_global_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_saml_global_stats: Statistics for the object global
  PLACEHOLDER
---

# thunder_aam_authentication_saml_global_stats (Data Source)

`thunder_aam_authentication_saml_global_stats`: Statistics for the object global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_saml_global_stats" "thunder_aam_authentication_saml_global_stats" {
}
output "get_aam_authentication_saml_global_stats" {
  value = ["${data.thunder_aam_authentication_saml_global_stats.thunder_aam_authentication_saml_global_stats}"]
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

- `acs_authz_fail` (Number) Total SAML Single-Sign-On Authorization Fail
- `acs_error` (Number) Total SAML Single-Sign-On Error
- `acs_req` (Number) Total SAML Single-Sign-On Request
- `acs_success` (Number) Total SAML Single-Sign-On Success
- `glo_slo_success` (Number) Total Global Logout Success
- `loc_slo_success` (Number) Total Local Logout Success
- `login_auth_req` (Number) Total Login Authentication Request
- `login_auth_resp` (Number) Total Login Authentication Response
- `other_error` (Number) Total Other Error
- `par_slo_success` (Number) Total Partial Logout Success
- `relay_error` (Number)
- `relay_fail` (Number)
- `relay_req` (Number)
- `relay_success` (Number)
- `requests_to_a10saml` (Number) Total Request to A10 SAML Service
- `responses_from_a10saml` (Number) Total Response from A10 SAML Service
- `slo_error` (Number) Total Single Logout Error
- `slo_req` (Number) Total Single Logout Request
- `slo_success` (Number) Total Single Logout Success
- `sp_metadata_export_req` (Number) Total Metadata Export Request
- `sp_metadata_export_success` (Number) Toal Metadata Export Success
- `sp_slo_req` (Number) Total SP-initiated Single Logout Request

