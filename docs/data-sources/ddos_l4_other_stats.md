---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_l4_other_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_l4_other_stats: Statistics for the object l4-other
  PLACEHOLDER
---

# thunder_ddos_l4_other_stats (Data Source)

`thunder_ddos_l4_other_stats`: Statistics for the object l4-other

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_l4_other_stats" "thunder_ddos_l4_other_stats" {
}
output "get_ddos_l4_other_stats" {
  value = ["${data.thunder_ddos_l4_other_stats.thunder_ddos_l4_other_stats}"]
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

- `dst_other_filter_action_blacklist` (Number) Dst Filter Action Blacklist
- `dst_other_filter_action_default_pass` (Number) Dst Filter Action Default Pass
- `dst_other_filter_action_drop` (Number) Dst Filter Action Drop
- `dst_other_filter_action_whitelist` (Number) Dst Filter Action WL
- `dst_other_filter_match` (Number) Dst Filter Match
- `dst_other_filter_not_match` (Number) Dst Filter No Match
- `other_any_exceed` (Number) OTHER Exceeded
- `other_drop_bl` (Number) OTHER Blacklisted Packets Dropped
- `other_drop_black_user_cfg_src` (Number) OTHER Src Blacklist User Packets Dropped
- `other_drop_black_user_cfg_src_dst` (Number) OTHER SrcDst Blacklist User Packets Dropped
- `other_dst_drop` (Number) OTHER Dst Packets Dropped
- `other_frag_drop` (Number) OTHER Frag Dropped
- `other_frag_rcvd` (Number) OTHER Frag Received
- `other_receive` (Number) OTHER Total Packets Received
- `other_src_drop` (Number) OTHER Src Packets Dropped
- `other_src_dst_drop` (Number) OTHER SrcDst Packets Dropped
- `other_total_bytes_drop` (Number) OTHER Total Bytes Dropped
- `other_total_bytes_rcv` (Number) OTHER Total Bytes Received
- `other_total_drop` (Number) OTHER Total Packets Dropped
- `src_dst_other_filter_action_blacklist` (Number) SrcDst Filter Action Blacklist
- `src_dst_other_filter_action_default_pass` (Number) SrcDst Filter Action Default Pass
- `src_dst_other_filter_action_drop` (Number) SrcDst Filter Action Drop
- `src_dst_other_filter_action_whitelist` (Number) SrcDst Filter Action WL
- `src_dst_other_filter_match` (Number) SrcDst Filter Match
- `src_dst_other_filter_not_match` (Number) SrcDst Filter No Match
- `src_other_filter_action_blacklist` (Number) Src Filter Action Blacklist
- `src_other_filter_action_default_pass` (Number) Src Filter Action Default Pass
- `src_other_filter_action_drop` (Number) Src Filter Action Drop
- `src_other_filter_action_whitelist` (Number) Src Filter Action WL
- `src_other_filter_match` (Number) Src Filter Match
- `src_other_filter_not_match` (Number) Src Filter No Match

