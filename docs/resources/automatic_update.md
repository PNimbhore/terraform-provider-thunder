---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_automatic_update Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_automatic_update: Automatic update configuration
  PLACEHOLDER
---

# thunder_automatic_update (Resource)

`thunder_automatic_update`: Automatic update configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_automatic_update" "thunder_automatic_update" {
  config_list {
    feature_name = "ca-bundle"
    schedule     = 1
    weekly       = 1
    week_day     = "Wednesday"
    week_time    = "11:11"
  }
  proxy_server {
    proxy_host    = "85"
    https_port    = 57158
    auth_type     = "basic"
    username      = "69"
    password      = 1
    secret_string = "asd"
  }
  use_mgmt_port = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `check_now` (Block List, Max: 1) (see [below for nested schema](#nestedblock--check_now))
- `checknow` (Block List, Max: 1) (see [below for nested schema](#nestedblock--checknow))
- `config_list` (Block List) (see [below for nested schema](#nestedblock--config_list))
- `glm_source_url` (String) Change GLM source url
- `info` (Block List, Max: 1) (see [below for nested schema](#nestedblock--info))
- `proxy_server` (Block List, Max: 1) (see [below for nested schema](#nestedblock--proxy_server))
- `reset` (Block List, Max: 1) (see [below for nested schema](#nestedblock--reset))
- `revert` (Block List, Max: 1) (see [below for nested schema](#nestedblock--revert))
- `use_mgmt_port` (Number) Use management port to connect
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--check_now"></a>
### Nested Schema for `check_now`

Optional:

- `feature_name` (String) 'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;
- `from_staging_server` (Number) Get files from GLM Staging storage
- `prod_ver` (String) update to this specific version, if this option is not configured, update to the latest version
- `stage_ver` (String) update this specific version


<a id="nestedblock--checknow"></a>
### Nested Schema for `checknow`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--config_list"></a>
### Nested Schema for `config_list`

Required:

- `feature_name` (String) 'app-fw': Application Firewall Configuration; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;

Optional:

- `daily` (Number) Every day
- `day_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `debug` (Number) Enable libcurl debug option
- `disable_ssl_verify` (Number) Disable peer server certificate verification
- `schedule` (Number)
- `uuid` (String) uuid of the object
- `week_day` (String) 'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;
- `week_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `weekly` (Number) Every week


<a id="nestedblock--info"></a>
### Nested Schema for `info`

Optional:

- `uuid` (String) uuid of the object


<a id="nestedblock--proxy_server"></a>
### Nested Schema for `proxy_server`

Optional:

- `auth_type` (String) 'ntlm': NTLM authentication(default); 'basic': Basic authentication;
- `domain` (String) Realm for NTLM authentication
- `https_port` (Number) Proxy server HTTPs port
- `password` (Number) Password for proxy authentication
- `proxy_host` (String) Proxy server hostname or IP address
- `secret_string` (String) password value
- `username` (String) Username for proxy authentication
- `uuid` (String) uuid of the object


<a id="nestedblock--reset"></a>
### Nested Schema for `reset`

Optional:

- `feature_name` (String) 'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle;


<a id="nestedblock--revert"></a>
### Nested Schema for `revert`

Optional:

- `feature_name` (String) 'app-fw': Application Firewall; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;


