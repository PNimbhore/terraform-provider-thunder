

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwRateLimitOper struct {
    
    Oper FwRateLimitOperOper `json:"oper"`
    Summary FwRateLimitOperSummary `json:"summary"`

}
type DataFwRateLimitOper struct {
    DtFwRateLimitOper FwRateLimitOper `json:"rate-limit"`
}


type FwRateLimitOperOper struct {
    RateLimitList []FwRateLimitOperOperRateLimitList `json:"rate-limit-list"`
    V4Address string `json:"v4-address"`
    V4Netmask string `json:"v4-netmask"`
    V6Prefix string `json:"v6-prefix"`
    TemplateId int `json:"template-id"`
}


type FwRateLimitOperOperRateLimitList struct {
    Address string `json:"address"`
    PrefixLen int `json:"prefix-len"`
    RuleName string `json:"rule-name"`
    TemplateId int `json:"template-id"`
    Type string `json:"type"`
    CpsReceived int `json:"cps-received"`
    CpsAllowed int `json:"cps-allowed"`
    UplinkTrafficReceived int `json:"uplink-traffic-received"`
    UplinkTrafficAllowed int `json:"uplink-traffic-allowed"`
    DownlinkTrafficReceived int `json:"downlink-traffic-received"`
    DownlinkTrafficAllowed int `json:"downlink-traffic-allowed"`
    TotalTrafficReceived int `json:"total-traffic-received"`
    TotalTrafficAllowed int `json:"total-traffic-allowed"`
    DropCount int `json:"drop-count"`
}


type FwRateLimitOperSummary struct {
    Oper FwRateLimitOperSummaryOper `json:"oper"`
}


type FwRateLimitOperSummaryOper struct {
    Mem_reserved int `json:"mem_reserved"`
    Mem_used int `json:"mem_used"`
    Alloc_failures int `json:"alloc_failures"`
    Total_num_entries int `json:"total_num_entries"`
    Total_entries_scope_aggregate int `json:"total_entries_scope_aggregate"`
    Total_entries_scope_subscriber_ip int `json:"total_entries_scope_subscriber_ip"`
    Total_entries_scope_subscriber_prefix int `json:"total_entries_scope_subscriber_prefix"`
    Total_entries_scope_parent int `json:"total_entries_scope_parent"`
    Total_entries_scope_parent_subscriber_ip int `json:"total_entries_scope_parent_subscriber_ip"`
    Total_entries_scope_parent_subscriberPrefix int `json:"total_entries_scope_parent_subscriber-prefix"`
}

func (p *FwRateLimitOper) GetId() string{
    return "1"
}

func (p *FwRateLimitOper) getPath() string{
    return "fw/rate-limit/oper"
}

func (p *FwRateLimitOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwRateLimitOper,error) {
logger.Println("FwRateLimitOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwRateLimitOper
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
