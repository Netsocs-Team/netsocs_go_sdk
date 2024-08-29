package apis

import "encoding/json"

type DeviceManagementApiLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeviceManagementApiLoginResponse struct {
	Auth  bool   `json:"auth"`
	Token string `json:"token"`
}

type DeviceManagementApiDeviceSchemaParams map[string]interface{}

func (d *DeviceManagementApiDeviceSchemaParams) UnmarshalJSON(data []byte) error {
	var params map[string]interface{}
	json.Unmarshal(data, &params)
	*d = params
	return nil
}

type DeviceManagementApiDeviceSchema struct {
	ID                 int    `json:"id"`
	IDManual           int    `json:"idManual"`
	IDSubDevice        int    `json:"idSubDevice"`
	IDSubSystem        int    `json:"idSubSystem"`
	IDLicense          int    `json:"idLicense"`
	IDAccount          int    `json:"idAccount"`
	IDUser             int    `json:"idUser"`
	IDCorporation      int    `json:"idCorporation"`
	IDGateway          int    `json:"idGateway"`
	IDDriver           int    `json:"idDriver"`
	IDDeviceGroup      int    `json:"idDeviceGroup"`
	IDManufacturer     int    `json:"idManufacturer"`
	IDBrand            int    `json:"idBrand"`
	IDModel            int    `json:"idModel"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	PhotoURL           string `json:"photoUrl"`
	DocumentsURL       string `json:"documentsUrl"`
	Map                string `json:"map"`
	Lat                string `json:"lat"`
	Long               string `json:"long"`
	GpsAddress         string `json:"gpsAddress"`
	GpsGeometry        string `json:"gpsGeometry"`
	DirectionalControl string `json:"directionalControl"`
	// Enrollment         bool   `json:"enrollment"`
	// Sync               bool   `json:"sync"`
	Certificate      string                                `json:"certificate"`
	SslPort          int                                   `json:"sslPort"`
	MaintenancePort  int                                   `json:"maintenancePort"`
	Port             int                                   `json:"port"`
	Soap             string                                `json:"soap"`
	API              string                                `json:"api"`
	Socket           string                                `json:"socket"`
	Ipv4Address      string                                `json:"ipv4Address"`
	Ipv4AddressB     string                                `json:"ipv4AddressB"`
	Ipv6Address      string                                `json:"ipv6Address"`
	Ipv6AddressB     string                                `json:"ipv6AddressB"`
	IPAddressPublic  string                                `json:"ipAddressPublic"`
	IPAddressPublicB string                                `json:"ipAddressPublicB"`
	MacAddress       string                                `json:"macAddress"`
	Username         string                                `json:"username"`
	Password         string                                `json:"password"`
	CreatedAt        string                                `json:"createdAt"`
	DisabledAt       string                                `json:"disabledAt"`
	DeletedAt        string                                `json:"deletedAt"`
	IDStatus         int                                   `json:"idStatus"`
	Params           DeviceManagementApiDeviceSchemaParams `json:"params"`
	// IsActive         bool   `json:"isActive"`
}
