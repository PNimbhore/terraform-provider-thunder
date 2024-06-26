---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_sport_rate_limit Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_sport_rate_limit: Configure source port rate-limit
  PLACEHOLDER
---

# thunder_slb_sport_rate_limit (Resource)

`thunder_slb_sport_rate_limit`: Configure source port rate-limit

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_sport_rate_limit" "thunder_slb_sport_rate_limit" {
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

- `counters1` (String) 'all': all; 'alloc_sport': Alloc'd src port entry; 'alloc_sportip': Alloc'd src port-ip entry; 'freed_sport': Freed src port entry; 'freed_sportip': Freed src port-ip entry; 'total_drop': Total rate exceed drop; 'total_reset': Total rate exceed reset; 'total_log': Total log sent;


