---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_health_external_edit Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_health_external_edit: Edit external health monitor script
  PLACEHOLDER
---

# thunder_health_external_edit (Resource)

`thunder_health_external_edit`: Edit external health monitor script

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_external_edit" "thunder_health_external_edit" {
  description = "56"
  file_name   = "2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `description` (String) Describe health external monitor script briefly
- `file_name` (String) External health monitor script file name

### Read-Only

- `id` (String) The ID of this resource.


