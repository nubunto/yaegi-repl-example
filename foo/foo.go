package foo

type Capability struct {}

func NewCapability() *Capability {
  return &Capability{}
}

func (c *Capability) DoSomething() string {
  return "yes this is foo!"
}
