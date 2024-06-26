---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_diameter Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_diameter: diameter template
  PLACEHOLDER
---

# thunder_slb_template_diameter (Resource)

`thunder_slb_template_diameter`: diameter template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_diameter" "test_thunder_slb_template_diameter" {
  name                       = "testing_diameter"
  customize_cea              = 1
  avp_code                   = 122
  avp_string                 = "test"
  service_group_name         = "sg1"
  dwr_time                   = 100
  idle_timeout               = 5
  multiple_origin_host       = 1
  origin_realm               = "test_realm"
  product_name               = "test_product"
  vendor_id                  = 12345
  session_age                = 456
  dwr_up_retry               = 2
  terminate_on_cca_t         = 1
  forward_unknown_session_id = 1
  forward_to_latest_server   = 1
  load_balance_on_session_id = 1
  message_code_list {
    message_code = 123
  }
  avp_list {
    avp       = 123
    int32     = 12
    mandatory = 1
  }
  origin_host {
    origin_host_name = "test_host"
  }
  user_tag = "tet_user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) diameter template Name

### Optional

- `avp_code` (Number) avp code
- `avp_list` (Block List) (see [below for nested schema](#nestedblock--avp_list))
- `avp_string` (String) pattern to be matched in the avp string name, max length 127 bytes
- `customize_cea` (Number) customizing cea response
- `dwr_time` (Number) dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option)
- `dwr_up_retry` (Number) number of successful dwr health-check before declaring target up
- `forward_to_latest_server` (Number) Forward client message to the latest server that sends message with the same session id
- `forward_unknown_session_id` (Number) Forward server message even it has unknown session id
- `idle_timeout` (Number) user sesison idle timeout (in minutes, default is 5)
- `load_balance_on_session_id` (Number) Load balance based on the session id
- `message_code_list` (Block List) (see [below for nested schema](#nestedblock--message_code_list))
- `multiple_origin_host` (Number) allowing multiple origin-host to a single server
- `origin_host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--origin_host))
- `origin_realm` (String) origin-realm name avp
- `product_name` (String) product name avp
- `relaxed_origin_host` (Number) Relaxed Origin-Host Format
- `service_group_name` (String) service group name, this is the service group that the message needs to be copied to
- `session_age` (Number) user session age allowed (default 10), this is not idle-time (in minutes)
- `terminate_on_cca_t` (Number) remove diameter session when receiving CCA-T message
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `vendor_id` (Number) vendor-id avp (Vendor Id)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--avp_list"></a>
### Nested Schema for `avp_list`

Optional:

- `avp` (Number) customize avps for cer to the server (avp number)
- `int32` (Number) 32 bits integer
- `int64` (Number) 64 bits integer
- `mandatory` (Number) mandatory avp
- `string` (String) String (string name, max length 127 bytes)


<a id="nestedblock--message_code_list"></a>
### Nested Schema for `message_code_list`

Optional:

- `message_code` (Number)


<a id="nestedblock--origin_host"></a>
### Nested Schema for `origin_host`

Optional:

- `origin_host_name` (String) origin-host name avp
- `uuid` (String) uuid of the object


