---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_icmp_v4 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_icmp_v4: ICMPv4 template Configuration
  PLACEHOLDER
---

# thunder_ddos_template_icmp_v4 (Resource)

`thunder_ddos_template_icmp_v4`: ICMPv4 template Configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_icmp_v4" "thunder_ddos_template_icmp_v4" {
  icmp_tmpl_name = "20"
  type_list {
    type_number = 196
    type_rate   = 4400054
    code {
      code_number = 122
      code_rate   = 14668888
    }
    code_other {
      code_other_rate = 15963103
    }
    user_tag = "25"
  }
  type_other {
    type_other_deny = 1
  }
  user_tag = "71"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `icmp_tmpl_name` (String) DDOS ICMPv4 Template Name

### Optional

- `type_list` (Block List) (see [below for nested schema](#nestedblock--type_list))
- `type_other` (Block List, Max: 1) (see [below for nested schema](#nestedblock--type_other))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--type_list"></a>
### Nested Schema for `type_list`

Required:

- `type_number` (Number) Specify ICMP type number

Optional:

- `code` (Block List) (see [below for nested schema](#nestedblock--type_list--code))
- `code_other` (Block List, Max: 1) (see [below for nested schema](#nestedblock--type_list--code_other))
- `type_deny` (Number) Reject this ICMP type
- `type_rate` (Number) Specify the whole rate with this type
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--type_list--code"></a>
### Nested Schema for `type_list.code`

Optional:

- `code_number` (Number) Specify the ICMP code
- `code_rate` (Number) Specify the rate with the code


<a id="nestedblock--type_list--code_other"></a>
### Nested Schema for `type_list.code_other`

Optional:

- `code_other_rate` (Number) Specify rate with other code



<a id="nestedblock--type_other"></a>
### Nested Schema for `type_other`

Optional:

- `type_other_deny` (Number) Deny all other type
- `type_other_rate` (Number) Specify rate with other type
- `uuid` (String) uuid of the object

