---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_icmp_v4_type Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_icmp_v4_type: Specify ICMP type
  PLACEHOLDER
---

# thunder_ddos_template_icmp_v4_type (Resource)

`thunder_ddos_template_icmp_v4_type`: Specify ICMP type

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v4_type" "thunder_ddos_template_icmp_v4_type" {
  icmp_tmpl_name = "20"
  code {
    code_number = 167
    code_rate   = 1248190
  }
  code_other {
    code_other_rate = 14325190
  }
  type_number = 230
  type_rate   = 14374233
  user_tag    = "66"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `icmp_tmpl_name` (String) IcmpTmplName
- `type_number` (Number) Specify ICMP type number

### Optional

- `code` (Block List) (see [below for nested schema](#nestedblock--code))
- `code_other` (Block List, Max: 1) (see [below for nested schema](#nestedblock--code_other))
- `type_deny` (Number) Reject this ICMP type
- `type_rate` (Number) Specify the whole rate with this type
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--code"></a>
### Nested Schema for `code`

Optional:

- `code_number` (Number) Specify the ICMP code
- `code_rate` (Number) Specify the rate with the code


<a id="nestedblock--code_other"></a>
### Nested Schema for `code_other`

Optional:

- `code_other_rate` (Number) Specify rate with other code


