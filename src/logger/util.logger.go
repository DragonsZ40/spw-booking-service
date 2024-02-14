package logger

import (
	"encoding/json"
	"flag"
	"strings"
	"time"

	utils "spw-booking-service/src/utils"
)

func MapStruct(LogModel *LogModel) map[string]interface{} {
	s, _ := json.Marshal(LogModel)
	j := make(map[string]*json.RawMessage)
	_ = json.Unmarshal(s, &j)
	jj := make(map[string]*json.RawMessage, len(j))
	for k, v := range j {
		jj[strings.ToLower(k)] = v
	}
	ss, _ := json.Marshal(jj)

	y := make(map[string]interface{})
	_ = json.Unmarshal(ss, &y)
	return y
}

func GetLogModel(req LogModelRequest) LogModel {
	var Log LogModel
	Log.Suffix = req.Suffix
	Log.ProjectName = req.Name
	Log.Team = req.Team
	Log.Product = req.Product
	Log.ServiceType = req.ServiceType
	Log.Method = req.Method
	Log.StepName = req.StepName
	Log.StartDate = utils.ConvDatetimeFormatLog(req.Start)
	Log.Requestor = utils.GetValueStr(&req.Channel)
	Log.Txid = req.TrackingID
	Log.Channel = utils.GetValueStr(&req.Channel)
	if req.DealerCode != "" {
		Log.DealerCode = utils.GetValueStr(&req.DealerCode)
	}
	if req.RefId != "" {
		Log.Refid = req.RefId
	} else if req.RefId == "" {
		Log.Refid = utils.GetValueStr(&req.CorrelationID)
	}
	if req.SearchKey != "" {
		Log.SearchKey = req.SearchKey
	} else if req.SearchKey == "" {
		Log.SearchKey = utils.GetValueStr(&req.CorrelationID)
	}
	if req.SubscriberNumber != "" {
		Log.Msisdn = utils.GetValueStr(&req.SubscriberNumber)
	}
	return Log
}

func LogOrder(req LogOrderRequest, l LogModel, start time.Time) {
	l.Request = req.Request
	l.Response = req.Response
	l.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	l.ElapsedTime = float64(time.Since(start).Milliseconds())
	l.ResultCode = req.ResultCode
	l.ResultDesc = req.ResultDesc
	l.CertificateId = req.CertificateId
	l.CertificateType = req.CertificateType
	if flag.Lookup("test.v") == nil {
		if req.ResultCode == "200" {
			l.LogOrderSuccess()
		} else {
			l.LogOrderFail()
		}
	}
}

func LogStep(req LogStepRequest, l LogModel, start time.Time) {
	l.StepName = req.StepName
	l.StartDate = utils.ConvDatetimeFormatLog(start)
	l.StepRequest = req.StepRequest
	l.StepResponse = req.StepResponse
	l.Endpoint = req.Endpoint
	l.ResultCode = req.ResultCode
	l.ResultDesc = req.ResultDesc
	l.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	l.ElapsedTime = float64(time.Since(start).Milliseconds())
	if flag.Lookup("test.v") == nil {
		if l.ResultCode == "" || l.ResultCode == "200" {
			l.ResultCode = "200"
			l.LogStepSuccess()
		} else {
			l.LogStepFail()
		}
	}
}
