/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Nest2Joboffer struct {

	// 県コード
	PrefCode string `json:"prefCode,omitempty"`

	// 県名
	PrefName string `json:"prefName,omitempty"`

	// 年度
	Year string `json:"year,omitempty"`

	// 表示内容
	Matter string `json:"matter,omitempty"`

	// 表示分類
	Class string `json:"class,omitempty"`

	// 数値
	Data string `json:"data,omitempty"`
}
