---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_license_manager_reminder Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_license_manager_reminder: Set the reminder for grace time
  PLACEHOLDER
---

# thunder_license_manager_reminder (Resource)

`thunder_license_manager_reminder`: Set the reminder for grace time

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_license_manager_reminder" "thunder_license_manager_reminder" {
  reminder_value = 700293
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `reminder_value` (Number) Configure reminder for grace time (Hour)

### Optional

- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


