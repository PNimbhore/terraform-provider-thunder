---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_map_encapsulation_domain Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_map_encapsulation_domain: MAP Encapsulation domain
  PLACEHOLDER
---

# thunder_cgnv6_map_encapsulation_domain (Resource)

`thunder_cgnv6_map_encapsulation_domain`: MAP Encapsulation domain

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_encapsulation_domain" "thunder_cgnv6_map_encapsulation_domain" {
  name = "test"
  basic_mapping_rule {
    rule_ipv4_address_port_settings = "prefix-addr"
    ea_length                       = 14
  }
  description = "201"
  format      = "draft-03"
  health_check_gateway {
    address_list {
      ipv4_gateway = "10.10.10.10"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
  tunnel_endpoint_address = "2001:db8:3333:4444:5555:6666:7777:8888"
  user_tag                = "71"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) MAP-E domain name

### Optional

- `basic_mapping_rule` (Block List, Max: 1) (see [below for nested schema](#nestedblock--basic_mapping_rule))
- `description` (String) MAP-E domain description
- `format` (String) 'draft-03': Construct IPv6 Interface Identifier according to draft-03;
- `health_check_gateway` (Block List, Max: 1) (see [below for nested schema](#nestedblock--health_check_gateway))
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `tunnel_endpoint_address` (String) Tunnel Endpoint Address for MAP-E domain
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--basic_mapping_rule"></a>
### Nested Schema for `basic_mapping_rule`

Optional:

- `ea_length` (Number) Length of Embedded Address (EA) bits
- `port_start` (Number) Starting Port, Must be Power of 2 value or zero
- `prefix_rule_list` (Block List) (see [below for nested schema](#nestedblock--basic_mapping_rule--prefix_rule_list))
- `rule_ipv4_address_port_settings` (String) 'prefix-addr': Each CE is assigned an IPv4 prefix; 'single-addr': Each CE is assigned an IPv4 address; 'shared-addr': Each CE is assigned a shared IPv4 address;
- `share_ratio` (Number) Port sharing ratio for each NAT IP. Must be Power of 2 value
- `uuid` (String) uuid of the object

<a id="nestedblock--basic_mapping_rule--prefix_rule_list"></a>
### Nested Schema for `basic_mapping_rule.prefix_rule_list`

Required:

- `name` (String) MAP BMR prefix rule name

Optional:

- `ea_length` (Number) Length of Embedded Address (EA) bits
- `ipv4_address_port_settings` (String) 'prefix-addr': Each CE is assigned an IPv4 prefix; 'single-addr': Each CE is assigned an IPv4 address; 'shared-addr': Each CE is assigned a shared IPv4 address;
- `ipv4_netmask` (String) Subnet mask (subnet bigger than /8 is not allowed)
- `port_start` (Number) Starting Port, Must be Power of 2 value or zero
- `rule_ipv4_prefix` (String) IPv4 prefix of BMR
- `rule_ipv6_prefix` (String) IPv6 prefix of BMR
- `share_ratio` (Number) Port sharing ratio for each NAT IP. Must be Power of 2 value
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object



<a id="nestedblock--health_check_gateway"></a>
### Nested Schema for `health_check_gateway`

Optional:

- `address_list` (Block List) (see [below for nested schema](#nestedblock--health_check_gateway--address_list))
- `ipv6_address_list` (Block List) (see [below for nested schema](#nestedblock--health_check_gateway--ipv6_address_list))
- `uuid` (String) uuid of the object
- `withdraw_route` (String) 'all-link-failure': Withdraw routes on health-check failure of all IPv4 gateways or all IPv6 gateways; 'any-link-failure': Withdraw routes on health-check failure of any gateway (default);

<a id="nestedblock--health_check_gateway--address_list"></a>
### Nested Schema for `health_check_gateway.address_list`

Optional:

- `ipv4_gateway` (String) IPv4 Gateway


<a id="nestedblock--health_check_gateway--ipv6_address_list"></a>
### Nested Schema for `health_check_gateway.ipv6_address_list`

Optional:

- `ipv6_gateway` (String) IPv6 Gateway



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'inbound_packet_received': Inbound IPv4 Packets Received; 'inbound_frag_packet_received': Inbound IPv4 Fragment Packets Received; 'inbound_addr_port_validation_failed': Inbound IPv4 Destination Address Port Validation Failed; 'inbound_rev_lookup_failed': Inbound IPv4 Reverse Route Lookup Failed; 'inbound_dest_unreachable': Inbound IPv6 Destination Address Unreachable; 'outbound_packet_received': Outbound IPv6 Packets Received; 'outbound_frag_packet_received': Outbound IPv6 Fragment Packets Received; 'outbound_addr_validation_failed': Outbound IPv6 Source Address Validation Failed; 'outbound_rev_lookup_failed': Outbound IPv6 Reverse Route Lookup Failed; 'outbound_dest_unreachable': Outbound IPv4 Destination Address Unreachable; 'packet_mtu_exceeded': Packet Exceeded MTU; 'frag_icmp_sent': ICMP Packet Too Big Sent; 'interface_not_configured': Interfaces not Configured Dropped; 'bmr_prefixrules_configured': BMR prefix rules configured; 'helper_count': Helper Count; 'active_dhcpv6_leases': Active DHCPv6 leases;

