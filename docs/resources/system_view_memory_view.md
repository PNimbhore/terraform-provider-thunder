---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_view_memory_view Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_view_memory_view: Configure System Parameters
  PLACEHOLDER
---

# thunder_system_view_memory_view (Resource)

`thunder_system_view_memory_view`: Configure System Parameters

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_view_memory_view" "thunder_system_view_memory_view" {
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

- `counters1` (String) 'all': all; 'usage-percentage': Usage percentage;


