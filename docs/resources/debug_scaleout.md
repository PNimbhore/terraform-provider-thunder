---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_scaleout Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_scaleout: Debug scaleout
  PLACEHOLDER
---

# thunder_debug_scaleout (Resource)

`thunder_debug_scaleout`: Debug scaleout

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_scaleout" "thunder_debug_scaleout" {
  config       = 0
  debug_level  = 3
  event        = 0
  packet       = 0
  session_sync = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config` (Number) Debug logs for scaleout config change
- `debug_level` (Number) Debug level (Level 1-3)
- `event` (Number) Debug logs for scaleout events
- `packet` (Number) Debug logs for scaleout packet flow
- `session_sync` (Number) Debug logs for scaleout session sync events
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

