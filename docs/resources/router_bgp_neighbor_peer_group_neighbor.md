---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_neighbor_peer_group_neighbor Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_router_bgp_neighbor_peer_group_neighbor: Specify a peer-group neighbor router
  PLACEHOLDER
---

# thunder_router_bgp_neighbor_peer_group_neighbor (Resource)

`thunder_router_bgp_neighbor_peer_group_neighbor`: Specify a peer-group neighbor router

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_neighbor_peer_group_neighbor" "BgpNeighborPeerGroupNei" {
  as_number                  = "300"
  peer_group                 = "A10"
  peer_group_key             = 1
  peer_group_remote_as       = 400
  activate                   = 1
  advertisement_interval     = 10
  allowas_in                 = 1
  allowas_in_count           = 10
  as_origination_interval    = 20
  dynamic                    = 1
  route_refresh              = 1
  extended_nexthop           = 0
  collide_established        = 1
  default_originate          = 1
  route_map                  = "Ipv6RouteMap"
  description                = "PEERGroup"
  disallow_infinite_holdtime = 1
  distribute_lists {
    distribute_list           = "DistIpv6BGP"
    distribute_list_direction = "in"
  }
  dont_capability_negotiate = 1
  ebgp_multihop             = 1
  ebgp_multihop_hop_count   = 100
  enforce_multihop          = 1
  bfd                       = 1
  multihop                  = 1
  neighbor_filter_lists {
    filter_list           = "FilterBgp"
    filter_list_direction = "in"
  }
  maximum_prefix       = 100
  maximum_prefix_thres = 80
  next_hop_self        = 1
  override_capability  = 0
  pass_value           = "PassPeergroup"
  passive              = 1

  remove_private_as = 1
  neighbor_route_map_lists {
    nbr_route_map      = "Ipv6RouteMap"
    nbr_rmap_direction = "in"
  }
  send_community_val      = "standard"
  inbound                 = 1
  shutdown                = 1
  strict_capability_match = 0
  timers_holdtime         = 200
  timers_keepalive        = 200
  connect                 = 1
  unsuppress_map          = "UnsuppIpv6"
  update_source_ip        = "1.1.1.1"
  weight                  = 300
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) AsNumber
- `peer_group` (String) Neighbor tag

### Optional

- `activate` (Number) Enable the Address Family for this Neighbor
- `advertisement_interval` (Number) Minimum interval between sending BGP routing updates (time in seconds)
- `allowas_in` (Number) Accept as-path with my AS present in it
- `allowas_in_count` (Number) Number of occurrences of AS number
- `as_origination_interval` (Number) Minimum interval between sending AS-origination routing updates (time in seconds)
- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `collide_established` (Number) Include Neighbor in Established State for Collision Detection
- `connect` (Number) BGP connect timer
- `default_originate` (Number) Originate default route to this neighbor
- `description` (String) Neighbor specific description (Up to 80 characters describing this neighbor)
- `dont_capability_negotiate` (Number) Do not perform capability negotiation
- `dynamic` (Number) Advertise dynamic capability to this neighbor
- `ebgp_multihop` (Number) Allow EBGP neighbors not on directly connected networks
- `ebgp_multihop_hop_count` (Number) maximum hop count
- `enforce_multihop` (Number) Enforce EBGP neighbors to perform multihop
- `ethernet` (Number) Ethernet interface (Port number)
- `extended_nexthop` (Number) Advertise extended-nexthop capability to this neighbor
- `inbound` (Number) Allow inbound soft reconfiguration for this neighbor
- `lif` (String) Logical interface (Lif interface name)
- `loopback` (Number) Loopback interface (Port number)
- `maximum_prefix` (Number) Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
- `maximum_prefix_thres` (Number) threshold-value, 1 to 100 percent
- `multihop` (Number) Enable multihop
- `neighbor_route_map_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor_route_map_lists))
- `override_capability` (Number) Override capability negotiation result
- `pass_value` (String) Key String
- `passive` (Number) Don't send open messages to this neighbor
- `peer_group_key` (Number) Configure peer-group
- `peer_group_remote_as` (String) Specify AS number of BGP neighbor
- `remove_private_as` (Number) Remove private AS number from outbound updates
- `route_map` (String) Route-map to specify criteria to originate default (route-map name)
- `route_refresh` (Number) Advertise route-refresh capability to this neighbor
- `shutdown` (Number) Administratively shut down this neighbor
- `strict_capability_match` (Number) Strict capability negotiation match
- `timers_holdtime` (Number) Holdtime
- `timers_keepalive` (Number) Keepalive interval
- `trunk` (Number) Trunk interface (Trunk interface number)
- `tunnel` (Number) Tunnel interface (Tunnel interface number)
- `update_source_ip` (String) IP address
- `update_source_ipv6` (String) IPv6 address
- `uuid` (String) uuid of the object
- `ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)
- `weight` (Number) Set default weight for routes from this neighbor

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--neighbor_route_map_lists"></a>
### Nested Schema for `neighbor_route_map_lists`

Optional:

- `nbr_rmap_direction` (String) 'in': in; 'out': out;
- `nbr_route_map` (String) Apply route map to neighbor (Name of route map)


