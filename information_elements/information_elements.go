package information_elements

type InformationElement interface {
	Serialize() []byte
}
