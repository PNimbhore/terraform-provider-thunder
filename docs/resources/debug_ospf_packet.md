---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_ospf_packet Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_ospf_packet: OSPFv3 packets
  PLACEHOLDER
---

# thunder_debug_ospf_packet (Resource)

`thunder_debug_ospf_packet`: OSPFv3 packets

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_ospf_packet" "thunder_debug_ospf_packet" {
  dd         = 0
  detail     = 0
  hello      = 0
  ls_ack     = 0
  ls_request = 0
  ls_update  = 0
  recv       = 0
  send       = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `dd` (Number) OSPFv3 Database Description
- `detail` (Number) Detail information
- `hello` (Number) OSPFv3 Hello
- `ls_ack` (Number) OSPFv3 Link State Acknowledgment
- `ls_request` (Number) OSPFv3 Link State Request
- `ls_update` (Number) OSPFv3 Link State Update
- `recv` (Number) Packet received
- `send` (Number) Packet sent
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


