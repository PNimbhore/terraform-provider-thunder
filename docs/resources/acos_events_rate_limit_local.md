---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_events_rate_limit_local Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_events_rate_limit_local: Configure Rate Limit for Local logs
  PLACEHOLDER
---

# thunder_acos_events_rate_limit_local (Resource)

`thunder_acos_events_rate_limit_local`: Configure Rate Limit for Local logs

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_rate_limit_local" "thunder_acos_events_rate_limit_local" {
  limit = 32
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (Number) Configure Rate Limit for Local logs
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

