---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_apps_skip_mac_overwrite Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_apps_skip_mac_overwrite: Skips overwriting dest MAC of flooded packets on Active node
  PLACEHOLDER
---

# thunder_scaleout_apps_skip_mac_overwrite (Resource)

`thunder_scaleout_apps_skip_mac_overwrite`: Skips overwriting dest MAC of flooded packets on Active node

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_apps_skip_mac_overwrite" "thunder_scaleout_apps_skip_mac_overwrite" {
  enable = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enable` (Number) Skips overwriting dest MAC of flooded packets on Active node
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


