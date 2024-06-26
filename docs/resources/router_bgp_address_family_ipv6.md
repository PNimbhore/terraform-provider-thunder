---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_address_family_ipv6 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_router_bgp_address_family_ipv6: ipv6 Address family
  PLACEHOLDER
---

# thunder_router_bgp_address_family_ipv6 (Resource)

`thunder_router_bgp_address_family_ipv6`: ipv6 Address family

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_address_family_ipv6" "thunderRouterBgpAddressFamilyIpv6Test" {
  as_number = "300"
  bgp {
    dampening                = 1
    dampening_half           = 1
    dampening_start_reuse    = 1
    dampening_start_supress  = 1
    dampening_max_supress    = 1
    dampening_unreachability = 1
  }
  distance {
    distance_ext   = 1
    distance_int   = 1
    distance_local = 1
  }
  maximum_paths_value = 2
  auto_summary        = 1
  synchronization     = 1
  originate           = 1
  network {
    synchronization {
      network_synchronization = 1
    }
    ipv6_network_list {
      network_ipv6 = "1111::1/64"
      route_map    = "ipv6Routemap"
      backdoor     = 1
      description  = "AddIpv6"
      comm_value   = "internet"
    }
  }
  aggregate_address_list {
    aggregate_address = "3333::1/64"
    as_set            = 1
    summary_only      = 1
  }

  redistribute {
    connected_cfg {
      connected = 1
      route_map = "ipv6"
    }
    floating_ip_cfg {
      floating_ip = 1
      route_map   = "ipv6"
    }
    nat64_cfg {
      nat64     = 1
      route_map = "ipv6"
    }
    nat_map_cfg {
      nat_map   = 1
      route_map = "ipv6"
    }
    lw4o6_cfg {
      lw4o6     = 1
      route_map = "ipv6"
    }
    static_nat_cfg {
      static_nat = 1
      route_map  = "ipv6"
    }
    ip_nat_cfg {
      ip_nat    = 1
      route_map = "ipv6"
    }
    ip_nat_list_cfg {
      ip_nat_list = 1
      route_map   = "ipv6"
    }
    isis_cfg {
      isis      = 1
      route_map = "ipv6"
    }
    ospf_cfg {
      ospf      = 1
      route_map = "ipv6"
    }
    rip_cfg {
      rip       = 1
      route_map = "ipv6"
    }
    static_cfg {
      static    = 1
      route_map = "ipv6"
    }
    vip {
      only_flagged_cfg {
        only_flagged = 1
        route_map    = "ipv6"
      }
      only_not_flagged_cfg {
        only_not_flagged = 1
        route_map        = "ipv6"
      }
    }

  }
  neighbor {
    ipv4_neighbor_list {
      neighbor_ipv4    = "1.1.1.1"
      peer_group_name  = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10
      neighbor_filter_lists {
        filter_list           = "FilterBgp"
        filter_list_direction = "in"
      }
      maximum_prefix       = 100
      maximum_prefix_thres = 80
      neighbor_prefix_lists {
        nbr_prefix_list           = "prefixIpv6"
        nbr_prefix_list_direction = "in"
      }

      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1
      weight  = 300
    }
    ipv6_neighbor_list {
      neighbor_ipv6    = "2222::1"
      peer_group_name  = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10
      neighbor_filter_lists {
        filter_list           = "FilterBgp"
        filter_list_direction = "in"
      }
      maximum_prefix       = 100
      maximum_prefix_thres = 80
      neighbor_prefix_lists {
        nbr_prefix_list           = "prefixIpv6"
        nbr_prefix_list_direction = "in"
      }

      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1
      weight  = 300
    }
    ethernet_neighbor_ipv6_list {
      ethernet        = 4
      peer_group_name = "A10"
    }
    ve_neighbor_ipv6_list {
      ve              = 300
      peer_group_name = "A10"
    }
    peer_group_neighbor_list {
      peer_group       = "A10"
      activate         = 1
      allowas_in       = 1
      allowas_in_count = 10

      maximum_prefix       = 100
      maximum_prefix_thres = 80
      next_hop_self        = 1
      remove_private_as    = 1
      neighbor_route_map_lists {
        nbr_route_map      = "Ipv6RouteMap"
        nbr_rmap_direction = "in"
      }
      inbound = 1

      weight = 300
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) AsNumber

### Optional

- `aggregate_address_list` (Block List) (see [below for nested schema](#nestedblock--aggregate_address_list))
- `auto_summary` (Number) Enable automatic network number summarization
- `bgp` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bgp))
- `distance` (Block List, Max: 1) (see [below for nested schema](#nestedblock--distance))
- `maximum_paths_value` (Number) Supported BGP multipath numbers
- `neighbor` (Block List, Max: 1) (see [below for nested schema](#nestedblock--neighbor))
- `network` (Block List, Max: 1) (see [below for nested schema](#nestedblock--network))
- `originate` (Number) Distribute an IPv6 default route
- `redistribute` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute))
- `synchronization` (Number) Perform IGP synchronization
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--aggregate_address_list"></a>
### Nested Schema for `aggregate_address_list`

Optional:

- `aggregate_address` (String) Configure BGP aggregate entries (Aggregate IPv6 prefix)
- `as_set` (Number) Generate AS set path information
- `summary_only` (Number) Filter more specific routes from updates


<a id="nestedblock--bgp"></a>
### Nested Schema for `bgp`

Optional:

- `dampening` (Number) Enable route-flap dampening
- `dampening_half` (Number) Reachability Half-life time for the penalty(minutes)
- `dampening_max_supress` (Number) Maximum duration to suppress a stable route(minutes)
- `dampening_start_reuse` (Number) Value to start reusing a route
- `dampening_start_supress` (Number) Value to start suppressing a route
- `dampening_unreachability` (Number) Un-reachability Half-life time for the penalty(minutes)
- `route_map` (String) Route-map to specify criteria for dampening (Route-map name)


<a id="nestedblock--distance"></a>
### Nested Schema for `distance`

Optional:

- `distance_ext` (Number) Distance for routes external to the AS
- `distance_int` (Number) Distance for routes internal to the AS
- `distance_local` (Number) Distance for local routes


<a id="nestedblock--neighbor"></a>
### Nested Schema for `neighbor`

Optional:

- `ethernet_neighbor_ipv6_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--ethernet_neighbor_ipv6_list))
- `ipv4_neighbor_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv4_neighbor_list))
- `ipv6_neighbor_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv6_neighbor_list))
- `peer_group_neighbor_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--peer_group_neighbor_list))
- `trunk_neighbor_ipv6_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--trunk_neighbor_ipv6_list))
- `ve_neighbor_ipv6_list` (Block List) (see [below for nested schema](#nestedblock--neighbor--ve_neighbor_ipv6_list))

<a id="nestedblock--neighbor--ethernet_neighbor_ipv6_list"></a>
### Nested Schema for `neighbor.ethernet_neighbor_ipv6_list`

Required:

- `ethernet` (Number) Ethernet interface number

Optional:

- `peer_group_name` (String)
- `uuid` (String) uuid of the object


<a id="nestedblock--neighbor--ipv4_neighbor_list"></a>
### Nested Schema for `neighbor.ipv4_neighbor_list`

Required:

- `neighbor_ipv4` (String) Neighbor address

Optional:

- `activate` (Number) Enable the Address Family for this Neighbor
- `allowas_in` (Number) Accept as-path with my AS present in it
- `allowas_in_count` (Number) Number of occurrences of AS number
- `default_originate` (Number) Originate default route to this neighbor
- `distribute_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv4_neighbor_list--distribute_lists))
- `graceful_restart` (Number) enable graceful-restart helper for this neighbor
- `inbound` (Number) Allow inbound soft reconfiguration for this neighbor
- `maximum_prefix` (Number) Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
- `maximum_prefix_thres` (Number) threshold-value, 1 to 100 percent
- `neighbor_filter_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv4_neighbor_list--neighbor_filter_lists))
- `neighbor_prefix_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv4_neighbor_list--neighbor_prefix_lists))
- `neighbor_route_map_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv4_neighbor_list--neighbor_route_map_lists))
- `next_hop_self` (Number) Disable the next hop calculation for this neighbor
- `peer_group_name` (String) Configure peer-group (peer-group name)
- `prefix_list_direction` (String) 'both': both; 'receive': receive; 'send': send;
- `remove_private_as` (Number) Remove private AS number from outbound updates
- `restart_min` (Number) restart value, 1 to 1440 minutes
- `route_map` (String) Route-map to specify criteria to originate default (route-map name)
- `send_community_val` (String) 'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes (Name of route map)
- `uuid` (String) uuid of the object
- `weight` (Number) Set default weight for routes from this neighbor

<a id="nestedblock--neighbor--ipv4_neighbor_list--distribute_lists"></a>
### Nested Schema for `neighbor.ipv4_neighbor_list.distribute_lists`

Optional:

- `distribute_list` (String) Filter updates to/from this neighbor (IP standard/extended/named access list)
- `distribute_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv4_neighbor_list--neighbor_filter_lists"></a>
### Nested Schema for `neighbor.ipv4_neighbor_list.neighbor_filter_lists`

