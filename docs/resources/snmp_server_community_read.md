---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_community_read Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_community_read: Define a read only community string
  PLACEHOLDER
---

# thunder_snmp_server_community_read (Resource)

`thunder_snmp_server_community_read`: Define a read only community string

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_community_read" "thunder_snmp_server_community_read" {
  oid_list {
    oid_val = "7"
    remote {
      host_list {
        dns_host = "5"
      }
      ipv4_list {
        ipv4_host = "10.10.10.10"
      }
      ipv6_list {
        ipv6_host = "2001:db8:3333:4444:5555:6666:7777:8888"
        ipv6_mask = 21
      }
    }
    user_tag = "84"
  }
  remote {
    host_list {
      dns_host = "25"
    }
    ipv4_list {
      ipv4_host = "10.10.10.10"
    }
    ipv6_list {
      ipv6_host = "2001:db8:3333:4444:5555:6666:7777:8888"
      ipv6_mask = 52
    }
  }
  user     = "2"
  user_tag = "57"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `user` (String) SNMPv1/v2c community string

### Optional

- `oid_list` (Block List) (see [below for nested schema](#nestedblock--oid_list))
- `remote` (Block List, Max: 1) (see [below for nested schema](#nestedblock--remote))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oid_list"></a>
### Nested Schema for `oid_list`

Required:

- `oid_val` (String) specific the oid (The oid value, object-key)

Optional:

- `remote` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oid_list--remote))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--oid_list--remote"></a>
### Nested Schema for `oid_list.remote`

Optional:

- `host_list` (Block List) (see [below for nested schema](#nestedblock--oid_list--remote--host_list))
- `ipv4_list` (Block List) (see [below for nested schema](#nestedblock--oid_list--remote--ipv4_list))
- `ipv6_list` (Block List) (see [below for nested schema](#nestedblock--oid_list--remote--ipv6_list))

<a id="nestedblock--oid_list--remote--host_list"></a>
### Nested Schema for `oid_list.remote.host_list`

Optional:

- `dns_host` (String) DNS remote host
- `ipv4_mask` (String) IPV4 mask


<a id="nestedblock--oid_list--remote--ipv4_list"></a>
### Nested Schema for `oid_list.remote.ipv4_list`

Optional:

- `ipv4_host` (String) IPV4 remote host
- `ipv4_mask` (String) IPV4 mask


<a id="nestedblock--oid_list--remote--ipv6_list"></a>
### Nested Schema for `oid_list.remote.ipv6_list`

Optional:

- `ipv6_host` (String) IPV6 remote host
- `ipv6_mask` (Number) IPV6 mask




<a id="nestedblock--remote"></a>
### Nested Schema for `remote`

Optional:

- `host_list` (Block List) (see [below for nested schema](#nestedblock--remote--host_list))
- `ipv4_list` (Block List) (see [below for nested schema](#nestedblock--remote--ipv4_list))
- `ipv6_list` (Block List) (see [below for nested schema](#nestedblock--remote--ipv6_list))

<a id="nestedblock--remote--host_list"></a>
### Nested Schema for `remote.host_list`

Optional:

- `dns_host` (String) DNS remote host
- `ipv4_mask` (String) IPV4 mask


<a id="nestedblock--remote--ipv4_list"></a>
### Nested Schema for `remote.ipv4_list`

Optional:

- `ipv4_host` (String) IPV4 remote host
- `ipv4_mask` (String) IPV4 mask


<a id="nestedblock--remote--ipv6_list"></a>
### Nested Schema for `remote.ipv6_list`

Optional:

- `ipv6_host` (String) IPV6 remote host
- `ipv6_mask` (Number) IPV6 mask


