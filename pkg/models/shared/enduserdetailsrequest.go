// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EndUserDetailsRequest struct {
	// The integration categories to show in Merge Link.
	Categories []CategoriesEnum `json:"categories" form:"name=categories" multipartForm:"name=categories"`
	// An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
	CommonModels []CommonModelScopesBodyRequest `json:"common_models,omitempty" form:"name=common_models,json" multipartForm:"name=common_models,json"`
	// Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
	EndUserEmailAddress string `json:"end_user_email_address" form:"name=end_user_email_address" multipartForm:"name=end_user_email_address"`
	// Your end user's organization.
	EndUserOrganizationName string `json:"end_user_organization_name" form:"name=end_user_organization_name" multipartForm:"name=end_user_organization_name"`
	// This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
	EndUserOriginID string `json:"end_user_origin_id" form:"name=end_user_origin_id" multipartForm:"name=end_user_origin_id"`
	// The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://www.merge.dev/docs/basics/integration-metadata/.
	Integration *string `json:"integration,omitempty" form:"name=integration" multipartForm:"name=integration"`
	// An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
	LinkExpiryMins *int64 `json:"link_expiry_mins,omitempty" form:"name=link_expiry_mins" multipartForm:"name=link_expiry_mins"`
	// Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
	ShouldCreateMagicLinkURL *bool `json:"should_create_magic_link_url,omitempty" form:"name=should_create_magic_link_url" multipartForm:"name=should_create_magic_link_url"`
}

func (o *EndUserDetailsRequest) GetCategories() []CategoriesEnum {
	if o == nil {
		return []CategoriesEnum{}
	}
	return o.Categories
}

func (o *EndUserDetailsRequest) GetCommonModels() []CommonModelScopesBodyRequest {
	if o == nil {
		return nil
	}
	return o.CommonModels
}

func (o *EndUserDetailsRequest) GetEndUserEmailAddress() string {
	if o == nil {
		return ""
	}
	return o.EndUserEmailAddress
}

func (o *EndUserDetailsRequest) GetEndUserOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.EndUserOrganizationName
}

func (o *EndUserDetailsRequest) GetEndUserOriginID() string {
	if o == nil {
		return ""
	}
	return o.EndUserOriginID
}

func (o *EndUserDetailsRequest) GetIntegration() *string {
	if o == nil {
		return nil
	}
	return o.Integration
}

func (o *EndUserDetailsRequest) GetLinkExpiryMins() *int64 {
	if o == nil {
		return nil
	}
	return o.LinkExpiryMins
}

func (o *EndUserDetailsRequest) GetShouldCreateMagicLinkURL() *bool {
	if o == nil {
		return nil
	}
	return o.ShouldCreateMagicLinkURL
}