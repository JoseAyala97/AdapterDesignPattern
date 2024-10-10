package main

import "AdapterDesignPatter/internal/models"

func main() {
	client := &models.Client{}
	dell := &models.Dell{}

	// Cliente inserta el conector USB en un Dell
	client.InsertUSBConnectorIntoComputer(dell)

	// Cliente inserta el conector USB en un Mac utilizando el adaptador
	macMachine := &models.Mac{}
	macMachineAdapter := &models.MacAdapter{
		MacMachine: macMachine,
	}
	client.InsertUSBConnectorIntoComputer(macMachineAdapter)
}
