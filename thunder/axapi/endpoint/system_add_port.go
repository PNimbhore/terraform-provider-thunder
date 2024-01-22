

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemAddPort struct {
	Inst struct {

    PortIndex int `json:"port-index"`

	} `json:"add-port"`
}

func (p *SystemAddPort) GetId() string{
    return "1"
}

func (p *SystemAddPort) getPath() string{
    return "system/add-port"
}

func (p *SystemAddPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAddPort::Post")
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

func (p *SystemAddPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAddPort::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *SystemAddPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAddPort::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SystemAddPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAddPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
