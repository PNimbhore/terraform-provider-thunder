---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_relay_kerberos_instance_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_relay_kerberos_instance_stats: Statistics for the object instance
  PLACEHOLDER
---

# thunder_aam_authentication_relay_kerberos_instance_stats (Data Source)

`thunder_aam_authentication_relay_kerberos_instance_stats`: Statistics for the object instance

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_relay_kerberos_instance_stats" "thunder_aam_authentication_relay_kerberos_instance_stats" {

  name = "test"
}
output "get_aam_authentication_relay_kerberos_instance_stats" {
  value = ["${data.thunder_aam_authentication_relay_kerberos_instance_stats.thunder_aam_authentication_relay_kerberos_instance_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify Kerberos authentication relay name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `current_requests_of_user` (Number) Current Pending Requests of User
- `request_send` (Number) Request Send
- `response_receive` (Number) Response Receive
- `tickets` (Number) Tickets


