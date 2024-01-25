---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_sctp_permit_payload_protocol_protocol_name Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_sctp_permit_payload_protocol_protocol_name: Configure SCTP permitted payload protocols
  PLACEHOLDER
---

# thunder_cgnv6_sctp_permit_payload_protocol_protocol_name (Resource)

`thunder_cgnv6_sctp_permit_payload_protocol_protocol_name`: Configure SCTP permitted payload protocols

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_sctp_permit_payload_protocol_protocol_name" "thunder_cgnv6_sctp_permit_payload_protocol_protocol_name" {
  protocol = "iua"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `protocol` (String) 'iua': IUA; 'm2ua': M2UA; 'm3ua': M3UA; 'sua': SUA; 'm2pa': M2PA; 'h.323': H.323;

### Optional

- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

