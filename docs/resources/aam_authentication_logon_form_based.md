---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_logon_form_based Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_logon_form_based: Form-based Authentication Logon
  PLACEHOLDER
---

# thunder_aam_authentication_logon_form_based (Resource)

`thunder_aam_authentication_logon_form_based`: Form-based Authentication Logon

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_logon_form_based" "thunder_aam_authentication_logon_form_based" {
  challenge_variable = "1"
  cp_page_cfg {
    changepassword_url = "56"
    cp_user_enum       = "changepassword-username-variable"
    cp_user_var        = "49"
    cp_old_pwd_enum    = "changepassword-old-password-variable"
    cp_old_pwd_var     = "44"
    cp_new_pwd_enum    = "changepassword-new-password-variable"
    cp_new_pwd_var     = "35"
    cp_cfm_pwd_enum    = "changepassword-password-confirm-variable"
    cp_cfm_pwd_var     = "40"
  }
  csp_support {
    specificuri         = "37"
    optional_second_uri = "69"
  }
  duration     = 1800
  hsts_timeout = 298757643
  logon_page_cfg {
    action_url                   = "10"
    username_variable            = "44"
    password_variable            = "59"
    passcode_variable            = "60"
    captcha_variable             = "11"
    login_failure_message        = "11"
    authz_failure_message        = "68"
    disable_change_password_link = 0
  }
  name                = "28"
  new_pin_variable    = "41"
  next_token_variable = "1"
  notify_cp_page_cfg {
    notifychangepassword_change_url   = "59"
    notifychangepassword_continue_url = "25"
  }
  retry    = 3
  user_tag = "92"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify form-based authentication logon name

### Optional

- `account_lock` (Number) Lock the account when the failed logon attempts is exceeded
- `challenge_variable` (String) Specify challenge variable name in form submission
- `cp_page_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--cp_page_cfg))
- `csp_support` (Block List, Max: 1) (see [below for nested schema](#nestedblock--csp_support))
- `duration` (Number) The time an account remains locked in seconds (default 1800)
- `hsts_timeout` (Number) Set HSTS policy expired timeout in seconds, 0 means to disable HSTS policy
- `logon_page_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--logon_page_cfg))
- `new_pin_variable` (String) Specify new-pin variable name in form submission
- `next_token_variable` (String) Specify next-token variable name in form submission
- `notify_cp_page_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--notify_cp_page_cfg))
- `portal` (Block List, Max: 1) (see [below for nested schema](#nestedblock--portal))
- `retry` (Number) Maximum number of consecutive failed logon attempts (default 3)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--cp_page_cfg"></a>
### Nested Schema for `cp_page_cfg`

Optional:

- `changepassword_url` (String) Specify changepassword form submission action url (changepassword action url)
- `cp_cfm_pwd_enum` (String) 'changepassword-password-confirm-variable': Specify password confirm variable name in form submission;
- `cp_cfm_pwd_var` (String) Specify password confirm variable name
- `cp_new_pwd_enum` (String) 'changepassword-new-password-variable': Specify new password variable name in form submission;
- `cp_new_pwd_var` (String) Specify new password variable name
- `cp_old_pwd_enum` (String) 'changepassword-old-password-variable': Specify old password variable name in form submission;
- `cp_old_pwd_var` (String) Specify old password variable name
- `cp_user_enum` (String) 'changepassword-username-variable': Specify username variable name in form submission;
- `cp_user_var` (String) Specify username variable name


<a id="nestedblock--csp_support"></a>
### Nested Schema for `csp_support`

Optional:

- `none` (Number) Set CSP frame-ancestors to none (also X-Frame-Options deny)
- `optional_second_uri` (String) Set optional second customized CSP URI
- `self` (Number) Set CSP frame-ancestors to self (also X-Frame-Options same-origin)
- `specificuri` (String) Set customized CSP frame-ancestors (maximum 2 URIs can be set)


<a id="nestedblock--logon_page_cfg"></a>
### Nested Schema for `logon_page_cfg`

Optional:

- `action_url` (String) Specify form submission action url
- `authz_failure_message` (String) Specify authorization failure message shown in logon page (Specify error string, default is "Authorization failed. Please contact your system administrator.")
- `captcha_variable` (String) Specify captcha variable name in form submission
- `disable_change_password_link` (Number) Don't display change password link on logon page forcibly even backend authentication server supports it (LDAP or Kerberos)
- `login_failure_message` (String) Specify login failure message shown in logon page (Specify error string, default is "Invalid username or password. Please try again.")
- `passcode_variable` (String) Specify passcode variable name in form submission
- `password_variable` (String) Specify password variable name in form submission
- `username_variable` (String) Specify username variable name in form submission


<a id="nestedblock--notify_cp_page_cfg"></a>
### Nested Schema for `notify_cp_page_cfg`

Optional:

- `notifychangepassword_change_url` (String) Specify change password action url for notifychangepassword form
- `notifychangepassword_continue_url` (String) Specify continue action url for notifychangepassword form


<a id="nestedblock--portal"></a>
### Nested Schema for `portal`

Optional:

- `challenge_page` (String) Specify challenge page name for RSA-RADIUS
- `changepasswordpage` (String) Specify change password page name
- `default_portal` (Number) Use default portal
- `failpage` (String) Specify logon fail page name (portal fail page name)
- `logon` (String) Specify logon page name
- `new_pin_page` (String) Specify new PIN page name for RSA-RADIUS
- `next_token_page` (String) Specify next token page name for RSA-RADIUS
- `notifychangepasswordpage` (String) Specify change password notification page name
- `portal_name` (String) Specify portal name


