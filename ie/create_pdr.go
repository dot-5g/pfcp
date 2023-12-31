package ie

type CreatePDR struct {
	IEType     uint16
	Length     uint16
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}
