---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_file_metrics Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_file_metrics: Persistent storage for metrics
  PLACEHOLDER
---

# thunder_visibility_file_metrics (Resource)

`thunder_visibility_file_metrics`: Persistent storage for metrics

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_file_metrics" "thunder_visibility_file_metrics" {
  action = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'enable': Enable persistent storage; 'disable': Disable persistent storage(default);
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


