---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_partition Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_partition: Create/unload a Network partition
  PLACEHOLDER
---

# thunder_partition (Resource)

`thunder_partition`: Create/unload a Network partition

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition" "thunder_partition" {
  partition_name = "test"
  id1            = 4
  user_tag       = "test_partition"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `partition_name` (String) Object partition name

### Optional

- `application_type` (String) 'adc': Application type ADC; 'cgnv6': Application type CGNv6;
- `id1` (Number) Specify unique Partition id
- `shared_vlan` (Block List, Max: 1) (see [below for nested schema](#nestedblock--shared_vlan))
- `template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--template))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--shared_vlan"></a>
### Nested Schema for `shared_vlan`

Optional:

- `allowable_ip_range` (Block List) (see [below for nested schema](#nestedblock--shared_vlan--allowable_ip_range))
- `allowable_ipv6_range` (Block List) (see [below for nested schema](#nestedblock--shared_vlan--allowable_ipv6_range))
- `mgmt_floating_ip_address` (String) IPv4 Address for Shared VLAN Mgmt IP Address
- `uuid` (String) uuid of the object
- `vlan` (Number)
- `vrid` (Number) Specify VRRP-A vrid

<a id="nestedblock--shared_vlan--allowable_ip_range"></a>
### Nested Schema for `shared_vlan.allowable_ip_range`


<a id="nestedblock--shared_vlan--allowable_ipv6_range"></a>
### Nested Schema for `shared_vlan.allowable_ipv6_range`



<a id="nestedblock--template"></a>
### Nested Schema for `template`

Optional:

- `resource_accounting` (String) Attach a resource accounting template (Name of the template)
- `uuid` (String) uuid of the object

