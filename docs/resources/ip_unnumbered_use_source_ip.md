---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_unnumbered_use_source_ip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_unnumbered_use_source_ip: Source IP address
  PLACEHOLDER
---

# thunder_ip_unnumbered_use_source_ip (Resource)

`thunder_ip_unnumbered_use_source_ip`: Source IP address

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_unnumbered_use_source_ip" "thunder_ip_unnumbered_use_source_ip" {
  update_source_ip = "10.10.10.10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `update_source_ip` (String) IP address
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


