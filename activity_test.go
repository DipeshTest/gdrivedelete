package gdrivedelete

import (
	"io/ioutil"
	"testing"

	"fmt"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

/*
func TestGDriveDeleteFile_Success(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GluqBQm6pA0dRKdjmOymoKWp8tjbgnAEyqyWzDjfKlWupI9buZDpZJMW61tljW__Nr5-YqopzM-3salKzqBLRsGPgJSAcL7laT8bEzUqxTZes1Vf0YY4YO4SW21-")
	tc.SetInput("fileId", "1j4VoNLllOZWwwrsUQ6ncz4NACHYCX0FZ")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")

}


func TestGDriveDeleteFile_InvalidToken(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GluqBQm6pA0dRKdjmOymoKWp8tjbgnAEyqyWzDjfKlWupI9buZDpZJMW61tljW__Nr5-YqopzM-3salKzqBLRsGPgJSAcL7laT8bEzUqxTZes1Vf0YY4YO4SW21-")
	tc.SetInput("fileId", "1j4VoNLllOZWwwrsUQ6ncz4NACHYCX0FZ")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "401")


	tc.SetInput("accessToken", "")

	act.Eval(tc)

	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "105")
} */

func TestGDriveDeleteFile_InvalidFileId(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileId", "1j4VoNLllOZWwwrsUQ6ncz4NACHYCX0FZ")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "404")

	tc.SetInput("fileId", "")

	act.Eval(tc)

	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "106")
}
