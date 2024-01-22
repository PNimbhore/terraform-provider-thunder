

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AdminAzureCred struct {
	Inst struct {

    Delete int `json:"delete"`
    FileUrl string `json:"file-url"`
    Import int `json:"import"`
    Show int `json:"show"`
    UseMgmtPort int `json:"use-mgmt-port"`
    User string 

	} `json:"azure-cred"`
}

func (p *AdminAzureCred) GetId() string{
    return "1"
}

func (p *AdminAzureCred) getPath() string{
    return "admin/" +p.Inst.User + "/azure-cred"
}

func (p *AdminAzureCred) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AdminAzureCred::Post")
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

func (p *AdminAzureCred) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AdminAzureCred::Get")
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
func (p *AdminAzureCred) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AdminAzureCred::Put")
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

func (p *AdminAzureCred) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AdminAzureCred::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
