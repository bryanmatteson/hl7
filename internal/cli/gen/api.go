package gen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HL7Version ...
type HL7Version string

// Versions ...
const (
	HL7v21  HL7Version = "2.1"
	HL7v22  HL7Version = "2.2"
	HL7v23  HL7Version = "2.3"
	HL7v231 HL7Version = "2.3.1"
	HL7v24  HL7Version = "2.4"
	HL7v25  HL7Version = "2.5"
	HL7v251 HL7Version = "2.5.1"
	HL7v26  HL7Version = "2.6"
	HL7v27  HL7Version = "2.7"
	HL7v271 HL7Version = "2.7.1"
	HL7v28  HL7Version = "2.8"
)

// CaristixClient ...
type CaristixClient struct {
	client  *http.Client
	baseURL string
}

const baseURL = "https://hl7-definition.caristix.com/v2-api/1/"

// NewCaristixClient ...
func NewCaristixClient(version HL7Version) *CaristixClient {
	return &CaristixClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: fmt.Sprintf("%sHL7v%s/", baseURL, version),
	}
}

func (c *CaristixClient) get(path string, data interface{}) error {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return err
	}

	r, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(data)
}

// GetDataTypesList ...
func (c *CaristixClient) GetDataTypesList() (items []APIItem, err error) {
	err = c.get("DataTypes", &items)
	return
}

// GetDataTypeDetail ...
func (c *CaristixClient) GetDataTypeDetail(id string) (item APIDataTypeItem, err error) {
	err = c.get("DataTypes/"+id, &item)
	return
}

// GetSegmentsList ...
func (c *CaristixClient) GetSegmentsList() (items []APIItem, err error) {
	err = c.get("Segments", &items)
	return
}

// GetSegmentDetail ...
func (c *CaristixClient) GetSegmentDetail(id string) (item APISegmentItem, err error) {
	err = c.get("Segments/"+id, &item)
	return
}

// GetTriggerEventsList ...
func (c *CaristixClient) GetTriggerEventsList() (items []APIItem, err error) {
	err = c.get("TriggerEvents", &items)
	return
}

// GetTriggerEventDetail ...
func (c *CaristixClient) GetTriggerEventDetail(id string) (item APIEventItem, err error) {
	err = c.get("TriggerEvents/"+id, &item)
	return
}

// GetTablesList ...
func (c *CaristixClient) GetTablesList() (items []APIItem, err error) {
	err = c.get("Tables", &items)
	return
}

// GetTableDetail ...
func (c *CaristixClient) GetTableDetail(id string) (item APITableItem, err error) {
	err = c.get("Tables/"+id, &item)
	return
}

// GetFieldDetail ...
func (c *CaristixClient) GetFieldDetail(id string) (item APIFieldItem, err error) {
	err = c.get("Fields/"+id, &item)
	return
}

// GetDataTypeFieldDetail ...
func (c *CaristixClient) GetDataTypeFieldDetail(id string) (item APIFieldItem, err error) {
	err = c.get("DataTypes/"+id, &item)
	return
}

// Kinds ...
const (
	DataTypeKind          = "DataType"
	DataTypeComponentKind = "DataTypeComponent"
	MessageKind           = "Message"
	SegmentKind           = "Segment"
	ComponentKind         = "Component"
	FieldKind             = "Field"
)

// APIItem ...
type APIItem struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Chapters    []string `json:"chapters"`
}

// APIDataTypeItem ...
type APIDataTypeItem struct {
	ID          string          `json:"id"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Fields      []*APIFieldInfo `json:"fields"`
}

// APIFieldInfo ...
type APIFieldInfo struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Position     string `json:"position"`
	Length       int    `json:"length"`
	DataType     string `json:"dataType"`
	DataTypeName string `json:"dataTypeName"`
	Usage        string `json:"usage"`
	Rpt          string `json:"rpt"`
	TableID      string `json:"tableId"`
	TableName    string `json:"tableName"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

// APIEventItem ...
type APIEventItem struct {
	ID          string            `json:"id"`
	MsgStructID string            `json:"msgStructId"`
	EventDesc   string            `json:"eventDesc"`
	Description string            `json:"description"`
	Sample      string            `json:"sample"`
	Chapters    []string          `json:"chapters"`
	Segments    []*APISegmentInfo `json:"segments"`
}

// APISegmentInfo ...
type APISegmentInfo struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	LongName string            `json:"longName"`
	Sequence string            `json:"sequence"`
	Usage    string            `json:"usage"`
	Rpt      string            `json:"rpt"`
	IsGroup  bool              `json:"isGroup"`
	Segments []*APISegmentInfo `json:"segments"`
}

// APIFieldItem ...
type APIFieldItem struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Position     string          `json:"position"`
	Name         string          `json:"name"`
	Length       int             `json:"length"`
	Usage        string          `json:"usage"`
	Rpt          string          `json:"rpt"`
	DataType     string          `json:"dataType"`
	DataTypeName string          `json:"dataTypeName"`
	TableID      string          `json:"tableId"`
	TableName    string          `json:"tableName"`
	Description  string          `json:"description"`
	Sample       string          `json:"sample"`
	Fields       []*APIFieldInfo `json:"fields"`
}

// APISegmentItem ...
type APISegmentItem struct {
	ID          string          `json:"id"`
	SegmentID   string          `json:"segmentId"`
	LongName    string          `json:"longName"`
	Description string          `json:"description"`
	Sample      string          `json:"sample"`
	Chapters    []string        `json:"chapters"`
	Fields      []*APIFieldInfo `json:"fields"`
}

// APITableItem ...
type APITableItem struct {
	ID        string           `json:"id"`
	TableID   string           `json:"tableId"`
	TableType string           `json:"tableType"`
	Name      string           `json:"name"`
	Chapters  []string         `json:"chapters"`
	Entries   []*APITableEntry `json:"entries"`
}

// APITableEntry ...
type APITableEntry struct {
	Value       string `json:"value"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
}
