// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IgnoreCommonModelRequest2 struct {
	Message *string     `form:"name=message" multipartForm:"name=message"`
	Reason  interface{} `form:"name=reason" multipartForm:"name=reason"`
}

func (o *IgnoreCommonModelRequest2) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *IgnoreCommonModelRequest2) GetReason() interface{} {
	if o == nil {
		return nil
	}
	return o.Reason
}
