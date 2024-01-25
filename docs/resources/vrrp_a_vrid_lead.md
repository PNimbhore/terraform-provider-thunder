---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_vrid_lead Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_vrid_lead: Define a vrid-lead
  PLACEHOLDER
---

# thunder_vrrp_a_vrid_lead (Resource)

`thunder_vrrp_a_vrid_lead`: Define a vrid-lead

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_vrid_lead" "thunder_vrrp_a_vrid_lead" {
  user_tag      = "test"
  vrid_lead_str = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vrid_lead_str` (String) VRRP-A VRID leader name

### Optional

- `partition` (Block List, Max: 1) (see [below for nested schema](#nestedblock--partition))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--partition"></a>
### Nested Schema for `partition`

Optional:

- `name_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--partition--name_cfg))
- `partition` (Number) Partition name
- `shared_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--partition--shared_cfg))

<a id="nestedblock--partition--name_cfg"></a>
### Nested Schema for `partition.name_cfg`

Optional:

- `name` (String) Partition name
- `vrid` (Number) VRRP-A id
- `vrid_value` (Number) Specify ha VRRP-A vrid


<a id="nestedblock--partition--shared_cfg"></a>
### Nested Schema for `partition.shared_cfg`

Optional:

- `shared` (Number) Shared partition
- `vrid` (Number) VRRP-A id
- `vrid_value` (Number) Specify ha VRRP-A vrid

