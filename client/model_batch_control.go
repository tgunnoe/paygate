/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// BatchControl struct for BatchControl
type BatchControl struct {
	// Batch ID
	ID string `json:"ID,omitempty"`
	// Same as ServiceClassCode in BatchHeaderRecord
	ServiceClassCode int32 `json:"serviceClassCode,omitempty"`
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda Record processed, within either the batch or file as appropriate.
	EntryAddendaCount int32 `json:"entryAddendaCount,omitempty"`
	// Validate the Receiving DFI Identification in each Entry Detail Record is hashed to provide a check against inadvertent alteration of data contents due to hardware failure or program error. In this context the Entry Hash is the sum of the corresponding fields in the Entry Detail Records on the file.
	EntryHash int32 `json:"entryHash,omitempty"`
	// Contains accumulated Entry debit totals within the batch.
	TotalDebit int32 `json:"totalDebit,omitempty"`
	// Contains accumulated Entry credit totals within the batch.
	TotalCredit int32 `json:"totalCredit,omitempty"`
	// Alphanumeric code used to identify an Originator The Company Identification Field must be included on all prenotification records and on each entry initiated pursuant to such prenotification. The Company ID may begin with the ANSI one-digit Identification Code Designator (ICD), followed by the identification number The ANSI Identification Numbers and related Identification Code IRS Employer Identification Number (EIN) \"1\" Data Universal Numbering Systems (DUNS) \"3\" User Assigned Number \"9\"
	CompanyIdentification string `json:"companyIdentification,omitempty"`
	// MAC is an eight character code derived from a special key used in conjunction with the DES algorithm. The purpose of the MAC is to validate the authenticity of ACH entries. The DES algorithm and key message standards must be in accordance with standards adopted by the American National Standards Institute. The remaining eleven characters of this field are blank.
	MessageAuthentication string `json:"messageAuthentication,omitempty"`
	// The routing number is used to identify the DFI originating entries within a given branch.
	ODFIIdentification string `json:"ODFIIdentification,omitempty"`
	// BatchNumber is assigned in ascending sequence to each batch by the ODFI or its Sending Point in a given file of entries. Since the batch number in the Batch Header Record and the Batch Control Record is the same, the ascending sequence number should be assigned by batch and not by record.
	BatchNumber string `json:"batchNumber,omitempty"`
}
