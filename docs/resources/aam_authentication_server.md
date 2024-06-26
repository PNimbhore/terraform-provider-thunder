---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_server Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_server: Authentication server configuration
  PLACEHOLDER
---

# thunder_aam_authentication_server (Resource)

`thunder_aam_authentication_server`: Authentication server configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server" "thunder_aam_authentication_server" {

  ocsp {
    sampling_enable {
      counters1 = "all"
    }
    instance_list {
      name         = "test"
      url          = "http://192.168.0.1:80/"
      http_version = 1
      version_type = "1.1"
      sampling_enable {
        counters1 = "all"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ldap` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ldap))
- `ocsp` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ocsp))
- `radius` (Block List, Max: 1) (see [below for nested schema](#nestedblock--radius))
- `uuid` (String) uuid of the object
- `windows` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ldap"></a>
### Nested Schema for `ldap`

Optional:

- `instance_list` (Block List) (see [below for nested schema](#nestedblock--ldap--instance_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--ldap--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--ldap--instance_list"></a>
### Nested Schema for `ldap.instance_list`

Required:

- `name` (String) Specify LDAP authentication server name

Optional:

- `admin_dn` (String) The LDAP server's admin DN
- `admin_secret` (Number) Specify the LDAP server's admin secret password
- `auth_type` (String) 'ad': Active Directory. Default; 'open-ldap': OpenLDAP;
- `base` (String) Specify the LDAP server's search base
- `bind_with_dn` (Number) Enforce using DN for LDAP binding(All user input name will be used to create DN)
- `ca_cert` (String) Specify the LDAPS CA cert filename (Trusted LDAPS CA cert filename)
- `default_domain` (String) Specify default domain for LDAP
- `derive_bind_dn` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ldap--instance_list--derive_bind_dn))
- `dn_attribute` (String) Specify Distinguished Name attribute, default is CN
- `health_check` (Number) Check server's health status
- `health_check_disable` (Number) Disable configured health check configuration
- `health_check_string` (String) Health monitor name
- `host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ldap--instance_list--host))
- `ldaps_conn_reuse_idle_timeout` (Number) Specify LDAPS connection reuse idle timeout value (in seconds) (Specify idle timeout value (in seconds), default is 0 (not reuse LDAPS connection))
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `port` (Number) Specify the LDAP server's authentication port, default is 389
- `port_hm` (String) Check port's health status
- `port_hm_disable` (Number) Disable configured port health check configuration
- `prompt_pw_change_before_exp` (Number) Prompt user to change password before expiration in N days. This option only takes effect when server type is AD (Prompt user to change password before expiration in N days, default is not to prompt the user)
- `protocol` (String) 'ldap': Use LDAP (default); 'ldaps': Use LDAP over SSL; 'starttls': Use LDAP StartTLS;
- `pwdmaxage` (Number) Specify the LDAP server's default password expiration time (in seconds) (The LDAP server's default password expiration time (in seconds), default is 0 (no expiration))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--ldap--instance_list--sampling_enable))
- `secret_string` (String) secret password
- `timeout` (Number) Specify timout for LDAP, default is 10 seconds (The timeout, default is 10 seconds)
- `uuid` (String) uuid of the object

<a id="nestedblock--ldap--instance_list--derive_bind_dn"></a>
### Nested Schema for `ldap.instance_list.derive_bind_dn`

Optional:

- `username_attr` (String) Specify attribute name of username


<a id="nestedblock--ldap--instance_list--host"></a>
### Nested Schema for `ldap.instance_list.host`

Optional:

- `hostip` (String) Server's hostname(Length 1-31) or IP address
- `hostipv6` (String) Server's IPV6 address


<a id="nestedblock--ldap--instance_list--sampling_enable"></a>
### Nested Schema for `ldap.instance_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'admin-bind-success': Admin Bind Success; 'admin-bind-failure': Admin Bind Failure; 'bind-success': User Bind Success; 'bind-failure': User Bind Failure; 'search-success': Search Success; 'search-failure': Search Failure; 'authorize-success': Authorization Success; 'authorize-failure': Authorization Failure; 'timeout-error': Timeout; 'other-error': Other Error; 'request': Request; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'pw_expiry': Password expiry; 'pw_change_success': Password change success; 'pw_change_failure': Password change failure;



