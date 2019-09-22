package main

import (
	"github.com/adonese/noebs/ebs_fields"
	"reflect"
	"testing"
)

//func Test_handleAdditionalFields(t *testing.T) {
//	type args struct {
//		fields *ebs_fields.GenericEBSResponseFields
//	}
//	tests := []struct {
//		name string
//		args args
//		want map[string]interface{}
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := handleAdditionalFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("handleAdditionalFields() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_additionalFieldsToHash(t *testing.T) {

	a := `SalesAmount=10.3;FixedFee=22.3;Token=23232;MeterNumber=12345;CustomerName=mohamed`
	want := map[string]interface{}{
		"SalesAmount": "10.3", "FixedFee": "22.3", "Token": "23232", "MeterNumber": "12345", "CustomerName": "mohamed",
	}
	tests := []struct {
		name string
		args string
		want map[string]interface{}
	}{
		{"successful case - nec", a, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := additionalFieldsToHash(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("additionalFieldsToHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_handleChan(t *testing.T) {
//	//var ch = make(chan *ebs_fields.GenericEBSResponseFields)
//	f := generateFields()
//	billChan <- *f
//	want := wantFields()
//	tests := []struct {
//		name  string
//		want  necBill
//		have necBill
//	}{
//		{"nec successful stuff", want, want},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, ok := handleChan()
//			if !ok {
//				t.Errorf("there's an error")
//			}
//			g := got.(necBill)
//			if ok := reflect.DeepEqual(g, tt.want); !ok{
//				t.Errorf("there is no workign here: have: :%v, %v", g, tt.want)
//			}
//		})
//	}
//}

func generateFields() *ebs_fields.GenericEBSResponseFields {
	var f ebs_fields.GenericEBSResponseFields
	f.AdditionalData = "SalesAmount=10.3;FixedFee=22.3;Token=23232;MeterNumber=12345;CustomerName=mohamed"
	return &f
}

func wantFields() necBill {
	v := necBill{
		SalesAmount:  10.3,
		FixedFee:     22.3,
		Token:        "23232",
		MeterNumber:  "12345",
		CustomerName: "mohamed",
	}
	return v
}