Optional:

- `filter_list` (String) Establish BGP filters (AS path access-list name)
- `filter_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv4_neighbor_list--neighbor_prefix_lists"></a>
### Nested Schema for `neighbor.ipv4_neighbor_list.neighbor_prefix_lists`

Optional:

- `nbr_prefix_list` (String) Filter updates to/from this neighbor (Name of a prefix list)
- `nbr_prefix_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv4_neighbor_list--neighbor_route_map_lists"></a>
### Nested Schema for `neighbor.ipv4_neighbor_list.neighbor_route_map_lists`

Optional:

- `nbr_rmap_direction` (String) 'in': in; 'out': out;
- `nbr_route_map` (String) Apply route map to neighbor (Name of route map)



<a id="nestedblock--neighbor--ipv6_neighbor_list"></a>
### Nested Schema for `neighbor.ipv6_neighbor_list`

Required:

- `neighbor_ipv6` (String) Neighbor IPv6 address

Optional:

- `activate` (Number) Enable the Address Family for this Neighbor
- `allowas_in` (Number) Accept as-path with my AS present in it
- `allowas_in_count` (Number) Number of occurrences of AS number
- `default_originate` (Number) Originate default route to this neighbor
- `distribute_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv6_neighbor_list--distribute_lists))
- `graceful_restart` (Number) enable graceful-restart helper for this neighbor
- `inbound` (Number) Allow inbound soft reconfiguration for this neighbor
- `maximum_prefix` (Number) Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
- `maximum_prefix_thres` (Number) threshold-value, 1 to 100 percent
- `neighbor_filter_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv6_neighbor_list--neighbor_filter_lists))
- `neighbor_prefix_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv6_neighbor_list--neighbor_prefix_lists))
- `neighbor_route_map_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--ipv6_neighbor_list--neighbor_route_map_lists))
- `next_hop_self` (Number) Disable the next hop calculation for this neighbor
- `peer_group_name` (String) Configure peer-group (peer-group name)
- `prefix_list_direction` (String) 'both': both; 'receive': receive; 'send': send;
- `remove_private_as` (Number) Remove private AS number from outbound updates
- `restart_min` (Number) restart value, 1 to 1440 minutes
- `route_map` (String) Route-map to specify criteria to originate default (route-map name)
- `send_community_val` (String) 'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes (Name of route map)
- `uuid` (String) uuid of the object
- `weight` (Number) Set default weight for routes from this neighbor

