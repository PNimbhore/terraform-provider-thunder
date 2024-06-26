---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_port_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_port_list: Configure port list
  PLACEHOLDER
---

# thunder_cgnv6_port_list (Resource)

`thunder_cgnv6_port_list`: Configure port list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_port_list" "thunder_cgnv6_port_list" {
  name = "test"
  port_config {
    original_port   = 55532
    translated_port = 26432
  }
  user_tag = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify name of the port list

### Optional

- `port_config` (Block List) (see [below for nested schema](#nestedblock--port_config))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--port_config"></a>
### Nested Schema for `port_config`

Optional:

- `original_port` (Number) Original port to be translated
- `translated_port` (Number) Port after translation


