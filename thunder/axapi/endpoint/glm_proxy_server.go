

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GlmProxyServer struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    Host string `json:"host"`
    Password int `json:"password"`
    Port int `json:"port"`
    SecretString string `json:"secret-string"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`

	} `json:"proxy-server"`
}

func (p *GlmProxyServer) GetId() string{
    return "1"
}

func (p *GlmProxyServer) getPath() string{
    return "glm/proxy-server"
}

func (p *GlmProxyServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmProxyServer::Post")
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

func (p *GlmProxyServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmProxyServer::Get")
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
func (p *GlmProxyServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmProxyServer::Put")
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

func (p *GlmProxyServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmProxyServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
