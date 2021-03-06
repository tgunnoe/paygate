/*
 * Paygate Admin API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// Configs struct for Configs
type Configs struct {
	CutoffTimes         []CutoffTime         `json:"CutoffTimes,omitempty"`
	FileTransferConfigs []FileTransferConfig `json:"FileTransferConfigs,omitempty"`
	FTPConfigs          []FtpConfig          `json:"FTPConfigs,omitempty"`
	SFTPConfigs         []SftpConfig         `json:"SFTPConfigs,omitempty"`
}
