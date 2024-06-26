---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_mlb Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_mlb: Configure mlb
  PLACEHOLDER
---

# thunder_slb_mlb (Resource)

`thunder_slb_mlb`: Configure mlb

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_mlb" "thunder_slb_mlb" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'client_msg_sent': Client message sent; 'server_msg_received': Server message received; 'server_conn_created': Server connection created; 'server_conn_rst': Server connection reset; 'server_conn_failed': Server connection failed; 'server_conn_closed': Server connection closed; 'client_conn_created': Client connection created; 'client_conn_closed': Client connection closed; 'client_conn_not_found': Client connection not found; 'msg_dropped': Message dropped; 'msg_rerouted': Message rerouted; 'mlb_dcmsg_sent': Dcmsg sent; 'mlb_dcmsg_received': Dcmsg received; 'mlb_dcmsg_error': Dcmsg error; 'mlb_dcmsg_alloc': Dcmsg alloc; 'mlb_dcmsg_free': Dcmsg free; 'mlb_server_probe': Server probe; 'mlb_server_down': Server down;


