---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_dns64_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_dns64_stats: Statistics for the object dns64
  PLACEHOLDER
---

# thunder_cgnv6_dns64_stats (Data Source)

`thunder_cgnv6_dns64_stats`: Statistics for the object dns64

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_dns64_stats" "thunder_cgnv6_dns64_stats" {
}
output "get_cgnv6_dns64_stats" {
  value = ["${data.thunder_cgnv6_dns64_stats.thunder_cgnv6_dns64_stats}"]
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

- `adjust` (Number) Translated
- `cache` (Number) Cache
- `drop` (Number) Dropped
- `query` (Number) Query
- `query_bad_pkt` (Number) Query Bad Packet
- `query_chg` (Number) Query Changed
- `query_parallel` (Number) Query Parallel
- `query_passive` (Number) Query Passive
- `resp` (Number) Response
- `resp_bad_pkt` (Number) Response Bad Packet
- `resp_bad_qr` (Number) Response Bad Query
- `resp_chg` (Number) Response Changed
- `resp_empty` (Number) Response Empty
- `resp_err` (Number) Response Error
- `resp_local` (Number) Response Local

