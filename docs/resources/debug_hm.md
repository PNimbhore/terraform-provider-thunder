---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_debug_hm Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_debug_hm: Debug health monitor
  PLACEHOLDER
---

# thunder_debug_hm (Resource)

`thunder_debug_hm`: Debug health monitor

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_hm" "thunder_debug_hm" {
  level       = 2
  method_type = "icmp"
  pin_uid     = 667918429
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `level` (Number) Debug level (Level 1-3)
- `method_type` (String) 'icmp': ICMP type; 'tcp': TCP type; 'udp': UDP type; 'ftp': FTP type; 'http': HTTP type; 'snmp': SNMP type; 'smtp': SMTP type; 'dns': DNS type; 'dns-tcp': DNS TCP type; 'pop3': POP3 type; 'imap': IMAP type; 'sip': SIP type; 'sip-tcp': SIP TCP type; 'radius': RADIUS type; 'ldap': LDAP type; 'rtsp': RTSP type; 'kerberos-kdc': Kerberos KDC type; 'database': DATABASE type; 'external': EXTERNAL type; 'https': HTTPS type; 'ntp': NTP type; 'compound': Compound type;
- `pin_uid` (Number) Debug Pin Unique Id
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


