---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_web_service_secure Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_web_service_secure: Web-service secure operation
  PLACEHOLDER
---

# thunder_web_service_secure (Resource)

`thunder_web_service_secure`: Web-service secure operation

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_service_secure" "thunder_web_service_secure" {
  restart = 1
  wipe    = 1
  generate {
    domain_name = "test.com"
    country     = "IN"
    state       = "MH"
  }
  regenerate {
    domain_name = "test.com"
    country     = "IN"
    state       = "MH"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `certificate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--certificate))
- `generate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--generate))
- `private_key` (Block List, Max: 1) (see [below for nested schema](#nestedblock--private_key))
- `regenerate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--regenerate))
- `restart` (Number) Restart WEB service
- `wipe` (Number) Wipe WEB private-key and certificate

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--certificate"></a>
### Nested Schema for `certificate`

Optional:

- `file_url` (String) File URL
- `load` (Number) Load WEB certificate
- `use_mgmt_port` (Number) Use management port as source port


<a id="nestedblock--generate"></a>
### Nested Schema for `generate`

Optional:

- `country` (String) The country name
- `domain_name` (String) The domain name
- `state` (String) The location


<a id="nestedblock--private_key"></a>
### Nested Schema for `private_key`

Optional:

- `file_url` (String) File URL
- `load` (Number) Load WEB private-key
- `use_mgmt_port` (Number) Use management port as source port


<a id="nestedblock--regenerate"></a>
### Nested Schema for `regenerate`

Optional:

- `country` (String) The country name
- `domain_name` (String) The domain name
- `state` (String) The location


