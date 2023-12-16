// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"time"
)

// URLType - The type of site.
//
// * `PERSONAL` - PERSONAL
// * `COMPANY` - COMPANY
// * `PORTFOLIO` - PORTFOLIO
// * `BLOG` - BLOG
// * `SOCIAL_MEDIA` - SOCIAL_MEDIA
// * `OTHER` - OTHER
// * `JOB_POSTING` - JOB_POSTING
type URLType string

const (
	URLTypePersonal    URLType = "PERSONAL"
	URLTypeCompany     URLType = "COMPANY"
	URLTypePortfolio   URLType = "PORTFOLIO"
	URLTypeBlog        URLType = "BLOG"
	URLTypeSocialMedia URLType = "SOCIAL_MEDIA"
	URLTypeOther       URLType = "OTHER"
	URLTypeJobPosting  URLType = "JOB_POSTING"
)

func (e URLType) ToPointer() *URLType {
	return &e
}

func (e *URLType) UnmarshalJSON(data []byte) error {
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
		*e = URLType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for URLType: %v", v)
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
	URLType *URLType `json:"url_type,omitempty"`
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

func (o *URL) GetURLType() *URLType {
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
