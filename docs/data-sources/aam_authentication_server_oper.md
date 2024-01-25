---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_server_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_server_oper: Operational Status for the object server
  PLACEHOLDER
---

# thunder_aam_authentication_server_oper (Data Source)

`thunder_aam_authentication_server_oper`: Operational Status for the object server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_server_oper" "thunder_aam_authentication_server_oper" {
}
output "get_aam_authentication_server_oper" {
  value = ["${data.thunder_aam_authentication_server_oper.thunder_aam_authentication_server_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ldap` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ldap))
- `ocsp` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ocsp))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `windows` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ldap"></a>
### Nested Schema for `ldap`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ldap--oper))

<a id="nestedblock--ldap--oper"></a>
### Nested Schema for `ldap.oper`

Optional:

- `ldaps_server_list` (Block List) (see [below for nested schema](#nestedblock--ldap--oper--ldaps_server_list))

<a id="nestedblock--ldap--oper--ldaps_server_list"></a>
### Nested Schema for `ldap.oper.ldaps_server_list`

Optional:

- `ldap_uri` (String)
- `ldaps_idle_conn_fd_list` (String)
- `ldaps_idle_conn_num` (Number)
- `ldaps_inuse_conn_fd_list` (String)
- `ldaps_inuse_conn_num` (Number)




<a id="nestedblock--ocsp"></a>
### Nested Schema for `ocsp`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ocsp--oper))

<a id="nestedblock--ocsp--oper"></a>
### Nested Schema for `ocsp.oper`

Optional:

- `name` (String)
- `stats_clear_type` (String)



<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `get_count` (String)
- `name` (String)
- `part_id` (Number)
- `rport_count` (Number)
- `rserver_count` (Number)
- `rserver_list` (Block List) (see [below for nested schema](#nestedblock--oper--rserver_list))

<a id="nestedblock--oper--rserver_list"></a>
### Nested Schema for `oper.rserver_list`

Optional:

- `hm` (String)
- `host` (String)
- `ip` (String)
- `max_conn` (Number)
- `rport_list` (Block List) (see [below for nested schema](#nestedblock--oper--rserver_list--rport_list))
- `server_name` (String)
- `status` (String)
- `weight` (Number)

<a id="nestedblock--oper--rserver_list--rport_list"></a>
### Nested Schema for `oper.rserver_list.rport_list`

Optional:

- `port` (Number)
- `port_hm` (String)
- `port_max_conn` (Number)
- `port_state` (String)
- `port_status` (String)
- `port_weight` (Number)
- `protocol` (String)
- `sg_list` (Block List) (see [below for nested schema](#nestedblock--oper--rserver_list--rport_list--sg_list))

<a id="nestedblock--oper--rserver_list--rport_list--sg_list"></a>
### Nested Schema for `oper.rserver_list.rport_list.sg_list`

Optional:

- `sg_name` (String)
- `sg_state` (String)





<a id="nestedblock--windows"></a>
### Nested Schema for `windows`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows--oper))

<a id="nestedblock--windows--oper"></a>
### Nested Schema for `windows.oper`

Optional:

- `name` (String)
- `stats_clear_type` (String)

