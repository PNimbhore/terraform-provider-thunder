---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_interface_http_health_check Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_interface_http_health_check: Use DST Interface-ip entry as HTTP health-check target
  PLACEHOLDER
---

# thunder_ddos_interface_http_health_check (Resource)

`thunder_ddos_interface_http_health_check`: Use DST Interface-ip entry as HTTP health-check target

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_interface_http_health_check" "thunder_ddos_interface_http_health_check" {
  challenge_method        = "http-redirect"
  challenge_redirect_code = "302"
  enable                  = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable` (String) 'enable': enable;

### Optional

- `challenge_method` (String) 'http-redirect': http-redirect; 'javascript': javascript;
- `challenge_redirect_code` (String) '302': 302 Found; '307': 307 Temporary Redirect;
- `challenge_uri_encode` (Number) Encode the challenge phrase in uri instead of in http cookie. Default encoded in http cookie
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


