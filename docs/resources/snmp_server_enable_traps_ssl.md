---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_enable_traps_ssl Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_enable_traps_ssl: Enable SSL group traps
  PLACEHOLDER
---

# thunder_snmp_server_enable_traps_ssl (Resource)

`thunder_snmp_server_enable_traps_ssl`: Enable SSL group traps

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_ssl" "thunder_snmp_server_enable_traps_ssl" {
  server_certificate_error = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `server_certificate_error` (Number) Enable SSL server certificate error trap
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


