/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Addenda05 struct for Addenda05
type Addenda05 struct {
	// Client defined string used as a reference to this record.
	Id string `json:"id,omitempty"`
	// 05 - NACHA regulations
	TypeCode string `json:"typeCode,omitempty"`
	// Text for describing the related payment
	PaymentRelatedInformation string `json:"paymentRelatedInformation,omitempty"`
	// SequenceNumber is consecutively assigned to each Addenda05 Record following an Entry Detail Record. The first addenda05 sequence number must always be a 1.
	SequenceNumber float32 `json:"sequenceNumber,omitempty"`
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry Detail or Corporate Entry Detail Record's trace number This number is the same as the last seven digits of the trace number of the related Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber float32 `json:"entryDetailSequenceNumber,omitempty"`
}
