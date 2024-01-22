

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementNtpOper struct {
    
    Oper EnableManagementNtpOperOper `json:"oper"`

}
type DataEnableManagementNtpOper struct {
    DtEnableManagementNtpOper EnableManagementNtpOper `json:"ntp"`
}


type EnableManagementNtpOperOper struct {
    PortList []EnableManagementNtpOperOperPortList `json:"port-list"`
}


type EnableManagementNtpOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementNtpOper) GetId() string{
    return "1"
}

func (p *EnableManagementNtpOper) getPath() string{
    return "enable-management/ntp/oper"
}

func (p *EnableManagementNtpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementNtpOper,error) {
logger.Println("EnableManagementNtpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementNtpOper
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
