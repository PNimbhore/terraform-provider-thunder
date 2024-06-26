---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_distributed_forwarding_fw Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_distributed_forwarding_fw: Enable Scaleout distributed-forwarding for fw
  PLACEHOLDER
---

# thunder_scaleout_distributed_forwarding_fw (Resource)

`thunder_scaleout_distributed_forwarding_fw`: Enable Scaleout distributed-forwarding for fw

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_distributed_forwarding_fw" "thunder_scaleout_distributed_forwarding_fw" {
  fw_value                  = "enable"
  session_offload_direction = "both"
  threshold {
    threshold_value = 51
    protocol_value  = "TCP"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `fw_value` (String) 'enable': Enable FW; 'disable': Disable FW;
- `session_offload_direction` (String) 'uplink': Enable session offload only in uplink direction; 'downlink': Enable session offload in downlink direction; 'both': Enable session offload in both direction;
- `threshold` (Block List) (see [below for nested schema](#nestedblock--threshold))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--threshold"></a>
### Nested Schema for `threshold`

Optional:

- `protocol_value` (String) 'UDP': configure threshold for udp session offload; 'TCP': configure threshold for tcp session offload;
- `threshold_value` (Number) configure packet threshold value to offload sessions(default 5)


