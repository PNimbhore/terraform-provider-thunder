provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_global" "thunder_cgnv6_nat64_global" {
  all                    = 1
  force_non_zero_ipv4_id = 1
  icmp {
    send_on_port_unavailable    = "disable"
    send_on_user_quota_exceeded = "admin-filtered"
  }
  inside {
    source {
      class_list = "9"
    }
  }
  sampling_enable {
    counters1 = "all"
    counters2 = "h323_alg_no_quota"
    counters3 = "nat_pool_force_delete"
  }
  tcp {
    mss_clamp {
      mss_clamp_type = "subtract"
      mss_subtract   = 20
      min            = 476
    }
    reset_on_error {
      outbound = "disable"
    }
  }
  user_quota_prefix_length = 128
}