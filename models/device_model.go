package models

//SearchCustomer struct for search customer
type SearchDevice struct {
	Keyword string `json:"keyword"`
}

//Device struct present database device talbe
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
