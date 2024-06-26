---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_partition_shared_vlan Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_partition_shared_vlan: Enable shared-vlan feature in Network partition
  PLACEHOLDER
---

# thunder_partition_shared_vlan (Resource)

`thunder_partition_shared_vlan`: Enable shared-vlan feature in Network partition

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition_shared_vlan" "thunder_partition_shared_vlan" {

  partition_name = "test"
  vlan           = 3

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `partition_name` (String) PartitionName

### Optional

- `allowable_ip_range` (Block List) (see [below for nested schema](#nestedblock--allowable_ip_range))
- `allowable_ipv6_range` (Block List) (see [below for nested schema](#nestedblock--allowable_ipv6_range))
- `mgmt_floating_ip_address` (String) IPv4 Address for Shared VLAN Mgmt IP Address
- `uuid` (String) uuid of the object
- `vlan` (Number)
- `vrid` (Number) Specify VRRP-A vrid

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--allowable_ip_range"></a>
### Nested Schema for `allowable_ip_range`


<a id="nestedblock--allowable_ipv6_range"></a>
### Nested Schema for `allowable_ipv6_range`


