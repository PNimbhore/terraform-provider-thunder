---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_web_category_statistics_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_web_category_statistics_oper: Operational Status for the object statistics
  PLACEHOLDER
---

# thunder_web_category_statistics_oper (Data Source)

`thunder_web_category_statistics_oper`: Operational Status for the object statistics

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_statistics_oper" "thunder_web_category_statistics_oper" {

}
output "get_web_category_statistics_oper" {
  value = ["${data.thunder_web_category_statistics_oper.thunder_web_category_statistics_oper}"]
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

- `clear_cache` (String)
- `num_dplane_threads` (Number)
- `num_lookup_threads` (Number)
- `per_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--per_cpu_list))
- `total_req_dropped` (Number)
- `total_req_lookup_processed` (Number)
- `total_req_processed` (Number)
- `total_req_queue` (Number)

<a id="nestedblock--oper--per_cpu_list"></a>
### Nested Schema for `oper.per_cpu_list`

Optional:

- `req_dropped` (Number)
- `req_lookup_processed` (Number)
- `req_processed` (Number)
- `req_queue` (Number)

