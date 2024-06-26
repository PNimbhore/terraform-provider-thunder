---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_ddos_protection_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_ddos_protection_stats: Statistics for the object ddos-protection
  PLACEHOLDER
---

# thunder_fw_ddos_protection_stats (Data Source)

`thunder_fw_ddos_protection_stats`: Statistics for the object ddos-protection

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_fw_ddos_protection_stats" "thunder_fw_ddos_protection_stats" {

}
output "get_fw_ddos_protection_stats" {
  value = ["${data.thunder_fw_ddos_protection_stats.thunder_fw_ddos_protection_stats}"]
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

- `ddos_entries_too_many` (Number) Too many DDOS entries
- `ddos_entry_add_to_bgp_failure` (Number) DDoS Entry BGP add failures
- `ddos_entry_added` (Number) DDOS entry added
- `ddos_entry_added_to_bgp` (Number) DDoS Entry added to BGP
- `ddos_entry_remove_from_bgp_failure` (Number) DDOS entry BGP remove failures
- `ddos_entry_removed` (Number) DDOS entry removed
- `ddos_entry_removed_from_bgp` (Number) DDoS Entry Removed from BGP
- `ddos_packet_dropped` (Number) DDOS Packet Drop


