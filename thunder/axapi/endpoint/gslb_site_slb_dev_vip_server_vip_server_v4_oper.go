

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV4Oper struct {
    
    Ipv4 string `json:"ipv4"`
    Oper GslbSiteSlbDevVipServerVipServerV4OperOper `json:"oper"`
    SiteName string 
    DeviceName string 

}
type DataGslbSiteSlbDevVipServerVipServerV4Oper struct {
    DtGslbSiteSlbDevVipServerVipServerV4Oper GslbSiteSlbDevVipServerVipServerV4Oper `json:"vip-server-v4"`
}


type GslbSiteSlbDevVipServerVipServerV4OperOper struct {
    Dev_vip_addr string `json:"dev_vip_addr"`
    Dev_vip_state string `json:"dev_vip_state"`
    NodeName string `json:"node-name"`
    ServiceIp string `json:"service-ip"`
    PortCount int `json:"port-count"`
    VirtualServer int `json:"virtual-server"`
    Disabled int `json:"disabled"`
    GslbProtocol int `json:"gslb-protocol"`
    LocalProtocol int `json:"local-protocol"`
    ManuallyHealthCheck int `json:"manually-health-check"`
    Use_gslb_state int `json:"use_gslb_state"`
    Dynamic int `json:"dynamic"`
    Hits int `json:"hits"`
    Recent int `json:"recent"`
    DevVipPortList []GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList `json:"dev-vip-port-list"`
}


type GslbSiteSlbDevVipServerVipServerV4OperOperDevVipPortList struct {
    DevVipPortNum int `json:"dev-vip-port-num"`
    DevVipPortState string `json:"dev-vip-port-state"`
}

func (p *GslbSiteSlbDevVipServerVipServerV4Oper) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevVipServerVipServerV4Oper) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/" +p.DeviceName + "/vip-server/vip-server-v4/"+p.Ipv4+"/oper"
}

func (p *GslbSiteSlbDevVipServerVipServerV4Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevVipServerVipServerV4Oper,error) {
logger.Println("GslbSiteSlbDevVipServerVipServerV4Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevVipServerVipServerV4Oper
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
