---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_oauth_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_oauth_global: Global Oauth statistics
  PLACEHOLDER
---

# thunder_aam_authentication_oauth_global (Resource)

`thunder_aam_authentication_oauth_global`: Global Oauth statistics

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_oauth_global" "thunder_aam_authentication_oauth_global" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'auth-req': auth-req; 'auth-succ': auth-succ; 'auth-fail': auth-fail; 'auth-error': auth-error; 'relay-req': relay-req; 'relay-succ': relay-succ; 'relay-fail': relay-fail; 'other-error': other-error;

