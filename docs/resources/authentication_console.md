---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_authentication_console Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_authentication_console: Configure console authentication type
  PLACEHOLDER
---

# thunder_authentication_console (Resource)

`thunder_authentication_console`: Configure console authentication type

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_authentication_console" "thunder_authentication_console" {
  type_cfg {
    type = 0
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `type_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--type_cfg))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--type_cfg"></a>
### Nested Schema for `type_cfg`

Optional:

- `console_type` (String)
- `type` (Number) The login authentication type


