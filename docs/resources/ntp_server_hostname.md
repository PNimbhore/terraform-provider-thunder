---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ntp_server_hostname Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ntp_server_hostname: IPV4 address, IPV6 address or host name of NTP server
  PLACEHOLDER
---

# thunder_ntp_server_hostname (Resource)

`thunder_ntp_server_hostname`: IPV4 address, IPV6 address or host name of NTP server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ntp_server_hostname" "thunder_ntp_server_hostname" {

  action          = "enable"
  host_servername = "5"
  key             = 41809
  prefer          = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host_servername` (String) IPV4 address, IPV6 address or host name of NTP server(string1~63)

### Optional

- `action` (String) 'enable': Enable this NTP server; 'disable': Disable this NTP server;
- `key` (Number) Use trusted key to authenticate this NTP server (The trusted key number to use)
- `prefer` (Number) Set this NTP server as preferred
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

