---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_plat_cpu_drop_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_plat_cpu_drop_oper: Operational Status for the object plat-cpu-drop
  PLACEHOLDER
---

# thunder_plat_cpu_drop_oper (Data Source)

`thunder_plat_cpu_drop_oper`: Operational Status for the object plat-cpu-drop

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_plat_cpu_drop_oper" "thunder_plat_cpu_drop_oper" {
}
output "get_plat_cpu_drop_oper" {
  value = ["${data.thunder_plat_cpu_drop_oper.thunder_plat_cpu_drop_oper}"]
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

- `drop_seg` (Block List) (see [below for nested schema](#nestedblock--oper--drop_seg))
- `fpga_seg` (Block List) (see [below for nested schema](#nestedblock--oper--fpga_seg))
- `rate_limit` (Number)
- `rate_limit_drp` (Block List) (see [below for nested schema](#nestedblock--oper--rate_limit_drp))

<a id="nestedblock--oper--drop_seg"></a>
### Nested Schema for `oper.drop_seg`

Optional:

- `drop_cnt` (Block List) (see [below for nested schema](#nestedblock--oper--drop_seg--drop_cnt))
- `drop_name` (String)

<a id="nestedblock--oper--drop_seg--drop_cnt"></a>
### Nested Schema for `oper.drop_seg.drop_cnt`

Optional:

- `drop_count` (String)



<a id="nestedblock--oper--fpga_seg"></a>
### Nested Schema for `oper.fpga_seg`

Optional:

- `fpga_seg_name` (String)


<a id="nestedblock--oper--rate_limit_drp"></a>
### Nested Schema for `oper.rate_limit_drp`

Optional:

- `rate_limit_drop` (String)


