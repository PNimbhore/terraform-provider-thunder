---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vrrp_a_detail_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vrrp_a_detail_oper: Operational Status for the object detail
  PLACEHOLDER
---

# thunder_vrrp_a_detail_oper (Data Source)

`thunder_vrrp_a_detail_oper`: Operational Status for the object detail

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_detail_oper" "thunder_vrrp_a_detail_oper" {
}
output "get_vrrp_a_detail_oper" {
  value = ["${data.thunder_vrrp_a_detail_oper.thunder_vrrp_a_detail_oper}"]
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

- `bad_group_rcv` (Number)
- `dup_id_rcv` (Number)
- `err_devid` (Number)
- `err_parid` (Number)
- `err_port` (Number)
- `ip6_pools_exceeded` (Number)
- `ip_pools_exceeded` (Number)
- `l2_no_route` (Number)
- `local_info_list` (Block List) (see [below for nested schema](#nestedblock--oper--local_info_list))
- `lock_try` (Number)
- `peer_info_list` (Block List) (see [below for nested schema](#nestedblock--oper--peer_info_list))
- `set_id_mismatch` (Number)
- `time_inaccurate_count` (Number)
- `vrrp_version_mismatch` (Number)

<a id="nestedblock--oper--local_info_list"></a>
### Nested Schema for `oper.local_info_list`

Optional:

- `local_eth_send_list` (Block List) (see [below for nested schema](#nestedblock--oper--local_info_list--local_eth_send_list))
- `switch_to_active` (Number)
- `switch_to_standby` (Number)
- `vrid` (Number)
- `vrrp_pkt_send` (Number)

<a id="nestedblock--oper--local_info_list--local_eth_send_list"></a>
### Nested Schema for `oper.local_info_list.local_eth_send_list`

Optional:

- `eth_pkt_send` (Number)
- `eth_port` (Number)



<a id="nestedblock--oper--peer_info_list"></a>
### Nested Schema for `oper.peer_info_list`

Optional:

- `peer_ip` (String)
- `peer_list` (Block List) (see [below for nested schema](#nestedblock--oper--peer_info_list--peer_list))
- `peer_pkt_recv` (Number)

<a id="nestedblock--oper--peer_info_list--peer_list"></a>
### Nested Schema for `oper.peer_info_list.peer_list`

Optional:

- `missing_heartbeat` (Number)
- `peer_id` (Number)
- `peer_port_list` (Block List) (see [below for nested schema](#nestedblock--oper--peer_info_list--peer_list--peer_port_list))
- `vrid` (Number)

<a id="nestedblock--oper--peer_info_list--peer_list--peer_port_list"></a>
### Nested Schema for `oper.peer_info_list.peer_list.peer_port_list`

Optional:

- `eth` (Number)
- `eth_miss` (Number)
- `vrrp_pkt_recv` (Number)

