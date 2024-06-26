---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_oauth_authorization_server Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_oauth_authorization_server: Authentication 2.0 authorization server
  PLACEHOLDER
---

# thunder_aam_authentication_oauth_authorization_server (Resource)

`thunder_aam_authentication_oauth_authorization_server`: Authentication 2.0 authorization server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_oauth_authorization_server" "thunder_aam_authentication_oauth_authorization_server" {
  authorization_endpoint = "90"
  client_method          = "ignored"
  issuer                 = "8"
  name                   = "30"
  sampling_enable {
    counters1 = "all"
  }
  server_method  = "post"
  token_endpoint = "73"
  user_tag       = "33"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify authorization server object name

### Optional

- `authorization_endpoint` (String) Specify URI for authorization
- `client_method` (String) 'ignored': Clients' browser will send data according to server spec (default); 'post': Clients' browser will send data by POST; 'get': Clients' browser will send data by GET;
- `issuer` (String) Specify openid provider name for authorization
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `server_method` (String) 'post': AX will send data to server by POST (default); 'get': AX will send data to server by GET;
- `token_endpoint` (String) Specify URI for token exchange
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `verification_cert` (String) Specify certificate to verify ID token signature
- `verification_jwks` (String) Specify jwks file to verify ID token signature

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'auth-req': auth-req; 'auth-succ': auth-succ; 'auth-fail': auth-fail; 'auth-error': auth-error; 'other-error': other-error;


