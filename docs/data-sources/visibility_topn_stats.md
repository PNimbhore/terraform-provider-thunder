---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_topn_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_topn_stats: Statistics for the object topn
  PLACEHOLDER
---

# thunder_visibility_topn_stats (Data Source)

`thunder_visibility_topn_stats`: Statistics for the object topn

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_topn_stats" "thunder_visibility_topn_stats" {
}
output "get_visibility_topn_stats" {
  value = ["${data.thunder_visibility_topn_stats.thunder_visibility_topn_stats}"]
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

- `heap_alloc_failed` (Number) Total heap node alloc failed
- `heap_alloc_oom` (Number) Total heap node alloc failed Out of Memory
- `heap_alloc_success` (Number) Total heap node allocated
- `heap_deleted` (Number) Total Heap node deleted
- `obj_deleted` (Number) Total Object node deleted
- `obj_reg_failed` (Number) Total object node alloc failed
- `obj_reg_oom` (Number) Total object node alloc failed Out of Memory
- `obj_reg_success` (Number) Total object node allocated

