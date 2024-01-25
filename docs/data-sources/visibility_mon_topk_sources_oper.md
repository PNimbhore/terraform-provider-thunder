---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_mon_topk_sources_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_mon_topk_sources_oper: Operational Status for the object sources
  PLACEHOLDER
---

# thunder_visibility_mon_topk_sources_oper (Data Source)

`thunder_visibility_mon_topk_sources_oper`: Operational Status for the object sources

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_mon_topk_sources_oper" "thunder_visibility_mon_topk_sources_oper" {

}
output "get_visibility_mon_topk_sources_oper" {
  value = ["${data.thunder_visibility_mon_topk_sources_oper.thunder_visibility_mon_topk_sources_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `metric_topk_list` (Block List) (see [below for nested schema](#nestedblock--oper--metric_topk_list))

<a id="nestedblock--oper--metric_topk_list"></a>
### Nested Schema for `oper.metric_topk_list`

Optional:

- `metric_name` (String)
- `topk_list` (Block List) (see [below for nested schema](#nestedblock--oper--metric_topk_list--topk_list))

<a id="nestedblock--oper--metric_topk_list--topk_list"></a>
### Nested Schema for `oper.metric_topk_list.topk_list`

Optional:

- `ip_addr` (String)
- `metric_value` (String)

