---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_gtp_msisdn_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_gtp_msisdn_list: Configure GTP MSISDN list
  PLACEHOLDER
---

# thunder_template_gtp_msisdn_list (Resource)

`thunder_template_gtp_msisdn_list`: Configure GTP MSISDN list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_msisdn_list" "thunder_template_gtp_msisdn_list" {
  name   = "test"
  action = "permit"
  str_list {
    msisdn = "20"
  }
  user_tag = "test_user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify name of the GTP MSISDN list

### Optional

- `action` (String) 'permit': Create a whitelist to permit the packets that match MSISDN filters; 'deny': Create a blacklist to deny the packets that match MSISDN filters;
- `str_list` (Block List) (see [below for nested schema](#nestedblock--str_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--str_list"></a>
### Nested Schema for `str_list`

Optional:

- `msisdn` (String) Specify the MSISDN filter

