/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ResultJoboffer struct {

	// 職業大分類1
	BroadName string `json:"broadName,omitempty"`

	// 有効求人数
	ValidJobOfferNumber float32 `json:"valid_job_offer_number,omitempty"`

	// 有効就職者
	ValidJobfinder float32 `json:"valid_jobfinder,omitempty"`

	// 就職件数
	FindingEmploymentCount float32 `json:"finding_employment_count,omitempty"`
}