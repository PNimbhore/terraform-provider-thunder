---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_common_quic Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_common_quic: Configure QUIC related settings
  PLACEHOLDER
---

# thunder_slb_common_quic (Resource)

`thunder_slb_common_quic`: Configure QUIC related settings

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_quic" "thunder_slb_common_quic" {
  cid_len          = 4
  cpu_offset       = 12
  enable_hash      = 0
  enable_signature = 0
  quic_lb_offset   = 8
  signature        = "122"
  signature_len    = 3
  signature_offset = 4
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cid_len` (Number) Length of CID
- `cpu_offset` (Number) Offset for Encoded CPU
- `enable_hash` (Number) Enable CID Hashing
- `enable_signature` (Number) Enable CID Signature Validation
- `quic_lb_offset` (Number) Offset for QUIC-LB
- `signature` (String) Set CID Signature
- `signature_len` (Number) Offset for CID Signature
- `signature_offset` (Number) Offset for CID Signature
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


