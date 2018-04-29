package gdrivedelete

import (
	s "strings"

	"github.com/DipeshTest/allstarsshared/GDrive"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-gdrivedelete")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	fileId := s.TrimSpace(context.GetInput("fileId").(string))
	timeout := s.TrimSpace(context.GetInput("timeout").(string))

	if len(s.TrimSpace(accessToken)) == 0 {
		// fmt.Println(Twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "Access Token field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(s.TrimSpace(fileId)) == 0 {
		// fmt.Println(Twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "106")

		context.SetOutput("message", "File ID field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else {
		if len(timeout) == 0 {
			timeout = "120"
		}
		code, msg := GDrive.DeleteFile(fileId, accessToken, timeout)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
