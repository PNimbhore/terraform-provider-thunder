---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service_geo_location Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service_geo_location: Geo location settings
  PLACEHOLDER
---

# thunder_gslb_zone_service_geo_location (Resource)

`thunder_gslb_zone_service_geo_location`: Geo location settings

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_geo_location" "thunder_gslb_zone_service_geo_location" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  action       = 1
  action_type  = "forward"
  forward_type = "both"
  geo_name     = "116"
  user_tag     = "79"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `geo_name` (String) Specify the geo-location
- `name` (String) Name
- `service_name` (String) ServiceName
- `service_port` (String) ServicePort

### Optional

- `action` (Number) Action for this geo-location
- `action_type` (String) 'allow': Allow query from this geo-location; 'drop': Drop query from this geo-location; 'forward': Forward packet for this geo-location; 'ignore': Send empty response to this geo-location; 'reject': Send refuse response to this geo-location;
- `alias` (Block List) (see [below for nested schema](#nestedblock--alias))
- `forward_type` (String) 'both': Forward both query and response; 'query': Forward query from this geo-location; 'response': Forward response to this geo-location;
- `policy` (String) Policy for this geo-location (Specify the policy name)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--alias"></a>
### Nested Schema for `alias`

Optional:

- `alias` (String) Send CNAME response for this geo-location (Specify a CNAME record)


