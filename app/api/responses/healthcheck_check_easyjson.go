// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package responses

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson469d3ca5DecodeOtusNotificationAppApiResponses(in *jlexer.Lexer, out *HealthcheckOkResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = Status(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson469d3ca5EncodeOtusNotificationAppApiResponses(out *jwriter.Writer, in HealthcheckOkResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HealthcheckOkResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson469d3ca5EncodeOtusNotificationAppApiResponses(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HealthcheckOkResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson469d3ca5EncodeOtusNotificationAppApiResponses(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HealthcheckOkResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson469d3ca5DecodeOtusNotificationAppApiResponses(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HealthcheckOkResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson469d3ca5DecodeOtusNotificationAppApiResponses(l, v)
}
