package ie

type InformationElement interface {
	Serialize() []byte
}
