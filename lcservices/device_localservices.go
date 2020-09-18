package lcservices

import (
	"github.com/yukiz97/cls-device-services/models"
	"github.com/yukiz97/utils/date"
	"github.com/yukiz97/utils/dbcon"
	"time"
)

var strDBConnect string
var mapDeviceField map[string]string

//InsertDevice insert device by value of struct
func InsertDevice(modelDevice models.Device) int64 {
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	insertQuery, err := db.Prepare("INSERT INTO device(" +
		"" + mapDeviceField["idproduct"] + ", " +
		"" + mapDeviceField["idcustomer"] + ", " +
		"" + mapDeviceField["code"] + ", " +
		"" + mapDeviceField["serial"] + ", " +
		"" + mapDeviceField["buydate"] + ", " +
		"" + mapDeviceField["expiredate"] + ", " +
		"" + mapDeviceField["expireyears"] + ") " +
		"VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	result, errInsert := insertQuery.Exec(
		modelDevice.IDProduct,
		modelDevice.IDCustomer,
		modelDevice.DeviceCode,
		modelDevice.DeviceSerial,
		modelDevice.BuyDate,
		modelDevice.GuaranteeExpireDate,
		modelDevice.GuaranteeYears,
	)

	if errInsert != nil {
		panic(errInsert)
	}

	idInserted, _ := result.LastInsertId()

	return idInserted
}

//UpdateDevice update device by value of struct
func UpdateDevice(modelDevice models.Device) bool {
	isUpdated := true
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	updateQuery, err := db.Prepare("UPDATE device SET " +
		"" + mapDeviceField["serial"] + " = ?, " +
		"" + mapDeviceField["buydate"] + " = ?, " +
		"" + mapDeviceField["expiredate"] + " = ?, " +
		"" + mapDeviceField["expireyears"] + "  = ? " +
		" WHERE " + mapDeviceField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, errUpdate := updateQuery.Exec(
		modelDevice.DeviceSerial,
		modelDevice.BuyDate,
		modelDevice.GuaranteeExpireDate,
		modelDevice.GuaranteeYears,
		modelDevice.ID,
	)

	if errUpdate != nil {
		panic(errUpdate)
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected == 0 {
		isUpdated = false
	}

	return isUpdated
}

//DeleteDevice delete device by id
func DeleteDevice(idDevice int) bool {
	isDeleted := true
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	deleteQuery, err := db.Prepare("DELETE FROM device WHERE " + mapDeviceField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, errDelete := deleteQuery.Exec(idDevice)

	if errDelete != nil {
		panic(errDelete)
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected == 0 {
		isDeleted = false
	}

	return isDeleted
}

//GetDeviceList get device list by keyword
func GetDeviceList(keyWord string) ([]models.Device,[]int) {
	listDevice := make([]models.Device, 0)
	listID := make([]int, 0)
	keyWord = "%" + keyWord + "%"

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()

	selectQuery, err := db.Prepare("SELECT * FROM device WHERE " + mapDeviceField["code"] + " LIKE ?")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query(keyWord)
	for result.Next() {
		var buyDate time.Time
		var expireDate time.Time
		var createDate time.Time
		modelDevice := models.Device{}

		result.Scan(
			&modelDevice.ID,
			&modelDevice.IDProduct,
			&modelDevice.IDCustomer,
			&modelDevice.DeviceCode,
			&modelDevice.DeviceSerial,
			&buyDate,
			&expireDate,
			&modelDevice.GuaranteeYears,
			&createDate,
		)
		modelDevice.BuyDate = date.FormatTimeToString(buyDate, date.Format2)
		modelDevice.GuaranteeExpireDate = date.FormatTimeToString(expireDate, date.Format2)
		modelDevice.CreateDate = date.FormatTimeToString(createDate, date.Format1)
		listDevice = append(listDevice, modelDevice)
		listID = append(listID,modelDevice.ID)
	}

	return listDevice, listID
}

//GetDeviceByID get device by device id
func GetDeviceByID(idDevice int) models.Device {
	var modelDevice models.Device

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()

	selectQuery, err := db.Prepare("SELECT * FROM device WHERE " + mapDeviceField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query(idDevice)
	for result.Next() {
		var buyDate time.Time
		var expireDate time.Time
		var createDate time.Time

		modelDevice = models.Device{}

		result.Scan(
			&modelDevice.ID,
			&modelDevice.IDProduct,
			&modelDevice.IDCustomer,
			&modelDevice.DeviceCode,
			&modelDevice.DeviceSerial,
			&buyDate,
			&expireDate,
			&modelDevice.GuaranteeYears,
			&createDate,
		)
		modelDevice.BuyDate = date.FormatTimeToString(buyDate, date.Format2)
		modelDevice.GuaranteeExpireDate = date.FormatTimeToString(expireDate, date.Format2)
		modelDevice.CreateDate = date.FormatTimeToString(createDate, date.Format1)
	}

	return modelDevice
}

//InitLocalServices init value for database functions
func InitLocalServices(host string, userName string, password string, db string) {
	strDBConnect = dbcon.GetMySQLOpenConnectString(host, userName, password, db)

	mapDeviceField = make(map[string]string)
	mapDeviceField["id"] = "IdDevice"
	mapDeviceField["idproduct"] = "IdProduct"
	mapDeviceField["idcustomer"] = "IdCustomer"
	mapDeviceField["code"] = "DeviceCode"
	mapDeviceField["serial"] = "DeviceSerial"
	mapDeviceField["buydate"] = "BuyDate"
	mapDeviceField["expiredate"] = "GuaranteeExpireDate"
	mapDeviceField["expireyears"] = "GuaranteeYears"
	mapDeviceField["createdate"] = "CreateDate"
}
