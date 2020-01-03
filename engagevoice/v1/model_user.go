/*
 * RingCentral Engage Voice API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagevoice

type User struct {
	Enabled    bool   `json:"enabled,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	ParentPath string `json:"parentPath,omitempty"`
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName,omitempty"`
}
