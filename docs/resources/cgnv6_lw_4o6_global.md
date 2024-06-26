---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_lw_4o6_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_lw_4o6_global: Configure LW-4over6 parameters
  PLACEHOLDER
---

# thunder_cgnv6_lw_4o6_global (Resource)

`thunder_cgnv6_lw_4o6_global`: Configure LW-4over6 parameters

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lw_4o6_global" "thunder_cgnv6_lw_4o6_global" {

  hairpinning = "filter-all"
  no_forward_match {
    send_icmpv6 = 1
  }
  no_reverse_match {
    send_icmp = 1
  }
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `hairpinning` (String) 'filter-all': Disable all Hairpinning; 'filter-none': Allow all Hairpinning (default); 'filter-self-ip': Block Hairpinning to same IP; 'filter-self-ip-port': Block hairpinning to same IP and Port combination;
- `icmp_inbound` (String) 'drop': Drop Inbound ICMP packets; 'handle': Handle Inbound ICMP packets(default);
- `inside_src_access_list` (Number) Access List for inside IPv4 addresses (ACL ID)
- `nat_prefix_list` (String) Configure LW-4over6 NAT Prefix List (LW-4over6 NAT Prefix Class-list)
- `no_forward_match` (Block List, Max: 1) (see [below for nested schema](#nestedblock--no_forward_match))
- `no_reverse_match` (Block List, Max: 1) (see [below for nested schema](#nestedblock--no_reverse_match))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `use_binding_table` (String) Bind LW-4over6 binding table for use (LW-4over6 Binding Table Name)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--no_forward_match"></a>
### Nested Schema for `no_forward_match`

Optional:

- `send_icmpv6` (Number) Send ICMPv6 Type 1 Code 5


<a id="nestedblock--no_reverse_match"></a>
### Nested Schema for `no_reverse_match`

Optional:

- `send_icmp` (Number) Send ICMP Type 3 Code 1


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'entry_count': Total Entries Configured; 'self_hairpinning_drop': Self-Hairpinning Drops; 'all_hairpinning_drop': All Hairpinning Drops; 'no_match_icmpv6_sent': No-Forward-Match ICMPv6 Sent; 'no_match_icmp_sent': No-Reverse-Match ICMP Sent; 'icmp_inbound_drop': Inbound ICMP Drops; 'fwd_lookup_failed': Forward Route Lookup Failed; 'rev_lookup_failed': Reverse Route Lookup Failed; 'interface_not_configured': LW-4over6 Interfaces not Configured Drops; 'no_binding_table_matches_fwd': No Forward Binding Table Entry Match Drops; 'no_binding_table_matches_rev': No Reverse Binding Table Entry Match Drops; 'session_count': LW-4over6 Session Count; 'system_address_drop': LW-4over6 System Address Drops;


