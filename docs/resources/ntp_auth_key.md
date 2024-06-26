---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ntp_auth_key Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ntp_auth_key: authentication key
  PLACEHOLDER
---

# thunder_ntp_auth_key (Resource)

`thunder_ntp_auth_key`: authentication key

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ntp_auth_key" "thunderNtpAuthKeyTest" {
  key      = 1
  alg_type = "SHA1"
  key_type = "ascii"
  asc_key  = "a"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key` (Number) authentication key

### Optional

- `alg_type` (String) 'M': encryption using MD5; 'SHA': encryption using SHA; 'SHA1': encryption using SHA1;
- `asc_key` (String)
- `hex_key` (String)
- `key_type` (String) 'ascii': key string in ASCII form; 'hex': key string in hex form;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


