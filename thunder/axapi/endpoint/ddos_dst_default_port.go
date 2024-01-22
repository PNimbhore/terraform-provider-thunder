

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDefaultPort struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Template DdosDstDefaultPortTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"port"`
}


type DdosDstDefaultPortTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}

func (p *DdosDstDefaultPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)+"+"+p.Inst.Protocol
}

func (p *DdosDstDefaultPort) getPath() string{
    return "ddos/dst/default/" +p.Inst.DefaultAddressType + "/port"
}

func (p *DdosDstDefaultPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultPort::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstDefaultPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultPort::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *DdosDstDefaultPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultPort::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstDefaultPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDefaultPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
