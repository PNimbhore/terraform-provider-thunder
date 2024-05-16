provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_sip" "thunder_ddos_template_sip" {
  sip_tmpl_name       = "test"
  idle_timeout        = 22
  ignore_zero_payload = 1
  action              = "drop"
  dst {
    sip_request_rate_limit {
      method {
        bye_cfg {
          dst_sip_bye_cfg_flag = 1
          dst_sip_bye_rate     = 3
        }
      }
    }
  }
  src {
    sip_request_rate_limit {
      method {
        invite_cfg {
          src_sip_invite_cfg_flag = 1
          src_sip_invite_rate     = 33
        }
      }
    }
  }
  user_tag = "test"
  malformed_sip {
    malformed_sip_check                   = "enable-check"
    malformed_sip_max_line_size           = 3
    malformed_sip_max_uri_length          = 4
    malformed_sip_max_header_name_length  = 2
    malformed_sip_max_header_value_length = 3
    malformed_sip_call_id_max_length      = 22
    malformed_sip_sdp_max_length          = 3
  }
  filter_header_list {
    sip_filter_header_seq       = 2
    sip_filter_header_regex     = "test"
    sip_filter_header_unmatched = 1
    sip_filter_header_blacklist = 1
    user_tag                    = "est"
  }
}