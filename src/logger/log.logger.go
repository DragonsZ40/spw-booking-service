package logger

import (
	"fmt"
	"os"
	"time"

	utils "spw-booking-service/src/utils"

	logs "github.com/sirupsen/logrus"
)

func init() {
	logs.SetFormatter(&logs.JSONFormatter{})
	logs.SetOutput(os.Stdout)
}

func (Log *LogModel) LogOrderSuccess() {
	fmt.Println("--- LogOrderSuccess ---")
	Log.StepTxid = Log.Txid
	Log.LogCat = "order"
	Log.ResultIndicator = "COMPLETED"
	Log.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	m := MapStruct(Log)
	logs.WithFields(m).Info(Log.ResultIndicator)
}

func (Log *LogModel) LogOrderFail() {
	fmt.Println("--- LogOrderFail ---")
	Log.StepTxid = Log.Txid
	Log.LogCat = "order"
	Log.ResultIndicator = "FAILED"
	Log.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	m := MapStruct(Log)
	logs.WithFields(m).Info(Log.ResultIndicator)
}

func (Log *LogModel) LogStepSuccess() {
	fmt.Println("--- LogStepSuccess ---")
	Log.StepTxid = utils.GetUUID()
	Log.LogCat = "detail"
	Log.ResultIndicator = "SUCCESS"
	Log.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	m := MapStruct(Log)
	logs.WithFields(m).Info(Log.ResultIndicator)
}

func (Log *LogModel) LogStepFail() {
	fmt.Println("--- LogStepFail ---")
	Log.StepTxid = utils.GetUUID()
	Log.LogCat = "detail"
	Log.ResultIndicator = "FAILED"
	Log.EndDate = utils.ConvDatetimeFormatLog(time.Now())
	m := MapStruct(Log)
	logs.WithFields(m).Info(Log.ResultIndicator)
}

func LogInfo(s string) {
	fmt.Println(s)
}
