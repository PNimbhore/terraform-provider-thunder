---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_ftp_proxy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_ftp_proxy: Configure FTP Proxy global
  PLACEHOLDER
---

# thunder_slb_ftp_proxy (Resource)

`thunder_slb_ftp_proxy`: Configure FTP Proxy global

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_ftp_proxy" "test_thunder_slb_ftp_proxy" {
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

- `counters1` (String) 'all': all; 'num': Num; 'curr': Current proxy conns; 'total': Total proxy conns; 'svrsel_fail': Server selection failure; 'no_route': no route failure; 'snat_fail': source nat failure; 'feat': feat packet; 'cc': clear ctrl port packet; 'data_ssl': data ssl force; 'line_too_long': line too long; 'line_mem_freed': request line freed; 'invalid_start_line': invalid start line; 'auth_tls': auth tls cmd; 'prot': prot cmd; 'pbsz': pbsz cmd; 'pasv': pasv cmd; 'port': port cmd; 'request_dont_care': other cmd; 'client_auth_tls': client auth tls; 'cant_find_pasv': cant find pasv; 'pasv_addr_ne_server': psv addr not equal to svr; 'smp_create_fail': smp create fail; 'data_server_conn_fail': data svr conn fail; 'data_send_fail': data send fail; 'epsv': epsv command; 'cant_find_epsv': cant find epsv; 'data_curr': Current Data Proxy; 'data_total': Total Data Proxy; 'auth_unsupported': Unsupported auth; 'adat': adat cmd; 'unsupported_pbsz_value': Unsupported PBSZ; 'unsupported_prot_value': Unsupported PROT; 'unsupported_command': Unsupported cmd; 'control_to_clear': Control chn clear txt; 'control_to_ssl': Control chn ssl; 'bad_sequence': Bad Sequence; 'rsv_persist_conn_fail': Serv Sel Persist fail; 'smp_v6_fail': Serv Sel SMPv6 fail; 'smp_v4_fail': Serv Sel SMPv4 fail; 'insert_tuple_fail': Serv Sel insert tuple fail; 'cl_est_err': Client EST state erro; 'ser_connecting_err': Serv CTNG state error; 'server_response_err': Serv RESP state error; 'cl_request_err': Client RQ state error; 'data_conn_start_err': Data Start state error; 'data_serv_connecting_err': Data Serv CTNG error; 'data_serv_connected_err': Data Serv CTED error; 'request': Total FTP Request; 'auth_req': Auth Request; 'auth_succ': Auth Success; 'auth_fail': Auth Failure; 'fwd_to_internet': Forward to Internet; 'fwd_to_sg': Total Forward to Service-group; 'drop': Total FTP Drop; 'ds_succ': Host Domain Name is resolved; 'ds_fail': Host Domain Name isn't resolved; 'open': open cmd; 'site': site cmd; 'user': user cmd; 'pass': pass cmd; 'quit': quit cmd; 'eprt': eprt cmd; 'cant_find_port': cant find port; 'cant_find_eprt': cant find eprt;


