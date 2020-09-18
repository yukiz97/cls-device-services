package models

//SearchDevice struct for search device
type SearchDevice struct {
	Keyword string `json:"keyword"`
}

//Device struct present database device table
type Device struct {
	ID                  int    `json:"id"`
	IDProduct           int    `json:"idproduct"`
	IDCustomer          int    `json:"idcustomer"`
	DeviceCode          string `json:"devicecode"`
	DeviceSerial        string `json:"deviceserial"`
	BuyDate             string `json:"buydate"`
	GuaranteeExpireDate string `json:"guaranteeexpiredate"`
	GuaranteeYears      uint8  `json:"guaranteeexpireyears"`
	CreateDate          string `json:"createdate"`
}

type DeviceWithAdditionalInfo struct {
	Device
	License DeviceLicenseSimplifyInfo `json:"license"`
	Customer DeviceCustomerSimplifyInfo `json:"customer"`
}

type DeviceLicenseSimplifyInfo struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	ProvideDate string `json:"providedate"`
	ExpireDate  string `json:"expiredate"`
	CreateDate  string `json:"createdate"`
}

type DeviceCustomerSimplifyInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
