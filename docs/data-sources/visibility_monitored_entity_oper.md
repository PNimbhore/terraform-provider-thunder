---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_monitored_entity_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_monitored_entity_oper: Operational Status for the object monitored-entity
  PLACEHOLDER
---

# thunder_visibility_monitored_entity_oper (Data Source)

`thunder_visibility_monitored_entity_oper`: Operational Status for the object monitored-entity

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_monitored_entity_oper" "thunder_visibility_monitored_entity_oper" {

}
output "get_visibility_monitored_entity_oper" {
  value = ["${data.thunder_visibility_monitored_entity_oper.thunder_visibility_monitored_entity_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `detail` (Block List, Max: 1) (see [below for nested schema](#nestedblock--detail))
- `mon_topk` (Block List, Max: 1) (see [below for nested schema](#nestedblock--mon_topk))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `secondary` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary))
- `sessions` (Block List, Max: 1) (see [below for nested schema](#nestedblock--sessions))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--detail"></a>
### Nested Schema for `detail`

Optional:

- `debug` (Block List, Max: 1) (see [below for nested schema](#nestedblock--detail--debug))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--detail--oper))

<a id="nestedblock--detail--debug"></a>
### Nested Schema for `detail.debug`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--detail--debug--oper))

<a id="nestedblock--detail--debug--oper"></a>
### Nested Schema for `detail.debug.oper`

Optional:

- `all_keys` (Number)
- `mon_entity_list` (Block List) (see [below for nested schema](#nestedblock--detail--debug--oper--mon_entity_list))
- `primary_keys` (Number)

<a id="nestedblock--detail--debug--oper--mon_entity_list"></a>
### Nested Schema for `detail.debug.oper.mon_entity_list`

Optional:

- `entity_key` (String)
- `entity_metric_list` (Block List) (see [below for nested schema](#nestedblock--detail--debug--oper--mon_entity_list--entity_metric_list))
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `sec_entity_list` (Block List) (see [below for nested schema](#nestedblock--detail--debug--oper--mon_entity_list--sec_entity_list))
- `uuid` (String)

<a id="nestedblock--detail--debug--oper--mon_entity_list--entity_metric_list"></a>
### Nested Schema for `detail.debug.oper.mon_entity_list.entity_metric_list`

Optional:

- `anomaly` (String)
- `current` (String)
- `max` (String)
- `mean` (String)
- `metric_name` (String)
- `min` (String)
- `std_dev` (String)
- `threshold` (String)


<a id="nestedblock--detail--debug--oper--mon_entity_list--sec_entity_list"></a>
### Nested Schema for `detail.debug.oper.mon_entity_list.sec_entity_list`

Optional:

- `entity_key` (String)
- `entity_metric_list` (Block List) (see [below for nested schema](#nestedblock--detail--debug--oper--mon_entity_list--sec_entity_list--entity_metric_list))
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `uuid` (String)

<a id="nestedblock--detail--debug--oper--mon_entity_list--sec_entity_list--entity_metric_list"></a>
### Nested Schema for `detail.debug.oper.mon_entity_list.sec_entity_list.entity_metric_list`

Optional:

- `anomaly` (String)
- `current` (String)
- `max` (String)
- `mean` (String)
- `metric_name` (String)
- `min` (String)
- `std_dev` (String)
- `threshold` (String)






<a id="nestedblock--detail--oper"></a>
### Nested Schema for `detail.oper`

Optional:

- `all_keys` (Number)
- `mon_entity_list` (Block List) (see [below for nested schema](#nestedblock--detail--oper--mon_entity_list))
- `primary_keys` (Number)

<a id="nestedblock--detail--oper--mon_entity_list"></a>
### Nested Schema for `detail.oper.mon_entity_list`

Optional:

- `entity_key` (String)
- `entity_metric_list` (Block List) (see [below for nested schema](#nestedblock--detail--oper--mon_entity_list--entity_metric_list))
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `sec_entity_list` (Block List) (see [below for nested schema](#nestedblock--detail--oper--mon_entity_list--sec_entity_list))
- `uuid` (String)

<a id="nestedblock--detail--oper--mon_entity_list--entity_metric_list"></a>
### Nested Schema for `detail.oper.mon_entity_list.entity_metric_list`

Optional:

- `anomaly` (String)
- `current` (String)
- `metric_name` (String)
- `threshold` (String)


<a id="nestedblock--detail--oper--mon_entity_list--sec_entity_list"></a>
### Nested Schema for `detail.oper.mon_entity_list.sec_entity_list`

Optional:

- `entity_key` (String)
- `entity_metric_list` (Block List) (see [below for nested schema](#nestedblock--detail--oper--mon_entity_list--sec_entity_list--entity_metric_list))
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `uuid` (String)

<a id="nestedblock--detail--oper--mon_entity_list--sec_entity_list--entity_metric_list"></a>
### Nested Schema for `detail.oper.mon_entity_list.sec_entity_list.entity_metric_list`

Optional:

- `anomaly` (String)
- `current` (String)
- `metric_name` (String)
- `threshold` (String)






<a id="nestedblock--mon_topk"></a>
### Nested Schema for `mon_topk`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--mon_topk--oper))
- `sources` (Block List, Max: 1) (see [below for nested schema](#nestedblock--mon_topk--sources))

<a id="nestedblock--mon_topk--oper"></a>
### Nested Schema for `mon_topk.oper`

Optional:

- `metric_topk_list` (Block List) (see [below for nested schema](#nestedblock--mon_topk--oper--metric_topk_list))

<a id="nestedblock--mon_topk--oper--metric_topk_list"></a>
### Nested Schema for `mon_topk.oper.metric_topk_list`

Optional:

- `metric_name` (String)
- `topk_list` (Block List) (see [below for nested schema](#nestedblock--mon_topk--oper--metric_topk_list--topk_list))

<a id="nestedblock--mon_topk--oper--metric_topk_list--topk_list"></a>
### Nested Schema for `mon_topk.oper.metric_topk_list.topk_list`

Optional:

- `ip_addr` (String)
- `metric_value` (String)
- `port` (Number)
- `protocol` (String)




<a id="nestedblock--mon_topk--sources"></a>
### Nested Schema for `mon_topk.sources`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--mon_topk--sources--oper))

<a id="nestedblock--mon_topk--sources--oper"></a>
### Nested Schema for `mon_topk.sources.oper`

Optional:

- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `metric_topk_list` (Block List) (see [below for nested schema](#nestedblock--mon_topk--sources--oper--metric_topk_list))

<a id="nestedblock--mon_topk--sources--oper--metric_topk_list"></a>
### Nested Schema for `mon_topk.sources.oper.metric_topk_list`

Optional:

- `metric_name` (String)
- `topk_list` (Block List) (see [below for nested schema](#nestedblock--mon_topk--sources--oper--metric_topk_list--topk_list))

<a id="nestedblock--mon_topk--sources--oper--metric_topk_list--topk_list"></a>
### Nested Schema for `mon_topk.sources.oper.metric_topk_list.topk_list`

Optional:

- `ip_addr` (String)
- `metric_value` (String)






<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `all_keys` (Number)
- `mon_entity_list` (Block List) (see [below for nested schema](#nestedblock--oper--mon_entity_list))
- `primary_keys` (Number)

<a id="nestedblock--oper--mon_entity_list"></a>
### Nested Schema for `oper.mon_entity_list`

Optional:

- `entity_key` (String)
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `sec_entity_list` (Block List) (see [below for nested schema](#nestedblock--oper--mon_entity_list--sec_entity_list))
- `uuid` (String)

<a id="nestedblock--oper--mon_entity_list--sec_entity_list"></a>
### Nested Schema for `oper.mon_entity_list.sec_entity_list`

Optional:

- `entity_key` (String)
- `flat_oid` (Number)
- `ha_state` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `mode` (String)
- `uuid` (String)




<a id="nestedblock--secondary"></a>
### Nested Schema for `secondary`

Optional:

- `mon_topk` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary--mon_topk))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary--oper))

<a id="nestedblock--secondary--mon_topk"></a>
### Nested Schema for `secondary.mon_topk`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary--mon_topk--oper))
- `sources` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary--mon_topk--sources))

<a id="nestedblock--secondary--mon_topk--oper"></a>
### Nested Schema for `secondary.mon_topk.oper`

Optional:

- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `metric_topk_list` (Block List) (see [below for nested schema](#nestedblock--secondary--mon_topk--oper--metric_topk_list))

<a id="nestedblock--secondary--mon_topk--oper--metric_topk_list"></a>
### Nested Schema for `secondary.mon_topk.oper.metric_topk_list`

Optional:

- `metric_name` (String)
- `topk_list` (Block List) (see [below for nested schema](#nestedblock--secondary--mon_topk--oper--metric_topk_list--topk_list))

<a id="nestedblock--secondary--mon_topk--oper--metric_topk_list--topk_list"></a>
### Nested Schema for `secondary.mon_topk.oper.metric_topk_list.topk_list`

Optional:

- `ip_addr` (String)
- `metric_value` (String)
- `port` (Number)
- `protocol` (String)




<a id="nestedblock--secondary--mon_topk--sources"></a>
### Nested Schema for `secondary.mon_topk.sources`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secondary--mon_topk--sources--oper))

<a id="nestedblock--secondary--mon_topk--sources--oper"></a>
### Nested Schema for `secondary.mon_topk.sources.oper`

Optional:

- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `metric_topk_list` (Block List) (see [below for nested schema](#nestedblock--secondary--mon_topk--sources--oper--metric_topk_list))

<a id="nestedblock--secondary--mon_topk--sources--oper--metric_topk_list"></a>
### Nested Schema for `secondary.mon_topk.sources.oper.metric_topk_list`

Optional:

- `metric_name` (String)
- `topk_list` (Block List) (see [below for nested schema](#nestedblock--secondary--mon_topk--sources--oper--metric_topk_list--topk_list))

<a id="nestedblock--secondary--mon_topk--sources--oper--metric_topk_list--topk_list"></a>
### Nested Schema for `secondary.mon_topk.sources.oper.metric_topk_list.topk_list`

Optional:

- `ip_addr` (String)
- `metric_value` (String)






<a id="nestedblock--secondary--oper"></a>
### Nested Schema for `secondary.oper`



<a id="nestedblock--sessions"></a>
### Nested Schema for `sessions`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--sessions--oper))

<a id="nestedblock--sessions--oper"></a>
### Nested Schema for `sessions.oper`

Optional:

- `mon_entity_list` (Block List) (see [below for nested schema](#nestedblock--sessions--oper--mon_entity_list))

<a id="nestedblock--sessions--oper--mon_entity_list"></a>
### Nested Schema for `sessions.oper.mon_entity_list`

Optional:

- `entity_key` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `sec_entity_list` (Block List) (see [below for nested schema](#nestedblock--sessions--oper--mon_entity_list--sec_entity_list))
- `session_list` (Block List) (see [below for nested schema](#nestedblock--sessions--oper--mon_entity_list--session_list))

<a id="nestedblock--sessions--oper--mon_entity_list--sec_entity_list"></a>
### Nested Schema for `sessions.oper.mon_entity_list.sec_entity_list`

Optional:

- `entity_key` (String)
- `ipv4_addr` (String)
- `ipv6_addr` (String)
- `l4_port` (Number)
- `l4_proto` (String)
- `session_list` (Block List) (see [below for nested schema](#nestedblock--sessions--oper--mon_entity_list--sec_entity_list--session_list))

<a id="nestedblock--sessions--oper--mon_entity_list--sec_entity_list--session_list"></a>
### Nested Schema for `sessions.oper.mon_entity_list.sec_entity_list.session_list`

Optional:

- `fwd_dst_ip` (String)
- `fwd_dst_port` (Number)
- `fwd_src_ip` (String)
- `fwd_src_port` (Number)
- `proto` (String)
- `rev_dst_ip` (String)
- `rev_dst_port` (Number)
- `rev_src_ip` (String)
- `rev_src_port` (Number)



<a id="nestedblock--sessions--oper--mon_entity_list--session_list"></a>
### Nested Schema for `sessions.oper.mon_entity_list.session_list`

Optional:

- `fwd_dst_ip` (String)
- `fwd_dst_port` (Number)
- `fwd_src_ip` (String)
- `fwd_src_port` (Number)
- `proto` (String)
- `rev_dst_ip` (String)
- `rev_dst_port` (Number)
- `rev_src_ip` (String)
- `rev_src_port` (Number)


