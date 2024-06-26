---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_port_zone_service_src_based_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_port_zone_service_src_based_policy: Configure src-based policy
  PLACEHOLDER
---

# thunder_ddos_dst_zone_port_zone_service_src_based_policy (Resource)

`thunder_ddos_dst_zone_port_zone_service_src_based_policy`: Configure src-based policy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_src_based_policy" "thunder_ddos_dst_zone_port_zone_service_src_based_policy" {
  zone_name             = "test"
  port_num              = "20"
  protocol              = "tcp"
  src_based_policy_name = "test"
  user_tag              = "testing"
  policy_class_list_list {
    class_list_name         = "test"
    glid                    = "3"
    action                  = "deny"
    max_dynamic_entry_count = 33
    user_tag                = "test"
    class_list_overflow_policy_list {
      dummy_name = "configuration"
      glid       = "3"
      action     = "deny"
      log_enable = 1
      user_tag   = "test"
    }

  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port_num` (String) PortNum
- `protocol` (String) Protocol
- `src_based_policy_name` (String) Specify name of the policy
- `zone_name` (String) ZoneName

### Optional

- `policy_class_list_list` (Block List) (see [below for nested schema](#nestedblock--policy_class_list_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--policy_class_list_list"></a>
### Nested Schema for `policy_class_list_list`

Required:

- `class_list_name` (String) Class-list name

Optional:

- `action` (String) 'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;
- `class_list_overflow_policy_list` (Block List) (see [below for nested schema](#nestedblock--policy_class_list_list--class_list_overflow_policy_list))
- `glid` (String) Global limit ID
- `glid_action` (String) 'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;
- `log_enable` (Number) Enable logging
- `log_periodic` (Number) Enable log periodic
- `max_dynamic_entry_count` (Number) Maximum count for dynamic source zone service entry allowed for this class-list
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--policy_class_list_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `zone_template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--policy_class_list_list--zone_template))

<a id="nestedblock--policy_class_list_list--class_list_overflow_policy_list"></a>
### Nested Schema for `policy_class_list_list.class_list_overflow_policy_list`

Required:

- `dummy_name` (String) 'configuration': Configure overflow policy for class-list;

Optional:

- `action` (String) 'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;
- `glid` (String) Global limit ID
- `log_enable` (Number) Enable logging
- `log_periodic` (Number) Enable log periodic
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `zone_template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--policy_class_list_list--class_list_overflow_policy_list--zone_template))

<a id="nestedblock--policy_class_list_list--class_list_overflow_policy_list--zone_template"></a>
### Nested Schema for `policy_class_list_list.class_list_overflow_policy_list.zone_template`

Optional:

- `dns` (String) DDOS dns template
- `encap` (String) DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `http` (String) DDOS http template
- `logging` (String) DDOS logging template
- `quic` (String) DDOS quic template
- `sip` (String) DDOS sip template
- `ssl_l4` (String) DDOS ssl-l4 template
- `tcp` (String) DDOS tcp template
- `udp` (String) DDOS udp template



<a id="nestedblock--policy_class_list_list--sampling_enable"></a>
### Nested Schema for `policy_class_list_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;


<a id="nestedblock--policy_class_list_list--zone_template"></a>
### Nested Schema for `policy_class_list_list.zone_template`

Optional:

- `dns` (String) DDOS dns template
- `encap` (String) DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `http` (String) DDOS http template
- `logging` (String) DDOS logging template
- `quic` (String) DDOS quic template
- `sip` (String) DDOS sip template
- `ssl_l4` (String) DDOS ssl-l4 template
- `tcp` (String) DDOS tcp template
- `udp` (String) DDOS udp template


