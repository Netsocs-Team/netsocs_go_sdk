package apis

/*
{
  topicKey : "key",
  dataValue: {...},
  idDevice: 1
}
*/

type RequestConfigToDeviceBody struct {
	TopicKey  string `json:"topicKey"`
	DataValue string `json:"dataValue"`
	IdDevice  int    `json:"idDevice"`
}
