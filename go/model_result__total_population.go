/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ResultTotalPopulation struct {

	// 総人口
	TotalPopulation float32 `json:"total_population,omitempty"`

	// 出生数
	BitrhNumber float32 `json:"bitrh_number,omitempty"`

	// 死亡数
	DiedNumber float32 `json:"died_number,omitempty"`
}
