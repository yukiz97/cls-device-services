package lcservices

import (
	"github.com/yukiz97/cls-device-services/models"
	"github.com/yukiz97/utils"
	"github.com/yukiz97/utils/date"
	"github.com/yukiz97/utils/dbcon"
	"time"
)

//GetDeviceLicenseSimplifyInfo get device license info of arr device id
func GetDeviceLicenseSimplifyInfo(arrId []int) map[int]models.DeviceLicenseSimplifyInfo{
	if len(arrId) == 0 {
		return nil
	}
	mapLicenseInfo := make(map[int]models.DeviceLicenseSimplifyInfo)

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	strIds := utils.IntArrayToString(arrId,",")
	selectQuery, err := db.Prepare("SELECT IdLicense,LicenseCode,ProvideDate,ExpireDate,CreateDate,IdDevice FROM license WHERE IdDevice IN ("+strIds+")")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query()

	for result.Next() {
		var provideDate time.Time
		var expireDate time.Time
		var createDate time.Time

		var idDevice int
		modelLicense := models.DeviceLicenseSimplifyInfo{}

		result.Scan(&modelLicense.ID, &modelLicense.Code, &provideDate, &expireDate,&createDate,&idDevice)

		modelLicense.ProvideDate = date.FormatTimeToString(provideDate, date.Format2)
		modelLicense.ExpireDate = date.FormatTimeToString(expireDate, date.Format2)
		modelLicense.CreateDate = date.FormatTimeToString(createDate, date.Format1)

		mapLicenseInfo[idDevice] = modelLicense
	}

	return mapLicenseInfo
}

//GetDeviceCustomerSimplifyInfo get device customer info of arr customer id
func GetDeviceCustomerSimplifyInfo(arrId []int) map[int]models.DeviceCustomerSimplifyInfo{
	if len(arrId) == 0 {
		return nil
	}
	mapCustomerInfo := make(map[int]models.DeviceCustomerSimplifyInfo)

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	strIds := utils.IntArrayToString(arrId,",")
	selectQuery, err := db.Prepare("SELECT customer.IdCustomer,CustomerName,CustomerAddress,IdDevice FROM customer JOIN device ON customer.IdCustomer = device.IdCustomer WHERE IdDevice IN ("+strIds+")")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query()

	for result.Next() {
		var idDevice int
		modelCustomer := models.DeviceCustomerSimplifyInfo{}

		result.Scan(&modelCustomer.ID, &modelCustomer.Name, &modelCustomer.Address,&idDevice)


		mapCustomerInfo[idDevice] = modelCustomer
	}

	return mapCustomerInfo
}