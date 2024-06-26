---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_mqtt_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_mqtt_oper: Operational Status for the object mqtt
  PLACEHOLDER
---

# thunder_slb_mqtt_oper (Data Source)

`thunder_slb_mqtt_oper`: Operational Status for the object mqtt

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_mqtt_oper" "thunder_slb_mqtt_oper" {
}
output "get_slb_mqtt_oper" {
  value = ["${data.thunder_slb_mqtt_oper.thunder_slb_mqtt_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `cpu_count` (Number)
- `mqtt_cpu_list` (Block List) (see [below for nested schema](#nestedblock--oper--mqtt_cpu_list))

<a id="nestedblock--oper--mqtt_cpu_list"></a>
### Nested Schema for `oper.mqtt_cpu_list`

Optional:

- `client_id_null` (Number)
- `conn_null` (Number)
- `curr_proxy` (Number)
- `insertion_failed` (Number)
- `insertion_successful` (Number)
- `parse_connect_fail` (Number)
- `parse_publish_fail` (Number)
- `parse_subscribe_fail` (Number)
- `parse_unsubscribe_fail` (Number)
- `recv_mqtt_auth` (Number)
- `recv_mqtt_connack` (Number)
- `recv_mqtt_connect` (Number)
- `recv_mqtt_disconnect` (Number)
- `recv_mqtt_other` (Number)
- `recv_mqtt_pingreq` (Number)
- `recv_mqtt_pingresp` (Number)
- `recv_mqtt_puback` (Number)
- `recv_mqtt_pubcomp` (Number)
- `recv_mqtt_publish` (Number)
- `recv_mqtt_pubrec` (Number)
- `recv_mqtt_pubrel` (Number)
- `recv_mqtt_suback` (Number)
- `recv_mqtt_subscribe` (Number)
- `recv_mqtt_unsuback` (Number)
- `recv_mqtt_unsubscribe` (Number)
- `request` (Number)
- `session_exist` (Number)
- `total_proxy` (Number)
- `tuple_already_linked` (Number)
- `tuple_not_linked` (Number)


