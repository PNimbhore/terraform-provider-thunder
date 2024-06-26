---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_relay_saml_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_relay_saml_stats: Statistics for the object saml
  PLACEHOLDER
---

# thunder_aam_authentication_relay_saml_stats (Data Source)

`thunder_aam_authentication_relay_saml_stats`: Statistics for the object saml

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_relay_saml_stats" "thunder_aam_authentication_relay_saml_stats" {

  name = "test"
}
output "get_aam_authentication_relay_saml_stats" {
  value = ["${data.thunder_aam_authentication_relay_saml_stats.thunder_aam_authentication_relay_saml_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify SAML authentication relay name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `error` (Number) Error
- `failure` (Number) Failure
- `request` (Number) Request
- `success` (Number) Success


