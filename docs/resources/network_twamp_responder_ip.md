---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_network_twamp_responder_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_network_twamp_responder_ip: Configure TWAMP responder
  PLACEHOLDER
---

# thunder_network_twamp_responder_ip (Resource)

`thunder_network_twamp_responder_ip`: Configure TWAMP responder

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_twamp_responder_ip" "thunder_network_twamp_responder_ip" {

  acl_id = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `acl_id` (Number) ACL id
- `acl_name` (String) Apply a named access list (Access List name)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

