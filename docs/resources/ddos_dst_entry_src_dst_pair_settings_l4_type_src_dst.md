---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst: DDOS L4 type
  PLACEHOLDER
---

# thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst (Resource)

`thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst`: DDOS L4 type

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst" "thunder_ddos_dst_entry_src_dst_pair_settings_l4_type_src_dst" {
  dst_entry_name          = "test"
  max_dynamic_entry_count = 536
  protocol                = "tcp"
  user_tag                = "107"
  all_types               = "all-types"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `all_types` (String) AllTypes
- `dst_entry_name` (String) DstEntryName
- `protocol` (String) 'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;

### Optional

- `apply_policy_on_overflow` (Number) Enable this flag to apply overflow policy when dynamic entry count overflows
- `max_dynamic_entry_count` (Number) Maximum count for dynamic src-dst entry
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


