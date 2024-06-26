---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_threat_intel_webroot_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_threat_intel_webroot_global: Global counters for webroot module
  PLACEHOLDER
---

# thunder_threat_intel_webroot_global (Resource)

`thunder_threat_intel_webroot_global`: Global counters for webroot module

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_threat_intel_webroot_global" "thunder_threat_intel_webroot_global" {
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

- `counters1` (String) 'all': all; 'spam-sources': Hits for spam sources; 'windows-exploits': Hits for windows exploits; 'web-attacks': Hits for web attacks; 'botnets': Hits for botnets; 'scanners': Hits for scanners; 'dos-attacks': Hits for dos attacks; 'reputation': Hits for reputation; 'phishing': Hits for phishing; 'proxy': Hits for proxy; 'mobile-threats': Hits for mobile threats; 'tor-proxy': Hits for tor-proxy; 'rtu-lookup': Number of lookups in RTU cache; 'database-lookup': Number of lookups in database; 'non-malicious-ips': IP's not found in database or RTU cache;


