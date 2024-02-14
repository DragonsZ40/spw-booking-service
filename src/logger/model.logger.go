package logger

import "time"

type LogModel struct {
	CertificateId        string  `json:"certificate_id"`
	CertificateType      string  `json:"certificate_type"`
	Channel              string  `json:"channel"` // ussd, web ...
	DealerCode           string  `json:"dealer_code"`
	EmployeeId           string  `json:"employee_id"`
	EndDate              string  `json:"end_date"`
	Endpoint             string  `json:"endpoint"`
	ElapsedTime          float64 `json:"elapsed_time"`
	ExtraInfo            string  `json:"extra_info"`
	Header               string  `json:"header"` // request header
	Info                 string  `json:"info"`
	KbUrl                string  `json:"kb_url"`
	LogCat               string  `json:"log_cat"` // order, detail
	Msisdn               string  `json:"msisdn"`  // ให้เก็บเป็น format ที่มีเลข 0 นำหน้าเท่านั้น ไม่ใช้ format 66
	Method               string  `json:"method"`
	Product              string  `json:"product"`
	ProductInfo          string  `json:"product_info"`
	ProjectName          string  `json:"project_name"`
	Remark               string  `json:"remark"`
	Requestor            string  `json:"requestor"`
	Request              string  `json:"request"`
	Response             string  `json:"response"`
	Refid                string  `json:"ref_id"`           // มันคือ `transaction_id` ที่คนอื่นยิงเข้ามาหา
	ResultIndicator      string  `json:"result_indicator"` // FAILED, COMPLETED, ...
	ResultCode           string  `json:"result_code"`
	ResultDesc           string  `json:"result_desc"`
	SearchKey            string  `json:"search_key"`
	ServiceType          string  `json:"service_type"`
	StackTrace           string  `json:"stack_trace"`
	StartDate            string  `json:"start_date"`
	StepName             string  `json:"step_name"`
	StepRequest          string  `json:"step_request"`
	StepResponse         string  `json:"step_response"`
	StepTxid             string  `json:"step_txid"` // APP Step Tracking id
	System               string  `json:"system"`
	Suffix               string  `json:"@suffix"`
	SubChannel           string  `json:"sub_channel"`
	Team                 string  `json:"@team"`
	Txid                 string  `json:"txid"`                   // APP Tracking id
	TransactionExtraInfo string  `json:"transaction_extra_info"` // อิสระเลยจะใช้กี่ Object ใน Array นี้ก็ได้ แต่ขอแค่เป็น name, value => // example: [{name: 'code', value: '200'}, {name: 'desc', value: 'Success'}]
	TransactionData      string  `json:"transaction_data"`
	AppVersion           string  `json:"app_version"`       // ค่าจาก Header หน้าบ้าน
	DeviceBrand          string  `json:"device_brand"`      // ค่าจาก Header หน้าบ้าน
	DeviceModel          string  `json:"device_model"`      // ค่าจาก Header หน้าบ้าน
	DeviceOS             string  `json:"device_os"`         // ค่าจาก Header หน้าบ้าน
	DeviceOSVersion      string  `json:"device_os_sersion"` // ค่าจาก Header หน้าบ้าน
	ShopCode             string  `json:"shop_code"`         // ค่าจาก Header หน้าบ้าน
	ShopName             string  `json:"shop_name"`         // ค่าจาก Header หน้าบ้าน
	DeviceType           string  `json:"device_type"`       // ค่าจาก Header หน้าบ้าน
}

type LogModelRequest struct {
	ServiceType      string
	Channel          string
	CorrelationID    string
	DealerCode       string
	SubscriberNumber string
	Method           string
	StepName         string
	Start            time.Time
	TrackingID       string
	Suffix           string
	Name             string
	Team             string
	Product          string
	RefId            string
	SearchKey        string
}

type LogOrderRequest struct {
	Request         string
	Response        string
	ResultCode      string
	ResultDesc      string
	CertificateId   string
	CertificateType string
}

type LogStepRequest struct {
	StepName     string
	StartDate    string
	StepRequest  string
	StepResponse string
	Endpoint     string
	ResultCode   string
	ResultDesc   string
	Method       string
	System       string
	Response     string
}
