---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_map_translation_fragmentation_inbound Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_map_translation_fragmentation_inbound: MAP-T fragmentation rules for inbound oversize packets
  PLACEHOLDER
---

# thunder_cgnv6_map_translation_fragmentation_inbound (Resource)

`thunder_cgnv6_map_translation_fragmentation_inbound`: MAP-T fragmentation rules for inbound oversize packets

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_map_translation_fragmentation_inbound" "thunder_cgnv6_map_translation_fragmentation_inbound" {

  df_set      = "drop"
  frag_action = "ipv6"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `df_set` (String) 'drop': Drop Silently; 'ipv6': Use IPv6 fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default);
- `frag_action` (String) 'drop': Drop Silently; 'ipv6': Use IPv6 fragmentation for oversize packets (default);
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


