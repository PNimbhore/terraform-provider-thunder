---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_waf Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_waf: Debug Web Application Firewall
  PLACEHOLDER
---

# thunder_debug_waf (Resource)

`thunder_debug_waf`: Debug Web Application Firewall

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_waf" "thunder_debug_waf" {
  level = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `level` (Number) Debug level (Level 1-4)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


