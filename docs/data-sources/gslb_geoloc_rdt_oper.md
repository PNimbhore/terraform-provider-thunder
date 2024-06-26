---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_geoloc_rdt_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_geoloc_rdt_oper: Operational Status for the object geoloc-rdt
  PLACEHOLDER
---

# thunder_gslb_geoloc_rdt_oper (Data Source)

`thunder_gslb_geoloc_rdt_oper`: Operational Status for the object geoloc-rdt

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_geoloc_rdt_oper" "thunder_gslb_geoloc_rdt_oper" {
}
output "get_gslb_geoloc_rdt_oper" {
  value = ["${data.thunder_gslb_geoloc_rdt_oper.thunder_gslb_geoloc_rdt_oper}"]
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

- `active_status` (String)
- `geo_name` (String)
- `geoloc_rdt_list` (Block List) (see [below for nested schema](#nestedblock--oper--geoloc_rdt_list))
- `site_name` (String)
- `total_rdt` (Number)

<a id="nestedblock--oper--geoloc_rdt_list"></a>
### Nested Schema for `oper.geoloc_rdt_list`

Optional:

- `age` (Number)
- `gl_name` (String)
- `rdt` (Number)
- `site_name` (String)
- `type` (String)


