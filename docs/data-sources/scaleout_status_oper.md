---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_status_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_status_oper: Operational Status for the object status
  PLACEHOLDER
---

# thunder_scaleout_status_oper (Data Source)

`thunder_scaleout_status_oper`: Operational Status for the object status

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_status_oper" "thunder_scaleout_status_oper" {
}
output "get_scaleout_status_oper" {
  value = ["${data.thunder_scaleout_status_oper.thunder_scaleout_status_oper}"]
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

- `active_interface_list` (Block List) (see [below for nested schema](#nestedblock--oper--active_interface_list))
- `advertised_redirect_ip_list` (Block List) (see [below for nested schema](#nestedblock--oper--advertised_redirect_ip_list))
- `advertised_session_sync_ip_list` (Block List) (see [below for nested schema](#nestedblock--oper--advertised_session_sync_ip_list))
- `cluster_mode` (String)
- `db_role` (String)
- `dest_redirect_ip_list` (Block List) (see [below for nested schema](#nestedblock--oper--dest_redirect_ip_list))
- `dest_session_sync_ip_list` (Block List) (see [below for nested schema](#nestedblock--oper--dest_session_sync_ip_list))
- `device_list` (Block List) (see [below for nested schema](#nestedblock--oper--device_list))
- `exclude_interface_ip_list` (Block List) (see [below for nested schema](#nestedblock--oper--exclude_interface_ip_list))
- `follow_shared_redirection` (Number)
- `follow_shared_session_sync` (Number)
- `l2redirect` (Number)
- `l2redirect_eth` (Number)
- `l2redirect_operational` (Number)
- `l2redirect_trunk` (Number)
- `l2redirect_valid` (Number)
- `l2redirect_vlan` (Number)
- `role` (String)

<a id="nestedblock--oper--active_interface_list"></a>
### Nested Schema for `oper.active_interface_list`

Optional:

- `interface` (String)


<a id="nestedblock--oper--advertised_redirect_ip_list"></a>
### Nested Schema for `oper.advertised_redirect_ip_list`

Optional:

- `ip` (String)


<a id="nestedblock--oper--advertised_session_sync_ip_list"></a>
### Nested Schema for `oper.advertised_session_sync_ip_list`

Optional:

- `ip` (String)


<a id="nestedblock--oper--dest_redirect_ip_list"></a>
### Nested Schema for `oper.dest_redirect_ip_list`

Optional:

- `device_id` (Number)
- `direction` (String)
- `ip` (String)


<a id="nestedblock--oper--dest_session_sync_ip_list"></a>
### Nested Schema for `oper.dest_session_sync_ip_list`

Optional:

- `device_id` (Number)
- `ip` (String)
- `ipv6` (String)


<a id="nestedblock--oper--device_list"></a>
### Nested Schema for `oper.device_list`

Optional:

- `id1` (Number)
- `is_local` (Number)
- `is_master` (Number)
- `state` (String)


<a id="nestedblock--oper--exclude_interface_ip_list"></a>
### Nested Schema for `oper.exclude_interface_ip_list`

Optional:

- `ip` (String)


