

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortIpFilteringPolicyOperOper struct {
    
    Oper DdosDstEntryPortIpFilteringPolicyOperOperOper `json:"oper"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortIpFilteringPolicyOperOper struct {
    DtDdosDstEntryPortIpFilteringPolicyOperOper DdosDstEntryPortIpFilteringPolicyOperOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryPortIpFilteringPolicyOperOperOper struct {
    RuleList []DdosDstEntryPortIpFilteringPolicyOperOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryPortIpFilteringPolicyOperOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}

func (p *DdosDstEntryPortIpFilteringPolicyOperOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortIpFilteringPolicyOperOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/ip-filtering-policy-oper/oper"
}

func (p *DdosDstEntryPortIpFilteringPolicyOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortIpFilteringPolicyOperOper,error) {
logger.Println("DdosDstEntryPortIpFilteringPolicyOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortIpFilteringPolicyOperOper
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
