---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_nat_nat_global_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_nat_nat_global_stats: Statistics for the object nat-global
  PLACEHOLDER
---

# thunder_ip_nat_nat_global_stats (Data Source)

`thunder_ip_nat_nat_global_stats`: Statistics for the object nat-global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_nat_global_stats" "thunder_ip_nat_nat_global_stats" {
}
output "get_ip_nat_nat_global_stats" {
  value = ["${data.thunder_ip_nat_nat_global_stats.thunder_ip_nat_nat_global_stats}"]
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


