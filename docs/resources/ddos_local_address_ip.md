---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_local_address_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_local_address_ip: Configure DDoS ipv4 address
  PLACEHOLDER
---

# thunder_ddos_local_address_ip (Resource)

`thunder_ddos_local_address_ip`: Configure DDoS ipv4 address

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_local_address_ip" "thunder_ddos_local_address_ip" {
  ip_addr = "10.10.10.10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ip_addr` (String) DDoS IPv4 Address for syn cookie usage

### Optional

- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


