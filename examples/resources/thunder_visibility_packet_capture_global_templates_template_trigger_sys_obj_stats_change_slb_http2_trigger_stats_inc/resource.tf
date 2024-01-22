provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2_trigger_stats_inc" {

  name                                    = "test"
  alloc_fail_total                        = 1
  bad_connection_preface                  = 1
  bad_frame_type_for_stream_state         = 1
  buff_alloc_error                        = 1
  cancel                                  = 1
  cant_allocate_control_frame             = 1
  cant_allocate_goaway_frame              = 1
  cant_allocate_ping_frame                = 1
  cant_allocate_rst_frame                 = 1
  cant_allocate_settings_frame            = 1
  cant_allocate_stream                    = 1
  cant_allocate_window_frame              = 1
  closed_state_unexpected_frame           = 1
  compression_error                       = 1
  connect_error                           = 1
  continuation_before_headers             = 1
  data_no_stream                          = 1
  data_queue_alloc_error                  = 1
  deflate_alloc_fail                      = 1
  enhance_your_calm                       = 1
  err_rcvd_total                          = 1
  err_sent_cancel                         = 1
  err_sent_compression_err                = 1
  err_sent_connect_err                    = 1
  err_sent_flow_control                   = 1
  err_sent_frame_size_err                 = 1
  err_sent_http11_required                = 1
  err_sent_inadequate_security            = 1
  err_sent_internal_err                   = 1
  err_sent_proto_err                      = 1
  err_sent_refused_stream                 = 1
  err_sent_setting_timeout                = 1
  err_sent_stream_closed                  = 1
  err_sent_total                          = 1
  err_sent_your_calm                      = 1
  error_max_invalid_stream                = 1
  exceeds_max_window_size_stream          = 1
  flow_control_error                      = 1
  frame_size_error                        = 1
  half_closed_remote_state_unexpected_fra = 1
  header_no_stream                        = 1
  header_padlen_gt_frame_payload          = 1
}