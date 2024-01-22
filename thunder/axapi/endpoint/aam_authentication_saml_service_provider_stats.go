

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlServiceProviderStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationSamlServiceProviderStatsStats `json:"stats"`

}
type DataAamAuthenticationSamlServiceProviderStats struct {
    DtAamAuthenticationSamlServiceProviderStats AamAuthenticationSamlServiceProviderStats `json:"service-provider"`
}


type AamAuthenticationSamlServiceProviderStatsStats struct {
    SpMetadataExportReq int `json:"sp-metadata-export-req"`
    SpMetadataExportSuccess int `json:"sp-metadata-export-success"`
    LoginAuthReq int `json:"login-auth-req"`
    LoginAuthResp int `json:"login-auth-resp"`
    AcsReq int `json:"acs-req"`
    AcsSuccess int `json:"acs-success"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    SloReq int `json:"slo-req"`
    SloSuccess int `json:"slo-success"`
    SloError int `json:"slo-error"`
    SpSloReq int `json:"sp-slo-req"`
    GloSloSuccess int `json:"glo-slo-success"`
    LocSloSuccess int `json:"loc-slo-success"`
    ParSloSuccess int `json:"par-slo-success"`
    OtherError int `json:"other-error"`
}

func (p *AamAuthenticationSamlServiceProviderStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationSamlServiceProviderStats) getPath() string{
    
    return "aam/authentication/saml/service-provider/"+p.Name+"/stats"
}

func (p *AamAuthenticationSamlServiceProviderStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSamlServiceProviderStats,error) {
logger.Println("AamAuthenticationSamlServiceProviderStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSamlServiceProviderStats
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
