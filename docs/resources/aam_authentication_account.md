---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_account Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_account: Authentication account configuration
  PLACEHOLDER
---

# thunder_aam_authentication_account (Resource)

`thunder_aam_authentication_account`: Authentication account configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_account" "thunder_aam_authentication_account" {

  kerberos_spn_list {
    name                   = "41"
    realm                  = "36"
    account                = "4"
    service_principal_name = "34"
    password               = 0
    user_tag               = "95"
  }
  sampling_enable {
    counters1 = "all"
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `kerberos_spn_list` (Block List) (see [below for nested schema](#nestedblock--kerberos_spn_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--kerberos_spn_list"></a>
### Nested Schema for `kerberos_spn_list`

Required:

- `name` (String) Specify AD account name

Optional:

- `account` (String) Specify domain account for SPN
- `password` (Number) Specify password of domain account
- `realm` (String) Specify Kerberos realm
- `secret_string` (String) Password of AD account
- `service_principal_name` (String) Specify service principal name
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response;


