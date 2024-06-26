---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_sport_rate_limit_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_sport_rate_limit_stats: Statistics for the object sport-rate-limit
  PLACEHOLDER
---

# thunder_slb_sport_rate_limit_stats (Data Source)

`thunder_slb_sport_rate_limit_stats`: Statistics for the object sport-rate-limit

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_sport_rate_limit_stats" "thunder_slb_sport_rate_limit_stats" {
}
output "get_slb_sport_rate_limit_stats" {
  value = ["${data.thunder_slb_sport_rate_limit_stats.thunder_slb_sport_rate_limit_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `alloc_sport` (Number) Alloc'd src port entry
- `alloc_sportip` (Number) Alloc'd src port-ip entry
- `freed_sport` (Number) Freed src port entry
- `freed_sportip` (Number) Freed src port-ip entry
- `total_drop` (Number) Total rate exceed drop
- `total_log` (Number) Total log sent
- `total_reset` (Number) Total rate exceed reset


