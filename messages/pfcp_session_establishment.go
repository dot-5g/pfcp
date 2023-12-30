package messages

import "github.com/dot-5g/pfcp/ie"

type PFCPSessionEstablishmentRequest struct {
	NodeID  ie.NodeID // Mandatory
	CPFSEID ie.FSEID  // Mandatory
	// CreatePDR ie.CreatePDR // Mandatory
	// CreateFAR ie.CreateFAR // Mandatory
}
