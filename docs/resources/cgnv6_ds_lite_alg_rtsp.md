---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_ds_lite_alg_rtsp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_ds_lite_alg_rtsp: DS-Lite RTSP ALG (default: disabled)
  PLACEHOLDER
---

# thunder_cgnv6_ds_lite_alg_rtsp (Resource)

`thunder_cgnv6_ds_lite_alg_rtsp`: DS-Lite RTSP ALG (default: disabled)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_ds_lite_alg_rtsp" "thunder_cgnv6_ds_lite_alg_rtsp" {
  rtsp_enable = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `rtsp_enable` (String) 'enable': Enable ALG;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


