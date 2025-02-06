package models

import "time"

type Router struct {
	DeviceID     uint32             `json:"device_id"`
	IPAddress    string             `json:"ip_address"`
	MACAddress   string             `json:"mac_address"`
	RoutingTable []Route            `json:"routing_table"`
	Interfaces   []NetworkInterface `json:"interfaces"`
	NATConfig    NAT                `json:"nat_config"`
	Status       string             `json:"status"`
}

type Route struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Metric      int    `json:"metric"`
}

type NetworkInterface struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	MAC    string `json:"mac"`
	Status string `json:"status"`
}

type NAT struct {
	Enabled bool      `json:"enabled"`
	Rules   []NATRule `json:"rules"`
}

type NATRule struct {
	SourceIP      string `json:"source_ip"`
	DestinationIP string `json:"destination_ip"`
	Protocol      string `json:"protocol"`
}

type Switch struct {
	DeviceID   string `json:"device_id"`
	MACAddress string `json:"mac_address"`
	Ports      []Port `json:"ports"`
	VLANs      []VLAN `json:"vlans"`
	Status     string `json:"status"`
}

type Port struct {
	Number     uint32 `json:"number"`
	MACAddress string `json:"mac_address"`
	Status     string `json:"status"`
}

type VLAN struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Ports []int  `json:"ports"`
}

type AccessPoint struct {
	DeviceID         uint32            `json:"device_id"`
	SSID             uint32            `json:"ssid"`
	Channel          int               `json:"channel"`
	ConnectedDevices []ConnectedDevice `json:"connected_devices"`
	SignalStrength   int               `json:"signal_strength"`
	Status           string            `json:"status"`
}

type ConnectedDevice struct {
	MACAddress     string `json:"mac_address"`
	IP             uint32 `json:"ip"`
	SignalStrength uint32 `json:"signal_strength"`
}

type Firewall struct {
	DeviceID uint32         `json:"device_id"`
	Rules    []FirewallRule `json:"rules"`
	Logs     []FirewallLog  `json:"logs"`
	Status   string         `json:"status"`
}

type FirewallRule struct {
	Action        string `json:"action"` // "allow" или "deny"
	SourceIP      string `json:"source_ip"`
	DestinationIP string `json:"destination_ip"`
	Protocol      string `json:"protocol"`
	Port          int    `json:"port"`
}

type FirewallLog struct {
	Timestamp     time.Time `json:"timestamp"`
	SourceIP      string    `json:"source_ip"`
	DestinationIP string    `json:"destination_ip"`
	Action        string    `json:"action"`
}

type OLT struct {
	DeviceID   uint32 `json:"device_id"`
	IPAddress  string `json:"ip_address"`
	MACAddress string `json:"mac_address"`
	ONUs       []ONU  `json:"onus"`
	VLANs      []VLAN `json:"vlans"`
	Status     string `json:"status"`
}

type ONU struct {
	DeviceID         uint32 `json:"device_id"`
	SerialNumber     string `json:"serial_number"`
	IPAddress        string `json:"ip_address"`
	MACAddress       string `json:"mac_address"`
	ConnectionStatus string `json:"connection_status"`
	ConnectionType   string `json:"connection_type"`  // "ethernet", "wifi"
	UpstreamSpeed    int    `json:"upstream_speed"`   // в Мбит/с
	DownstreamSpeed  int    `json:"downstream_speed"` // в Мбит/с
	Status           string `json:"status"`
}

type PONPort struct {
	PortNumber     uint32 `json:"port_number"`
	Status         string `json:"status"`
	ONUs           []ONU  `json:"onus"`
	SignalStrength int    `json:"signal_strength"` // в dBm
}

// ------------ Общая структура VLAN оборудования

type NetworkDevice struct {
	DeviceID   uint32   `json:"device_id"`
	DeviceType struct{} `json:"device_type"` // "router", "switch", "access_point", "firewall"
	Status     string   `json:"status"`
}

// ------- Общая структура GPon оборудования
type GPONDevice struct {
	DeviceID   uint32   `json:"device_id"`
	DeviceType struct{} `json:"device_type"` // "olt", "onu"
	Status     string   `json:"status"`
}
