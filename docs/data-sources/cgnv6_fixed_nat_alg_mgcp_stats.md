---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_fixed_nat_alg_mgcp_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_fixed_nat_alg_mgcp_stats: Statistics for the object mgcp
  PLACEHOLDER
---

# thunder_cgnv6_fixed_nat_alg_mgcp_stats (Data Source)

`thunder_cgnv6_fixed_nat_alg_mgcp_stats`: Statistics for the object mgcp

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_alg_mgcp_stats" "thunder_cgnv6_fixed_nat_alg_mgcp_stats" {
}
output "get_cgnv6_fixed_nat_alg_mgcp_stats" {
  value = ["${data.thunder_cgnv6_fixed_nat_alg_mgcp_stats.thunder_cgnv6_fixed_nat_alg_mgcp_stats}"]
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

- `aucx` (Number) MGCP AUCX
- `auep` (Number) MGCP AUEP
- `crcx` (Number) MGCP CRCX
- `dlcx` (Number) MGCP DLCX
- `epcf` (Number) MGCP EPCF
- `mdcx` (Number) MGCP MDCX
- `ntfy` (Number) MGCP NTFY
- `rqnt` (Number) MGCP RQNT
- `rsip` (Number) MGCP RSIP