<a id="nestedblock--neighbor--ipv6_neighbor_list--distribute_lists"></a>
### Nested Schema for `neighbor.ipv6_neighbor_list.distribute_lists`

Optional:

- `distribute_list` (String) Filter updates to/from this neighbor (IP standard/extended/named access list)
- `distribute_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv6_neighbor_list--neighbor_filter_lists"></a>
### Nested Schema for `neighbor.ipv6_neighbor_list.neighbor_filter_lists`

Optional:

- `filter_list` (String) Establish BGP filters (AS path access-list name)
- `filter_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv6_neighbor_list--neighbor_prefix_lists"></a>
### Nested Schema for `neighbor.ipv6_neighbor_list.neighbor_prefix_lists`

Optional:

- `nbr_prefix_list` (String) Filter updates to/from this neighbor (Name of a prefix list)
- `nbr_prefix_list_direction` (String) 'in': in; 'out': out;


<a id="nestedblock--neighbor--ipv6_neighbor_list--neighbor_route_map_lists"></a>
### Nested Schema for `neighbor.ipv6_neighbor_list.neighbor_route_map_lists`

Optional:

- `nbr_rmap_direction` (String) 'in': in; 'out': out;
- `nbr_route_map` (String) Apply route map to neighbor (Name of route map)



<a id="nestedblock--neighbor--peer_group_neighbor_list"></a>
### Nested Schema for `neighbor.peer_group_neighbor_list`

