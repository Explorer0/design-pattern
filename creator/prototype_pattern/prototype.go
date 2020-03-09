package prototype_pattern

type Customer struct {
	name 	string
	address string
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetAddress() string {
	return c.address
}

func (c *Customer) SetName(name string) {
	c.name = name
}

func (c *Customer) SetAddress(address string) {
	c.address = address
}

func (c *Customer) Clone() *Customer{
	newCustomer := new(Customer)
	newCustomer.SetName(c.GetName())
	newCustomer.SetAddress(c.GetAddress())

	return newCustomer
}


type PrototypeManager struct {
	CustomerTemplateMap map[string]*Customer
}

func (pm *PrototypeManager) GetTemplateCustomer(templateName string) *Customer {
	if customer, ok := pm.CustomerTemplateMap[templateName]; ok {
		return customer.Clone()
	}
	return nil
}

func (pm *PrototypeManager) AddTemplateCustomer(templateName string, c *Customer)  {
	pm.CustomerTemplateMap[templateName] = c
}

func CreatePrototypeManager() *PrototypeManager {
	pm := new(PrototypeManager)

	fuckerCustomer := &Customer{"fucker", "China"}
	shitCustomer := &Customer{ "shiter", "American"}

	pm.CustomerTemplateMap["fucker"] = fuckerCustomer
	pm.CustomerTemplateMap["shiter"] = shitCustomer

	return pm
}

