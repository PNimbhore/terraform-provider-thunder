---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_ssl_ca_cert_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_ssl_ca_cert_oper: Operational Status for the object ssl-ca-cert
  PLACEHOLDER
---

# thunder_slb_ssl_ca_cert_oper (Data Source)

`thunder_slb_ssl_ca_cert_oper`: Operational Status for the object ssl-ca-cert

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_ssl_ca_cert_oper" "thunder_slb_ssl_ca_cert_oper" {
}
output "get_slb_ssl_ca_cert_oper" {
  value = ["${data.thunder_slb_ssl_ca_cert_oper.thunder_slb_ssl_ca_cert_oper}"]
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

- `exact_match` (Number)
- `partition` (String)
- `sortby_exp` (Number)
- `sortby_name` (Number)
- `ssl_certs` (Block List) (see [below for nested schema](#nestedblock--oper--ssl_certs))

<a id="nestedblock--oper--ssl_certs"></a>
### Nested Schema for `oper.ssl_certs`

Optional:

- `common_name` (String)
- `issuer` (String)
- `keysize` (Number)
- `name` (String)
- `notafter` (String)
- `notafter_number` (Number)
- `notbefore` (String)
- `organization` (String)
- `serial` (String)
- `status` (String)
- `subject` (String)
- `type` (String)


