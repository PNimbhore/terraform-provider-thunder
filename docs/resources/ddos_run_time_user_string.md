---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_run_time_user_string Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_run_time_user_string: Configure run time user string
  PLACEHOLDER
---

# thunder_ddos_run_time_user_string (Resource)

`thunder_ddos_run_time_user_string`: Configure run time user string

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_run_time_user_string" "thunder_ddos_run_time_user_string" {
  value = "testing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `uuid` (String) uuid of the object
- `value` (String) Add run time user string

### Read-Only

- `id` (String) The ID of this resource.


