package models

import "fmt"

// Definir la interfaz correctamente
type Computer interface {
	InsertIntoUSBPort()
}

// Estructuras básicas
type Dell struct{}

func (d *Dell) InsertIntoUSBPort() {
	fmt.Println("USB connected into a Dell.")
}

type Mac struct{}

func (m *Mac) InsertIntoUSBPort() {
	fmt.Println("USB connected into a Mac.")
}

// Acción del cliente
type Client struct{}

func (c *Client) InsertUSBConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts USB connector into computer.")
	com.InsertIntoUSBPort()
}

// Adaptador para Mac
type MacAdapter struct {
	MacMachine *Mac
}

func (m *MacAdapter) InsertIntoUSBPort() {
	fmt.Println("Mac Adapter converts USB to USBC.")
	m.MacMachine.InsertIntoUSBPort()
}
