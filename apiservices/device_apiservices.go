package apiservices

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yukiz97/cls-device-services/lcservices"
	"github.com/yukiz97/cls-device-services/models"
	"github.com/yukiz97/utils/restapi"
	"log"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, _ *http.Request) {
	restapi.RespondWithJSON(response, http.StatusOK, "Welcome to restful API of cls - device services")
}

func insertDevice(response http.ResponseWriter, request *http.Request) {
	modelInput := models.Device{}

	err := json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.IDProduct == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`idproduct` must be greater than 0")
		return
	} else if modelInput.IDCustomer == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`idcustomer` must be greater than 0")
		return
	} else if modelInput.DeviceCode == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`devicecode` must not be empty")
		return
	} else if modelInput.DeviceSerial == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`deviceserial` must not be empty")
		return
	} else if modelInput.BuyDate == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`buydate` must not be empty")
		return
	} else if modelInput.GuaranteeExpireDate == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`guaranteeexpiredate` must not be empty")
		return
	} else if modelInput.GuaranteeYears == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`guaranteeexpireyears` must not be empty")
		return
	}

	idInserted := lcservices.InsertDevice(modelInput)

	if idInserted != 0 {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idInserted, Message: "Insert new device successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Insert new device fail, please try again!")
	}
}

func updateDevice(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idDevice, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID device must be a integer")
		return
	}

	modelInput := models.Device{}
	err = json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.IDProduct == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`idproduct` must be greater than 0")
		return
	} else if modelInput.IDCustomer == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`idcustomer` must be greater than 0")
		return
	} else if modelInput.DeviceCode == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`devicecode` must not be empty")
		return
	} else if modelInput.DeviceSerial == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`deviceserial` must not be empty")
		return
	} else if modelInput.BuyDate == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`buydate` must not be empty")
		return
	} else if modelInput.GuaranteeExpireDate == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`guaranteeexpiredate` must not be empty")
		return
	} else if modelInput.GuaranteeYears == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "`guaranteeexpireyears` must not be empty")
		return
	}

	modelInput.ID = idDevice
	isUpdated := lcservices.UpdateDevice(modelInput)

	if isUpdated {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idDevice, Message: "Updated device successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Device with ID "+strconv.Itoa(idDevice)+" doesn't exists or value doesn't change!")
	}
}

func deleteDevice(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idDevice, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID device must be a integer")
		return
	}

	isDeleted := lcservices.DeleteDevice(idDevice)

	if isDeleted {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idDevice, Message: "Deleted device successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Device with ID "+strconv.Itoa(idDevice)+" doesn't exists!")
	}
}

func getDeviceList(response http.ResponseWriter, _ *http.Request) {
	listDevice := lcservices.GetDeviceList("")

	restapi.RespondWithJSON(response, http.StatusOK, listDevice)
}

func getDevice(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idDevice, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID device must be a integer")
		return
	}

	modelDevice := lcservices.GetDeviceByID(idDevice)
	if modelDevice.ID == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "Device with ID "+strconv.Itoa(idDevice)+" doest not exist!")
		return
	}

	restapi.RespondWithJSON(response, http.StatusOK, modelDevice)
}

func searchDeviceList(response http.ResponseWriter, request *http.Request) {
	modelInput := models.SearchDevice{}

	err := json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.Keyword == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`keyword` must not be empty")
		return
	}

	listDevice := lcservices.GetDeviceList(modelInput.Keyword)

	restapi.RespondWithJSON(response, http.StatusOK, listDevice)
}

//InitRestfulAPIServices init customer restfull api
func InitRestfulAPIServices(listenPort int) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	router.HandleFunc("/getDeviceList/", getDeviceList).Methods("GET")
	router.HandleFunc("/getDevice/id/{id}", getDevice).Methods("GET")

	router.HandleFunc("/insertDevice/", insertDevice).Methods("POST")
	router.HandleFunc("/searchDeviceList/", searchDeviceList).Methods("POST")

	router.HandleFunc("/updateDevice/id/{id}", updateDevice).Methods("PUT")

	router.HandleFunc("/deleteDevice/id/{id}", deleteDevice).Methods("DELETE")

	println("Running CLS - device services.... - Listen to port :" + strconv.Itoa(listenPort))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listenPort), router))
}