Required:

- `peer_group` (String) Neighbor tag

Optional:

- `activate` (Number) Enable the Address Family for this Neighbor
- `allowas_in` (Number) Accept as-path with my AS present in it
- `allowas_in_count` (Number) Number of occurrences of AS number
- `inbound` (Number) Allow inbound soft reconfiguration for this neighbor
- `maximum_prefix` (Number) Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
- `maximum_prefix_thres` (Number) threshold-value, 1 to 100 percent
- `neighbor_route_map_lists` (Block List) (see [below for nested schema](#nestedblock--neighbor--peer_group_neighbor_list--neighbor_route_map_lists))
- `next_hop_self` (Number) Disable the next hop calculation for this neighbor
- `remove_private_as` (Number) Remove private AS number from outbound updates
- `uuid` (String) uuid of the object
- `weight` (Number) Set default weight for routes from this neighbor

<a id="nestedblock--neighbor--peer_group_neighbor_list--neighbor_route_map_lists"></a>
### Nested Schema for `neighbor.peer_group_neighbor_list.neighbor_route_map_lists`

Optional:

- `nbr_rmap_direction` (String) 'in': in; 'out': out;
- `nbr_route_map` (String) Apply route map to neighbor (Name of route map)



<a id="nestedblock--neighbor--trunk_neighbor_ipv6_list"></a>
### Nested Schema for `neighbor.trunk_neighbor_ipv6_list`

Required:

- `trunk` (Number) Trunk interface number

Optional:

- `peer_group_name` (String)
- `uuid` (String) uuid of the object


<a id="nestedblock--neighbor--ve_neighbor_ipv6_list"></a>
### Nested Schema for `neighbor.ve_neighbor_ipv6_list`

Required:

- `ve` (Number) Virtual ethernet interface number

Optional:

- `peer_group_name` (String)
- `uuid` (String) uuid of the object



<a id="nestedblock--network"></a>
### Nested Schema for `network`

Optional:

- `ipv6_network_list` (Block List) (see [below for nested schema](#nestedblock--network--ipv6_network_list))
- `monitor` (Block List, Max: 1) (see [below for nested schema](#nestedblock--network--monitor))
- `synchronization` (Block List, Max: 1) (see [below for nested schema](#nestedblock--network--synchronization))

<a id="nestedblock--network--ipv6_network_list"></a>
### Nested Schema for `network.ipv6_network_list`

Required:

- `network_ipv6` (String) Specify a network to announce via BGP

Optional:

- `backdoor` (Number) Specify a BGP backdoor route
- `comm_value` (String) community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export
- `description` (String) Network specific description (Up to 80 characters describing this network)
- `lcomm_value` (String) Large community value in the format XX:YY:ZZ
- `route_map` (String) Route-map to modify the attributes (Name of the route map)
- `uuid` (String) uuid of the object


<a id="nestedblock--network--monitor"></a>
### Nested Schema for `network.monitor`

Optional:

- `default` (Block List, Max: 1) (see [below for nested schema](#nestedblock--network--monitor--default))

<a id="nestedblock--network--monitor--default"></a>
### Nested Schema for `network.monitor.default`

Optional:

- `network_monitor_default` (Number) default route monitoring
- `uuid` (String) uuid of the object



<a id="nestedblock--network--synchronization"></a>
### Nested Schema for `network.synchronization`

Optional:

- `network_synchronization` (Number) Perform IGP synchronization
- `uuid` (String) uuid of the object



<a id="nestedblock--redistribute"></a>
### Nested Schema for `redistribute`

Optional:

- `connected_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--connected_cfg))
- `floating_ip_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--floating_ip_cfg))
- `ip_nat_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--ip_nat_cfg))
- `ip_nat_list_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--ip_nat_list_cfg))
- `isis_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--isis_cfg))
- `lw4o6_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--lw4o6_cfg))
- `nat64_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--nat64_cfg))
- `nat_map_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--nat_map_cfg))
- `ospf_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--ospf_cfg))
- `rip_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--rip_cfg))
- `static_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--static_cfg))
- `static_nat_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--static_nat_cfg))
- `uuid` (String) uuid of the object
- `vip` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--vip))

