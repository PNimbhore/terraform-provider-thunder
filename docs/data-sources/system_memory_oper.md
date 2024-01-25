---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_memory_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_memory_oper: Operational Status for the object memory
  PLACEHOLDER
---

# thunder_system_memory_oper (Data Source)

`thunder_system_memory_oper`: Operational Status for the object memory

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_memory_oper" "thunder_system_memory_oper" {
}
output "get_system_memory_oper" {
  value = ["${data.thunder_system_memory_oper.thunder_system_memory_oper}"]
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

- `aflex_memory` (Block List) (see [below for nested schema](#nestedblock--oper--aflex_memory))
- `aflex_memory_counts` (Number)
- `buffers` (Number)
- `cached` (Number)
- `free` (Number)
- `n2_memory` (Block List) (see [below for nested schema](#nestedblock--oper--n2_memory))
- `n2_memory_counts` (Number)
- `shared` (Number)
- `ssl_memory` (Block List) (see [below for nested schema](#nestedblock--oper--ssl_memory))
- `ssl_memory_counts` (Number)
- `system_memory` (Block List) (see [below for nested schema](#nestedblock--oper--system_memory))
- `system_memory_counts` (Number)
- `tcp_memory` (Block List) (see [below for nested schema](#nestedblock--oper--tcp_memory))
- `tcp_memory_counts` (Number)
- `total` (Number)
- `usage` (String)
- `used` (Number)

<a id="nestedblock--oper--aflex_memory"></a>
### Nested Schema for `oper.aflex_memory`

Optional:

- `allocated` (Number)
- `max` (Number)
- `object_size` (Number)


<a id="nestedblock--oper--n2_memory"></a>
### Nested Schema for `oper.n2_memory`

Optional:

- `allocated` (Number)
- `max` (Number)
- `object_size` (Number)


<a id="nestedblock--oper--ssl_memory"></a>
### Nested Schema for `oper.ssl_memory`

Optional:

- `allocated` (Number)
- `max` (Number)
- `object_size` (Number)


<a id="nestedblock--oper--system_memory"></a>
### Nested Schema for `oper.system_memory`

Optional:

- `allocated` (Number)
- `max` (Number)
- `object_size` (Number)


<a id="nestedblock--oper--tcp_memory"></a>
### Nested Schema for `oper.tcp_memory`

Optional:

- `allocated` (Number)
- `max` (Number)
- `object_size` (Number)

