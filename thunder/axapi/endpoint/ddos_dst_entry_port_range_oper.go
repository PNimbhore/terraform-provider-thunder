

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangeOper struct {
    
    IpFilteringPolicyOper DdosDstEntryPortRangeOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstEntryPortRangeOperOper `json:"oper"`
    PatternRecognition DdosDstEntryPortRangeOperPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryPortRangeOperPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
    PortInd DdosDstEntryPortRangeOperPortInd `json:"port-ind"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    ProgressionTracking DdosDstEntryPortRangeOperProgressionTracking `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    TopkSources DdosDstEntryPortRangeOperTopkSources `json:"topk-sources"`
    DstEntryName string 

}
type DataDdosDstEntryPortRangeOper struct {
    DtDdosDstEntryPortRangeOper DdosDstEntryPortRangeOper `json:"port-range"`
}


type DdosDstEntryPortRangeOperIpFilteringPolicyOper struct {
    Oper DdosDstEntryPortRangeOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryPortRangeOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryPortRangeOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryPortRangeOperOper struct {
    Ddos_entry_list []DdosDstEntryPortRangeOperOperDdos_entry_list `json:"ddos_entry_list"`
    ResourceLimitConfig string `json:"resource-limit-config"`
    ReourceLimitAlloc string `json:"reource-limit-alloc"`
    ResourceLimitRemain string `json:"resource-limit-remain"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    ResourceUsage int `json:"resource-usage"`
    L4ExtRate int `json:"l4-ext-rate"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryPortRangeOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryPortRangeOperPatternRecognition struct {
    Oper DdosDstEntryPortRangeOperPatternRecognitionOper `json:"oper"`
}


type DdosDstEntryPortRangeOperPatternRecognitionOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryPortRangeOperPatternRecognitionOperFilterList `json:"filter-list"`
}


type DdosDstEntryPortRangeOperPatternRecognitionOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryPortRangeOperPatternRecognitionPuDetails struct {
    Oper DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper `json:"oper"`
}


type DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOper struct {
    AllFilters []DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}


type DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstEntryPortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryPortRangeOperPortInd struct {
    Oper DdosDstEntryPortRangeOperPortIndOper `json:"oper"`
}


type DdosDstEntryPortRangeOperPortIndOper struct {
    Indicators []DdosDstEntryPortRangeOperPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryPortRangeOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}


type DdosDstEntryPortRangeOperProgressionTracking struct {
    Oper DdosDstEntryPortRangeOperProgressionTrackingOper `json:"oper"`
}


type DdosDstEntryPortRangeOperProgressionTrackingOper struct {
    Indicators []DdosDstEntryPortRangeOperProgressionTrackingOperIndicators `json:"indicators"`
}


type DdosDstEntryPortRangeOperProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstEntryPortRangeOperTopkSources struct {
    Oper DdosDstEntryPortRangeOperTopkSourcesOper `json:"oper"`
}


type DdosDstEntryPortRangeOperTopkSourcesOper struct {
    Indicators []DdosDstEntryPortRangeOperTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryPortRangeOperTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryPortRangeOperTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryPortRangeOperTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryPortRangeOperTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryPortRangeOperTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryPortRangeOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangeOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port-range/" +strconv.Itoa(p.PortRangeStart)+"+" +strconv.Itoa(p.PortRangeEnd)+"+"+p.Protocol+"/oper"
}

func (p *DdosDstEntryPortRangeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortRangeOper,error) {
logger.Println("DdosDstEntryPortRangeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortRangeOper
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
