package mui

//Prop - MUI Lang generic prop
type Prop struct {
	key   string
	value string
}

//SetKey -
func (prop *Prop) SetKey(key string) {
	prop.key = key
}

//GetKey -
func (prop *Prop) GetKey() string {
	return prop.key
}

//SetValue -
func (prop *Prop) SetValue(value string) {
	prop.value = value
}

//GetValue -
func (prop *Prop) GetValue() string {
	return prop.value
}

//NewProp -
func NewProp(key, value string) *Prop {
	return &Prop{key, value}
}
