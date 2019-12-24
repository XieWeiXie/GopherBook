package internal

type Version struct {
	Value string
}

func NewVersion(v string) *Version {
	return &Version{Value: v}
}

func (V *Version) SetValue(v string) {
	V.Value = v
}

func (V *Version) GetValue() string {
	return V.Value
}
