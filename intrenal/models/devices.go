package models

import "time"

type Router struct {
	DeviceID     uint32             `json:"device_id" gorm:"primaryKey"`
	IPAddress    string             `json:"ip_address"`
	MACAddress   string             `json:"mac_address"`
	RoutingTable []Route            `json:"routing_table" gorm:"foreignKey:RouterID"`
	Interfaces   []NetworkInterface `json:"interfaces" gorm:"foreignKey:RouterID"`
	NATConfig    NAT                `json:"nat_config" gorm:"foreignKey:RouterID"`
	Status       string             `json:"status"`
}

type Route struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	RouterID    uint32 `json:"router_id"`
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Metric      int    `json:"metric"`
}

type NetworkInterface struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	RouterID uint32 `json:"router_id"`
	Name     string `json:"name"`
	IP       string `json:"ip"`
	MAC      string `json:"mac"`
	Status   string `json:"status"`
}

type NAT struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	RouterID uint32    `json:"router_id"`
	Enabled  bool      `json:"enabled"`
	Rules    []NATRule `json:"rules" gorm:"foreignKey:NATID"`
}

type NATRule struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	NATID         uint   `json:"nat_id"`
	SourceIP      string `json:"source_ip"`
	DestinationIP string `json:"destination_ip"`
	Protocol      string `json:"protocol"`
}

type Switch struct {
	DeviceID   string `json:"device_id" gorm:"primaryKey"`
	MACAddress string `json:"mac_address"`
	Ports      []Port `json:"ports" gorm:"foreignKey:SwitchID"`
	VLANs      []VLAN `json:"vlans" gorm:"foreignKey:SwitchID"`
	Status     string `json:"status"`
}

type Port struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	SwitchID   string `json:"switch_id"`
	Number     uint32 `json:"number"`
	MACAddress string `json:"mac_address"`
	Status     string `json:"status"`
}

type VLAN struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	SwitchID string `json:"switch_id"`
	Name     string `json:"name"`
	Ports    []int  `json:"ports" gorm:"type:integer[]"`
}

type AccessPoint struct {
	DeviceID         uint32            `json:"device_id" gorm:"primaryKey"`
	SSID             uint32            `json:"ssid"`
	Channel          int               `json:"channel"`
	ConnectedDevices []ConnectedDevice `json:"connected_devices" gorm:"foreignKey:AccessPointID"`
	SignalStrength   int               `json:"signal_strength"`
	Status           string            `json:"status"`
}

type ConnectedDevice struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	AccessPointID  uint32 `json:"access_point_id"`
	MACAddress     string `json:"mac_address"`
	IP             uint32 `json:"ip"`
	SignalStrength uint32 `json:"signal_strength"`
}

type Firewall struct {
	DeviceID uint32         `json:"device_id" gorm:"primaryKey"`
	Rules    []FirewallRule `json:"rules" gorm:"foreignKey:FirewallID"`
	Logs     []FirewallLog  `json:"logs" gorm:"foreignKey:FirewallID"`
	Status   string         `json:"status"`
}

type FirewallRule struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	FirewallID    uint32 `json:"firewall_id"`
	Action        string `json:"action"`
	SourceIP      string `json:"source_ip"`
	DestinationIP string `json:"destination_ip"`
	Protocol      string `json:"protocol"`
	Port          int    `json:"port"`
}

type FirewallLog struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	FirewallID    uint32    `json:"firewall_id"`
	Timestamp     time.Time `json:"timestamp"`
	SourceIP      string    `json:"source_ip"`
	DestinationIP string    `json:"destination_ip"`
	Action        string    `json:"action"`
}

type OLT struct {
	DeviceID   uint32 `json:"device_id" gorm:"primaryKey"`
	IPAddress  string `json:"ip_address"`
	MACAddress string `json:"mac_address"`
	ONUs       []ONU  `json:"onus" gorm:"foreignKey:OLTID"`
	VLANs      []VLAN `json:"vlans" gorm:"foreignKey:ID"`
	Status     string `json:"status"`
}

type ONU struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	OLTID            uint32 `json:"olt_id"`
	DeviceID         uint32 `json:"device_id"`
	SerialNumber     string `json:"serial_number"`
	IPAddress        string `json:"ip_address"`
	MACAddress       string `json:"mac_address"`
	ConnectionStatus string `json:"connection_status"`
	ConnectionType   string `json:"connection_type"`
	UpstreamSpeed    int    `json:"upstream_speed"`
	DownstreamSpeed  int    `json:"downstream_speed"`
	Status           string `json:"status"`
}

type PONPort struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	PortNumber     uint32 `json:"port_number"`
	Status         string `json:"status"`
	ONUs           []ONU  `json:"onus" gorm:"foreignKey:DeviceID"`
	SignalStrength int    `json:"signal_strength"`
}

// ------------ Общая структура VLAN оборудования

type NetworkDevice struct {
	DeviceID   uint32 `json:"device_id" gorm:"primaryKey"`
	DeviceType string `json:"device_type"` // "router", "switch", "access_point", "firewall"
	Status     string `json:"status"`
}

// ------- Общая структура GPon оборудования
type GPONDevice struct {
	DeviceID   uint32 `json:"device_id" gorm:"primaryKey"`
	DeviceType string `json:"device_type"` // "olt", "onu"
	Status     string `json:"status"`
}
