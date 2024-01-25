---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_template_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_template_policy: Policy config
  PLACEHOLDER
---

# thunder_cgnv6_template_policy (Resource)

`thunder_cgnv6_template_policy`: Policy config

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_policy" "thunder_cgnv6_template_policy" {
  name     = "test1"
  user_tag = "57"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Policy template name

### Optional

- `class_list` (Block List, Max: 1) (see [below for nested schema](#nestedblock--class_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--class_list"></a>
### Nested Schema for `class_list`

Optional:

- `client_ip_l3_dest` (Number) Use destination IP as client IP address
- `client_ip_l7_header` (Number) Use extract client IP address from L7 header
- `header_name` (String) Specify L7 header name
- `lid_list` (Block List) (see [below for nested schema](#nestedblock--class_list--lid_list))
- `name` (String) Class list name
- `uuid` (String) uuid of the object

<a id="nestedblock--class_list--lid_list"></a>
### Nested Schema for `class_list.lid_list`

Required:

- `lidnum` (Number) Specify a limit ID

Optional:

- `action_value` (String) 'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;
- `conn_limit` (Number) Connection limit
- `conn_per` (Number) Per (Specify interval in number of 100ms)
- `conn_rate_limit` (Number) Specify connection rate limit
- `dns64` (Block List, Max: 1) (see [below for nested schema](#nestedblock--class_list--lid_list--dns64))
- `interval` (Number) Specify log interval in minutes, by default system will log every over limit instance
- `lockout` (Number) Don't accept any new connection for certain time (Lockout duration in minutes)
- `log` (Number) Log a message
- `over_limit_action` (Number) Set action when exceeds limit
- `request_limit` (Number) Request limit (Specify request limit)
- `request_per` (Number) Per (Specify interval in number of 100ms)
- `request_rate_limit` (Number) Request rate limit (Specify request rate limit)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--class_list--lid_list--dns64"></a>
### Nested Schema for `class_list.lid_list.dns64`

Optional:

- `disable` (Number) Disable
- `exclusive_answer` (Number) Exclusive Answer in DNS Response
- `prefix` (String) IPv6 prefix