<a id="nestedblock--ldap--sampling_enable"></a>
### Nested Schema for `ldap.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'admin-bind-success': Total Admin Bind Success; 'admin-bind-failure': Total Admin Bind Failure; 'bind-success': Total User Bind Success; 'bind-failure': Total User Bind Failure; 'search-success': Total Search Success; 'search-failure': Total Search Failure; 'authorize-success': Total Authorization Success; 'authorize-failure': Total Authorization Failure; 'timeout-error': Total Timeout; 'other-error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'ldaps-idle-conn-num': LDAPS Idle Connection Number; 'ldaps-inuse-conn-num': LDAPS In-use Connection Number; 'pw-expiry': Total Password expiry; 'pw-change-success': Total password change success; 'pw-change-failure': Total password change failure;



<a id="nestedblock--ocsp"></a>
### Nested Schema for `ocsp`

Optional:

- `instance_list` (Block List) (see [below for nested schema](#nestedblock--ocsp--instance_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--ocsp--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--ocsp--instance_list"></a>
### Nested Schema for `ocsp.instance_list`

Required:

- `name` (String) Specify OCSP authentication server name

Optional:

- `health_check` (Number) Check server's health status
- `health_check_disable` (Number) Disable configured health check configuration
- `health_check_string` (String) Health monitor name
- `http_version` (Number) Set HTTP version (default 1.0)
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `port_health_check` (String) Check port's health status
- `port_health_check_disable` (Number) Disable configured port health check configuration
- `responder_ca` (String) Specify the trusted OCSP responder's CA cert filename
- `responder_cert` (String) Specify the trusted OCSP responder's cert filename
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--ocsp--instance_list--sampling_enable))
- `url` (String) Specify the OCSP server's address (Format: http://host[:port]/) (The OCSP server's address(Format: http://host[:port]/))
- `uuid` (String) uuid of the object
- `version_type` (String) '1.1': HTTP version 1.1;

<a id="nestedblock--ocsp--instance_list--sampling_enable"></a>
### Nested Schema for `ocsp.instance_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'request': Request; 'certificate-good': Good Certificate Response; 'certificate-revoked': Revoked Certificate Response; 'certificate-unknown': Unknown Certificate Response; 'timeout': Timeout; 'fail': Handle OCSP response failed; 'stapling-request': OCSP Stapling Request Send; 'stapling-certificate-good': OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': OCSP Stapling Unknown Certificate Response; 'stapling-timeout': OCSP Stapling Timeout; 'stapling-fail': Handle OCSP response failed;



<a id="nestedblock--ocsp--sampling_enable"></a>
### Nested Schema for `ocsp.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'stapling-certificate-good': Total OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': Total OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': Total OCSP Stapling Unknown Certificate Response; 'stapling-request-normal': Total OSCP Stapling Normal Request; 'stapling-request-dropped': Total OCSP Stapling Dropped Request; 'stapling-response-success': Total OCSP Stapling Success Response; 'stapling-response-failure': Total OCSP Stapling Failure Response; 'stapling-response-error': Total OCSP Stapling Error Response; 'stapling-response-timeout': Total OCSP Stapling Timeout Response; 'stapling-response-other': Total OCSP Stapling Other Response; 'request-normal': Total OSCP Normal Request; 'request-dropped': Total OCSP Dropped Request; 'response-success': Total OCSP Success Response; 'response-failure': Total OCSP Failure Response; 'response-error': Total OCSP Error Response; 'response-timeout': Total OCSP Timeout Response; 'response-other': Total OCSP Other Response; 'job-start-error': Total OCSP Job Start Error; 'polling-control-error': Total OCSP Polling Control Error;



<a id="nestedblock--radius"></a>
### Nested Schema for `radius`

Optional:

- `instance_list` (Block List) (see [below for nested schema](#nestedblock--radius--instance_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--radius--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--radius--instance_list"></a>
### Nested Schema for `radius.instance_list`

Required:

- `name` (String) Specify RADIUS authentication server name

Optional:

- `accounting_port` (Number) Specify the RADIUS server's accounting port, default is 1813
- `acct_port_hm` (String) Specify accounting port health check method
- `acct_port_hm_disable` (Number) Disable configured accounting port health check configuration
- `auth_type` (String) 'pap': PAP authentication. Default; 'mschapv2': MS-CHAPv2 authentication; 'mschapv2-pap': Use MS-CHAPv2 first. If server doesn't support it, try PAP;
- `health_check` (Number) Check server's health status
- `health_check_disable` (Number) Disable configured health check configuration
- `health_check_string` (String) Health monitor name
- `host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--radius--instance_list--host))
- `interval` (Number) Specify the interval time for resend the request (second), default is 3 seconds (The interval time(second), default is 3 seconds)
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `port` (Number) Specify the RADIUS server's authentication port, default is 1812
- `port_hm` (String) Check port's health status
- `port_hm_disable` (Number) Disable configured port health check configuration
- `retry` (Number) Specify the retry number for resend the request, default is 5 (The retry number, default is 5)
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--radius--instance_list--sampling_enable))
- `secret` (Number) Specify the RADIUS server's secret
- `secret_string` (String) The RADIUS server's secret
- `uuid` (String) uuid of the object

<a id="nestedblock--radius--instance_list--host"></a>
### Nested Schema for `radius.instance_list.host`

Optional:

- `hostip` (String) Server's hostname(Length 1-31) or IP address
- `hostipv6` (String) Server's IPV6 address


<a id="nestedblock--radius--instance_list--sampling_enable"></a>
### Nested Schema for `radius.instance_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'authen_success': Authentication Success; 'authen_failure': Authentication Failure; 'authorize_success': Authorization Success; 'authorize_failure': Authorization Failure; 'access_challenge': Access-Challenge Message Receive; 'timeout_error': Timeout; 'other_error': Other Error; 'request': Request; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;



<a id="nestedblock--radius--sampling_enable"></a>
### Nested Schema for `radius.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'authen_success': Total Authentication Success; 'authen_failure': Total Authentication Failure; 'authorize_success': Total Authorization Success; 'authorize_failure': Total Authorization Failure; 'access_challenge': Total Access-Challenge Message Receive; 'timeout_error': Total Timeout; 'other_error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;



<a id="nestedblock--windows"></a>
### Nested Schema for `windows`

Optional:

- `instance_list` (Block List) (see [below for nested schema](#nestedblock--windows--instance_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--windows--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--windows--instance_list"></a>
### Nested Schema for `windows.instance_list`

Required:

- `name` (String) Specify Windows authentication server name

Optional:

- `auth_protocol` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows--instance_list--auth_protocol))
- `health_check` (Number) Check server's health status
- `health_check_disable` (Number) Disable configured health check configuration
- `health_check_string` (String) Health monitor name
- `host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows--instance_list--host))
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `realm` (String) Specify realm of Windows server
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--windows--instance_list--sampling_enable))
- `support_apacheds_kdc` (Number) Enable weak cipher (DES CRC/MD5/MD4) and merge AS-REQ in single packet
- `timeout` (Number) Specify connection timeout to server, default is 10 seconds
- `uuid` (String) uuid of the object

<a id="nestedblock--windows--instance_list--auth_protocol"></a>
### Nested Schema for `windows.instance_list.auth_protocol`

Optional:

- `kdc_validate` (Number) Enable KDC validation
- `kerberos_disable` (Number) Disable Kerberos authentication protocol
- `kerberos_kdc_validation` (Block List, Max: 1) (see [below for nested schema](#nestedblock--windows--instance_list--auth_protocol--kerberos_kdc_validation))
- `kerberos_password_change_port` (Number) Specify the Kerbros password change port, default is 464
- `kerberos_port` (Number) Specify the Kerberos port, default is 88
- `kport_hm` (String) Check Kerberos port's health status
- `kport_hm_disable` (Number) Disable configured Kerberos port health check configuration
- `ntlm_disable` (Number) Disable NTLM authentication protocol
- `ntlm_health_check` (String) Check NTLM port's health status
- `ntlm_health_check_disable` (Number) Disable configured NTLM port health check configuration
- `ntlm_version` (Number) Specify NTLM version, default is 2

<a id="nestedblock--windows--instance_list--auth_protocol--kerberos_kdc_validation"></a>
### Nested Schema for `windows.instance_list.auth_protocol.kerberos_kdc_validation`

Optional:

- `kdc_account` (String) Specify account for KDC validation
- `kdc_password` (Number) Specify account password
- `kdc_pwd` (String) Account password
- `kdc_spn` (String) Specify SPN for KDC validation



<a id="nestedblock--windows--instance_list--host"></a>
### Nested Schema for `windows.instance_list.host`

Optional:

- `hostip` (String) Specify the Windows server's hostname(Length 1-31) or IP address
- `hostipv6` (String) Specify the Windows server's IPV6 address


<a id="nestedblock--windows--instance_list--sampling_enable"></a>
### Nested Schema for `windows.instance_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'krb_send_req_success': Kerberos Request; 'krb_get_resp_success': Kerberos Response; 'krb_timeout_error': Kerberos Timeout; 'krb_other_error': Kerberos Other Error; 'krb_pw_expiry': Kerberos password expiry; 'krb_pw_change_success': Kerberos password change success; 'krb_pw_change_failure': Kerberos password change failure; 'ntlm_proto_nego_success': NTLM Protocol Negotiation Success; 'ntlm_proto_nego_failure': NTLM Protocol Negotiation Failure; 'ntlm_session_setup_success': NTLM Session Setup Success; 'ntlm_session_setup_failure': NTLM Session Setup Failure; 'ntlm_prepare_req_success': NTLM Prepare Request Success; 'ntlm_prepare_req_error': NTLM Prepare Request Error; 'ntlm_auth_success': NTLM Authentication Success; 'ntlm_auth_failure': NTLM Authentication Failure; 'ntlm_timeout_error': NTLM Timeout; 'ntlm_other_error': NTLM Other Error; 'krb_validate_kdc_success': Kerberos KDC Validation Success; 'krb_validate_kdc_failure': Kerberos KDC Validation Failure;



<a id="nestedblock--windows--sampling_enable"></a>
### Nested Schema for `windows.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'kerberos-request-send': Total Kerberos Request; 'kerberos-response-get': Total Kerberos Response; 'kerberos-timeout-error': Total Kerberos Timeout; 'kerberos-other-error': Total Kerberos Other Error; 'ntlm-authentication-success': Total NTLM Authentication Success; 'ntlm-authentication-failure': Total NTLM Authentication Failure; 'ntlm-proto-negotiation-success': Total NTLM Protocol Negotiation Success; 'ntlm-proto-negotiation-failure': Total NTLM Protocol Negotiation Failure; 'ntlm-session-setup-success': Total NTLM Session Setup Success; 'ntlm-session-setup-failed': Total NTLM Session Setup Failure; 'kerberos-request-normal': Total Kerberos Normal Request; 'kerberos-request-dropped': Total Kerberos Dropped Request; 'kerberos-response-success': Total Kerberos Success Response; 'kerberos-response-failure': Total Kerberos Failure Response; 'kerberos-response-error': Total Kerberos Error Response; 'kerberos-response-timeout': Total Kerberos Timeout Response; 'kerberos-response-other': Total Kerberos Other Response; 'kerberos-job-start-error': Total Kerberos Job Start Error; 'kerberos-polling-control-error': Total Kerberos Polling Control Error; 'ntlm-prepare-req-success': Total NTLM Prepare Request Success; 'ntlm-prepare-req-failed': Total NTLM Prepare Request Failed; 'ntlm-timeout-error': Total NTLM Timeout; 'ntlm-other-error': Total NTLM Other Error; 'ntlm-request-normal': Total NTLM Normal Request; 'ntlm-request-dropped': Total NTLM Dropped Request; 'ntlm-response-success': Total NTLM Success Response; 'ntlm-response-failure': Total NTLM Failure Response; 'ntlm-response-error': Total NTLM Error Response; 'ntlm-response-timeout': Total NTLM Timeout Response; 'ntlm-response-other': Total NTLM Other Response; 'ntlm-job-start-error': Total NTLM Job Start Error; 'ntlm-polling-control-error': Total NTLM Polling Control Error; 'kerberos-pw-expiry': Total Kerberos password expiry; 'kerberos-pw-change-success': Total Kerberos password change success; 'kerberos-pw-change-failure': Total Kerberos password change failure; 'kerberos-validate-kdc-success': Total Kerberos KDC Validation Success; 'kerberos-validate-kdc-failure': Total Kerberos KDC Validation Failure; 'kerberos-generate-kdc-keytab-success': Total Kerberos KDC Keytab Generation Success; 'kerberos-generate-kdc-keytab-failure': Total Kerberos KDC Keytab Generation Failure; 'kerberos-delete-kdc-keytab-success': Total Kerberos KDC Keytab Deletion Success; 'kerberos-delete-kdc-keytab-failure': Total Kerberos KDC Keytab Deletion Failure; 'kerberos-kdc-keytab-count': Current Kerberos KDC Keytab Count;


