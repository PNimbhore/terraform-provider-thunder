

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv4InIpv6FragStats struct {
    
    Stats Ipv4InIpv6FragStatsStats `json:"stats"`

}
type DataIpv4InIpv6FragStats struct {
    DtIpv4InIpv6FragStats Ipv4InIpv6FragStats `json:"frag"`
}


type Ipv4InIpv6FragStatsStats struct {
    SessionInserted int `json:"session-inserted"`
    SessionExpired int `json:"session-expired"`
    IcmpRcv int `json:"icmp-rcv"`
    Icmpv6Rcv int `json:"icmpv6-rcv"`
    UdpRcv int `json:"udp-rcv"`
    TcpRcv int `json:"tcp-rcv"`
    IpipRcv int `json:"ipip-rcv"`
    Ipv6ipRcv int `json:"ipv6ip-rcv"`
    OtherRcv int `json:"other-rcv"`
    IcmpDropped int `json:"icmp-dropped"`
    Icmpv6Dropped int `json:"icmpv6-dropped"`
    UdpDropped int `json:"udp-dropped"`
    TcpDropped int `json:"tcp-dropped"`
    IpipDropped int `json:"ipip-dropped"`
    Ipv6ipDropped int `json:"ipv6ip-dropped"`
    OtherDropped int `json:"other-dropped"`
    OverlapError int `json:"overlap-error"`
    BadIpLen int `json:"bad-ip-len"`
    TooSmall int `json:"too-small"`
    FirstTcpTooSmall int `json:"first-tcp-too-small"`
    FirstL4TooSmall int `json:"first-l4-too-small"`
    TotalSessionsExceeded int `json:"total-sessions-exceeded"`
    NoSessionMemory int `json:"no-session-memory"`
    FastAgingSet int `json:"fast-aging-set"`
    FastAgingUnset int `json:"fast-aging-unset"`
    FragmentQueueSuccess int `json:"fragment-queue-success"`
    UnalignedLen int `json:"unaligned-len"`
    ExceededLen int `json:"exceeded-len"`
    DuplicateFirstFrag int `json:"duplicate-first-frag"`
    DuplicateLastFrag int `json:"duplicate-last-frag"`
    TotalFragmentsExceeded int `json:"total-fragments-exceeded"`
    FragmentQueueFailure int `json:"fragment-queue-failure"`
    ReassemblySuccess int `json:"reassembly-success"`
    MaxLenExceeded int `json:"max-len-exceeded"`
    ReassemblyFailure int `json:"reassembly-failure"`
    PolicyDrop int `json:"policy-drop"`
    ErrorDrop int `json:"error-drop"`
    HighCpuThreshold int `json:"high-cpu-threshold"`
    LowCpuThreshold int `json:"low-cpu-threshold"`
    CpuThresholdDrop int `json:"cpu-threshold-drop"`
    IpdEntryDrop int `json:"ipd-entry-drop"`
    MaxPacketsExceeded int `json:"max-packets-exceeded"`
    SessionPacketsExceeded int `json:"session-packets-exceeded"`
    SctpRcv int `json:"sctp-rcv"`
    SctpDropped int `json:"sctp-dropped"`
    FirstGtpPacketTooSmall int `json:"first-gtp-packet-too-small"`
}

func (p *Ipv4InIpv6FragStats) GetId() string{
    return "1"
}

func (p *Ipv4InIpv6FragStats) getPath() string{
    return "ipv4-in-ipv6/frag/stats"
}

func (p *Ipv4InIpv6FragStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv4InIpv6FragStats,error) {
logger.Println("Ipv4InIpv6FragStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv4InIpv6FragStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
