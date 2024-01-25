---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_file_ca_cert Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_file_ca_cert (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_ca_cert" "thunder_file_ca_cert" {
  name             = "CACERT"
  protocol         = "http"
  host             = "10.64.3.190"
  path             = "/test123.pem"
  use_mgmt_port    = 1
  certificate_type = "pem"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) Remote site (IP or domain name)
- `name` (String) Local file name
- `path` (String) Remote path
- `protocol` (String) Transfer protocol

### Optional

- `certificate_type` (String) Certificate format
- `overwrite` (Number) Overwrite existing file
- `password` (String) Password for the remote site
- `pfx_password` (String) Password for pfx format
- `secured` (Number) Mark as non-exportable (for pfx format)
- `use_mgmt_port` (Number) Use management port as source port
- `username` (String) Username for the remote site

### Read-Only

- `id` (String) The ID of this resource.

