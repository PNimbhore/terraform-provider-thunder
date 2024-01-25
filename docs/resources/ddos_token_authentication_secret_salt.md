---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_token_authentication_secret_salt Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_token_authentication_secret_salt: Token Authentication Secret Salt
  PLACEHOLDER
---

# thunder_ddos_token_authentication_secret_salt (Resource)

`thunder_ddos_token_authentication_secret_salt`: Token Authentication Secret Salt

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_token_authentication_secret_salt" "thunder_ddos_token_authentication_secret_salt" {

  current_salt  = 4191031059
  previous_salt = 176240879

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `current_salt` (Number) Current salt value
- `previous_salt` (Number) Previous salt value
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