<a id="nestedblock--redistribute--connected_cfg"></a>
### Nested Schema for `redistribute.connected_cfg`

Optional:

- `connected` (Number) Connected
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--floating_ip_cfg"></a>
### Nested Schema for `redistribute.floating_ip_cfg`

Optional:

- `floating_ip` (Number) Floating IP
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--ip_nat_cfg"></a>
### Nested Schema for `redistribute.ip_nat_cfg`

Optional:

- `ip_nat` (Number) IP NAT
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--ip_nat_list_cfg"></a>
### Nested Schema for `redistribute.ip_nat_list_cfg`

Optional:

- `ip_nat_list` (Number) IP NAT list
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--isis_cfg"></a>
### Nested Schema for `redistribute.isis_cfg`

Optional:

- `isis` (Number) ISO IS-IS
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--lw4o6_cfg"></a>
### Nested Schema for `redistribute.lw4o6_cfg`

Optional:

- `lw4o6` (Number) LW4O6 Prefix
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--nat64_cfg"></a>
### Nested Schema for `redistribute.nat64_cfg`

Optional:

- `nat64` (Number) NAT64 Prefix
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--nat_map_cfg"></a>
### Nested Schema for `redistribute.nat_map_cfg`

Optional:

- `nat_map` (Number) NAT MAP Prefix
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--ospf_cfg"></a>
### Nested Schema for `redistribute.ospf_cfg`

Optional:

- `ospf` (Number) Open Shortest Path First (OSPF)
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--rip_cfg"></a>
### Nested Schema for `redistribute.rip_cfg`

Optional:

- `rip` (Number) Routing Information Protocol (RIP)
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--static_cfg"></a>
### Nested Schema for `redistribute.static_cfg`

Optional:

- `route_map` (String) Route map reference (Pointer to route-map entries)
- `static` (Number) Static routes


<a id="nestedblock--redistribute--static_nat_cfg"></a>
### Nested Schema for `redistribute.static_nat_cfg`

Optional:

- `route_map` (String) Route map reference (Pointer to route-map entries)
- `static_nat` (Number) Static NAT Prefix


<a id="nestedblock--redistribute--vip"></a>
### Nested Schema for `redistribute.vip`

Optional:

- `only_flagged_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--vip--only_flagged_cfg))
- `only_not_flagged_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--redistribute--vip--only_not_flagged_cfg))

<a id="nestedblock--redistribute--vip--only_flagged_cfg"></a>
### Nested Schema for `redistribute.vip.only_flagged_cfg`

Optional:

- `only_flagged` (Number) Selected Virtual IP (VIP)
- `route_map` (String) Route map reference (Pointer to route-map entries)


<a id="nestedblock--redistribute--vip--only_not_flagged_cfg"></a>
### Nested Schema for `redistribute.vip.only_not_flagged_cfg`

Optional:

- `only_not_flagged` (Number) Only not flagged
- `route_map` (String) Route map reference (Pointer to route-map entries)


