---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_bfd_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_bfd_stats: Statistics for the object bfd
  PLACEHOLDER
---

# thunder_system_bfd_stats (Data Source)

`thunder_system_bfd_stats`: Statistics for the object bfd

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_bfd_stats" "thunder_system_bfd_stats" {
}
output "get_system_bfd_stats" {
  value = ["${data.thunder_system_bfd_stats.thunder_system_bfd_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `auth_failed` (Number) Authentication failures
- `auth_key_id_mismatch` (Number) Authentication key-id mismatch
- `auth_key_mismatch` (Number) Authentication key mismatch
- `auth_length_invalid` (Number) Invalid authentication length
- `auth_mismatch` (Number) Authentication mismatch
- `auth_seqnum_invalid` (Number) Invalid authentication sequence number
- `auth_type_mismatch` (Number) Authentication type mismatch
- `data_is_short` (Number) Packet data length too short
- `dest_unreachable` (Number) Destination unreachable
- `invalid_detect_mult` (Number) Invalid detect multiplier
- `invalid_multipoint` (Number) Invalid multipoint setting
- `invalid_my_disc` (Number) Invalid my descriptor
- `invalid_ttl` (Number) Invalid TTL
- `ip_checksum_error` (Number) IP packet checksum errors
- `length_too_small` (Number) Packets too small
- `local_state_admin_down` (Number) Local admin down session state
- `multihop_mismatch` (Number) Multihop session or packet mismatch
- `no_ipv6_enable` (Number) No IPv6 enable
- `other_error` (Number) Other errors
- `session_not_found` (Number) Session not found
- `udp_checksum_error` (Number) UDP packet checksum errors
- `version_mismatch` (Number) BFD version mismatch

