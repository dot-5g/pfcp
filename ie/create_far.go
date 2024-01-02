package ie

type CreateFAR struct {
	IEType      uint16
	Length      uint16
	FARID       FARID
	ApplyAction ApplyAction
}
