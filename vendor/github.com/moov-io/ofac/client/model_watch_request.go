/*
 * OFAC API
 *
 * OFAC (Office of Foreign Assets Control) API is designed to facilitate the enforcement of US government economic sanctions programs required by federal law. This project implements a modern REST HTTP API for companies and organizations to obey federal law and use OFAC data in their applications.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Webhook or other means of notification on search criteria. OFAC will make a POST request with a body of the customer (SDN, AltNames, and Address).
type WatchRequest struct {
	// Private token supplied by clients to be used for authenticating webhooks.
	AuthToken string `json:"authToken"`
	// HTTPS url for webhook on search match
	Webhook string `json:"webhook"`
}