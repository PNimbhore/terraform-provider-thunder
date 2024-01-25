---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_portal_logon Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_portal_logon: Logon page configuration
  PLACEHOLDER
---

# thunder_aam_authentication_portal_logon (Resource)

`thunder_aam_authentication_portal_logon`: Logon page configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_logon" "thunder_aam_authentication_portal_logon" {

  name       = "default-portal"
  action_url = "16"
  background {
    bgfile  = "4"
    bgstyle = "tile"
  }
  username_var = "8"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `action_url` (String) Specify form action URL in default logon page (Default: /logon.fo)
- `background` (Block List, Max: 1) (see [below for nested schema](#nestedblock--background))
- `captcha_type` (String) 'reCAPTCHAv2-checkbox': Google reCAPTCHAv2 Checkbox; 'reCAPTCHAv2-invisible': Google reCAPTCHAv2 Invisible; 'reCAPTCHAv3': Google reCAPTCHAv3;
- `enable_captcha` (Number) Enable CAPTCHA in deafult logon page
- `enable_passcode` (Number) Enable passcode field in default logon page
- `fail_msg_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--fail_msg_cfg))
- `passcode_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--passcode_cfg))
- `passcode_var` (String) Specify passcode variable name in default logon page (Default: passcode)
- `password_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--password_cfg))
- `password_var` (String) Specify password variable name in default logon page (Default: pwd)
- `recaptcha_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--recaptcha_cfg))
- `site_key_string` (String) Site key string
- `submit_text` (String) Specify submit button text in default logon page (Default: Log In)
- `username_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--username_cfg))
- `username_var` (String) Specify username variable name in default logon page (Default: user)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--background"></a>
### Nested Schema for `background`

Optional:

- `bgcolor_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `bgcolor_value` (String) Specify 6-digit HEX color value
- `bgfile` (String) Specify background image filename
- `bgstyle` (String) 'tile': Tile; 'stretch': Stretch; 'fit': Fit;


<a id="nestedblock--fail_msg_cfg"></a>
### Nested Schema for `fail_msg_cfg`

Optional:

- `authz_fail_msg` (String) Configure authorization failure message in default logon page, its text attributes follow fail-msg's (Specify authorization failure message (Default: Authorization failed. Please contact your system administrator.))
- `fail_color` (Number) Specify font color (Default: red)
- `fail_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `fail_color_value` (String) Specify 6-digit HEX color value
- `fail_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `fail_font` (Number) Sepcify font (Default: Arial)
- `fail_font_custom` (String) Specify custom font
- `fail_msg` (Number) Configure login failure message in default logon page
- `fail_size` (Number) Specify font size (Default: 5)
- `fail_text` (String) Specify login failure message (Default: Invalid username or password. Please try again.)


<a id="nestedblock--passcode_cfg"></a>
### Nested Schema for `passcode_cfg`

Optional:

- `passcode` (Number) Configure passcode text in default logon page
- `passcode_color` (Number) Specify font color (Default: black)
- `passcode_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `passcode_color_value` (String) Specify 6-digit HEX color value
- `passcode_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `passcode_font` (Number) Sepcify font (Default: Arial)
- `passcode_font_custom` (String) Specify custom font
- `passcode_size` (Number) Specify font size (Default: 3)
- `passcode_text` (String) Specify passcode text (Default: Passcode)


<a id="nestedblock--password_cfg"></a>
### Nested Schema for `password_cfg`

Optional:

- `pass_color` (Number) Specify font color (Default: black)
- `pass_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `pass_color_value` (String) Specify 6-digit HEX color value
- `pass_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `pass_font` (Number) Sepcify font (Default: Arial)
- `pass_font_custom` (String) Specify custom font
- `pass_size` (Number) Specify font size (Default: 3)
- `pass_text` (String) Specify password text (Default: Password)
- `password` (Number) Configure password text in default logon page


<a id="nestedblock--recaptcha_cfg"></a>
### Nested Schema for `recaptcha_cfg`

Optional:

- `recaptcha_action` (String) Specify reCAPTCHA action (Specify action string, only accept alphanumeric, underscore, and slash (Default: A10_DEFAULT_LOGON))
- `recaptcha_badge` (String) 'bottom-left': bottom left corner; 'bottom-right': bottom right corner;
- `recaptcha_size` (String) 'normal': normal size; 'compact': compact size;
- `recaptcha_theme` (String) 'light': light theme; 'dark': dark theme;


<a id="nestedblock--username_cfg"></a>
### Nested Schema for `username_cfg`

Optional:

- `user_color` (Number) Specify font color (Default: black)
- `user_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `user_color_value` (String) Specify 6-digit HEX color value
- `user_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `user_font` (Number) Sepcify font (Default: Arial)
- `user_font_custom` (String) Specify custom font
- `user_size` (Number) Specify font size (Default: 3)
- `user_text` (String) Specify username text (Default: User Name)
- `username` (Number) Configure username text in default logon page

