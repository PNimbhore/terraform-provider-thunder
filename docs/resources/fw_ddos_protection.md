---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_ddos_protection Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_ddos_protection: Configure FW DDoS Protection
  PLACEHOLDER
---

# thunder_fw_ddos_protection (Resource)

`thunder_fw_ddos_protection`: Configure FW DDoS Protection

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_ddos_protection" "thunder_fw_ddos_protection" {
  action {
    action_type = "drop"
    expiration  = 5
  }
  dynamic_blacklist {
    dynamic_blacklist_action = "enable"
    dir                      = "both"
  }
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (Block List, Max: 1) (see [below for nested schema](#nestedblock--action))
- `dynamic_blacklist` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dynamic_blacklist))
- `logging` (Block List, Max: 1) (see [below for nested schema](#nestedblock--logging))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--action"></a>
### Nested Schema for `action`

Optional:

- `action_type` (String) 'drop': Log, and drop all packets (default); 'redistribute-route': Log, Drop, and Notify upstream router to reroute the packets;
- `expiration` (Number) To specify time in minutes to revert the action (Expiration time, in minutes (default is 5 mins))
- `expiration_route` (Number) To specify time in minutes to revert the action (Expiration time, in minutes (default is 60 mins))
- `remove_wait_timer` (Number) Max time to wait before removing IP from blackhole (Max value in seconds (default 300))
- `route_map` (String) Route map name
- `timer_multiply_max` (Number) To specify max value of timer multiplier for attacks lasted long time (Max value of timer multiplier (default is 6))


<a id="nestedblock--dynamic_blacklist"></a>
### Nested Schema for `dynamic_blacklist`

Optional:

- `cpu_threshold` (Number) Core-level CPU usage threshold for dynamic blacklist creation (Core-level CPU usage threshold for dynamic blacklist creation (default is 60))
- `dir` (String) 'inbound': enable in inbound direction; 'outbound': enable in outbound direction; 'both': enable in both directions;
- `dynamic_blacklist_action` (String) 'enable': Enable protection against volumetric attacks using dynamic blacklist; 'disable': Disable protection against volumetric attacks using dynamic blacklist;
- `timeout` (Number) Timeout value (in seconds) for dynamic blacklist (Timeout value (in seconds) for dynamic blacklist(default is 5 seconds))


<a id="nestedblock--logging"></a>
### Nested Schema for `logging`

Optional:

- `enable_action` (String) 'local': Enable local logs only; 'remote': Enable logging to remote server & IPFIX; 'both': Enable both local & remote logs;
- `logging_action` (String) 'enable': enable FW DDoS protection logging; 'disable': Disable both local & remote FW DDoS protection logging;


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'ddos_entries_too_many': Too many DDOS entries; 'ddos_entry_added': DDOS entry added; 'ddos_entry_removed': DDOS entry removed; 'ddos_entry_added_to_bgp': DDoS Entry added to BGP; 'ddos_entry_removed_from_bgp': DDoS Entry Removed from BGP; 'ddos_entry_add_to_bgp_failure': DDoS Entry BGP add failures; 'ddos_entry_remove_from_bgp_failure': DDOS entry BGP remove failures; 'ddos_packet_dropped': DDOS Packet Drop;

