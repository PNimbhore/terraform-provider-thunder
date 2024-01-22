

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileLogBackup struct {
	Inst struct {

    All int `json:"all"`
    Date int `json:"date"`
    Day int `json:"day"`
    Expedite int `json:"expedite"`
    FileHandle string `json:"file-handle"`
    Month int `json:"month"`
    Period int `json:"period"`
    StatsData int `json:"stats-data"`
    Week int `json:"week"`
    FileContent []byte `json:"-"`
    File string `json:"file"`

	} `json:"log-backup"`
}
type DeleteFileLogBackup struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"log-backup"`
}

func (p *FileLogBackup) GetId() string{
    return "1"
}

func (p *FileLogBackup) getPath() string{
    return "file/log-backup"
}

func (p *FileLogBackup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileLogBackup::Post")
    headers := axapi.GenRequestHeader(authToken)
        f, error := os.Open(p.Inst.FileHandle)
        if error != nil {
            logger.Println("Failed to open a file: ", error)
            return error
        }
        data, error := ioutil.ReadAll(f)
        defer f.Close()
        if error != nil {
            logger.Println("Failed to read file: ", error)
            return error
        }
        s := &FileLogBackup{}
            s.Inst.All = p.Inst.All
            s.Inst.Date = p.Inst.Date
            s.Inst.Day = p.Inst.Day
            s.Inst.Expedite = p.Inst.Expedite
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.Month = p.Inst.Month
            s.Inst.Period = p.Inst.Period
            s.Inst.StatsData = p.Inst.StatsData
            s.Inst.Week = p.Inst.Week
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileLogBackup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileLogBackup::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/log-backup/oper", "", nil, headers, logger)
    if err == nil {
        if len(axResp) != 0 {
        err = json.Unmarshal(axResp, &p)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
                     }
            }
   }
    return err
}
func (p *FileLogBackup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileLogBackup::Put")
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

func (p *FileLogBackup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileLogBackup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileLogBackup{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/log-backup", payloadBytes, headers, logger)
    return err
}
