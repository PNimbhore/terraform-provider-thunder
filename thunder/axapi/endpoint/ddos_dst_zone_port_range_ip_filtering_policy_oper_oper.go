

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeIpFilteringPolicyOperOper struct {
    
    Oper DdosDstZonePortRangeIpFilteringPolicyOperOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangeIpFilteringPolicyOperOper struct {
    DtDdosDstZonePortRangeIpFilteringPolicyOperOper DdosDstZonePortRangeIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstZonePortRangeIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstZonePortRangeIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangeIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstZonePortRangeIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstZonePortRangeIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangeIpFilteringPolicyOperOper
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
