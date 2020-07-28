package mui

//Prop - MUI Lang generic prop
type Prop struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//SetKey -
func (prop *Prop) SetKey(key string) {
	prop.Key = key
}

//GetKey -
func (prop *Prop) GetKey() string {
	return prop.Key
}

//SetValue -
func (prop *Prop) SetValue(value string) {
	prop.Value = value
}

//GetValue -
func (prop *Prop) GetValue() string {
	return prop.Value
}

//NewProp -
func NewProp(key, value string) *Prop {
	return &Prop{key, value}
}
