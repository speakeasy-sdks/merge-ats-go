// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"time"
)

// URLURLType - * `PERSONAL` - PERSONAL
// * `COMPANY` - COMPANY
// * `PORTFOLIO` - PORTFOLIO
// * `BLOG` - BLOG
// * `SOCIAL_MEDIA` - SOCIAL_MEDIA
// * `OTHER` - OTHER
// * `JOB_POSTING` - JOB_POSTING
type URLURLType string

const (
	URLURLTypePersonal    URLURLType = "PERSONAL"
	URLURLTypeCompany     URLURLType = "COMPANY"
	URLURLTypePortfolio   URLURLType = "PORTFOLIO"
	URLURLTypeBlog        URLURLType = "BLOG"
	URLURLTypeSocialMedia URLURLType = "SOCIAL_MEDIA"
	URLURLTypeOther       URLURLType = "OTHER"
	URLURLTypeJobPosting  URLURLType = "JOB_POSTING"
)

func (e URLURLType) ToPointer() *URLURLType {
	return &e
}

func (e *URLURLType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PERSONAL":
		fallthrough
	case "COMPANY":
		fallthrough
	case "PORTFOLIO":
		fallthrough
	case "BLOG":
		fallthrough
	case "SOCIAL_MEDIA":
		fallthrough
	case "OTHER":
		fallthrough
	case "JOB_POSTING":
		*e = URLURLType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for URLURLType: %v", v)
	}
}

// URL - # The Url Object
// ### Description
// The `Url` object is used to represent hyperlinks associated with the parent model.
// ### Usage Example
// Fetch from the `GET Candidate` endpoint and view their website urls.
type URL struct {
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The type of site.
	//
	// * `PERSONAL` - PERSONAL
	// * `COMPANY` - COMPANY
	// * `PORTFOLIO` - PORTFOLIO
	// * `BLOG` - BLOG
	// * `SOCIAL_MEDIA` - SOCIAL_MEDIA
	// * `OTHER` - OTHER
	// * `JOB_POSTING` - JOB_POSTING
	URLType *URLURLType `json:"url_type,omitempty"`
	// The site's url.
	Value *string `json:"value,omitempty"`
}

func (u URL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *URL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *URL) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *URL) GetURLType() *URLURLType {
	if o == nil {
		return nil
	}
	return o.URLType
}

func (o *URL) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
