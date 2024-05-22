package inpatient_wl

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_waiting_list.go

// CreateWaitingListEntry - Saves new entry into waiting list
func (this *implInpatientWaitingListAPI) CreateWaitingListEntry(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance,  interface{},  int){
    var entry WaitingListEntry

    if err := c.ShouldBindJSON(&entry); err != nil {
        return nil, gin.H{
            "status": http.StatusBadRequest,
            "message": "Invalid request body",
            "error": err.Error(),
        }, http.StatusBadRequest
    }

    if entry.PatientId == "" {
        return nil, gin.H{
            "status": http.StatusBadRequest,
            "message": "Patient ID is required",
        }, http.StatusBadRequest
    }

    if entry.Id == "" || entry.Id == "@new" {
        entry.Id = uuid.NewString()
    }

    conflictIndx := slices.IndexFunc( ambulance.WaitingList, func(waiting WaitingListEntry) bool {
        return entry.Id == waiting.Id || entry.PatientId == waiting.PatientId
    })

    if conflictIndx >= 0 {
        return nil, gin.H{
            "status": http.StatusConflict,
            "message": "Entry already exists",
        }, http.StatusConflict
    }

    ambulance.WaitingList = append(ambulance.WaitingList, entry)
    // entry was copied by value return reconciled value from the list
    entryIndx := slices.IndexFunc( ambulance.WaitingList, func(waiting WaitingListEntry) bool {
        return entry.Id == waiting.Id
    })
    if entryIndx < 0 {
        return nil, gin.H{
            "status": http.StatusInternalServerError,
            "message": "Failed to save entry",
        }, http.StatusInternalServerError
    }
    return ambulance, ambulance.WaitingList[entryIndx], http.StatusOK
  })
}

// DeleteWaitingListEntry - Deletes specific entry
func (this *implInpatientWaitingListAPI) DeleteWaitingListEntry(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
    entryId := ctx.Param("entryId")

    if entryId == "" {
        return nil, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Entry ID is required",
        }, http.StatusBadRequest
    }

    entryIndx := slices.IndexFunc(ambulance.WaitingList, func(waiting WaitingListEntry) bool {
        return entryId == waiting.Id
    })

    if entryIndx < 0 {
        return nil, gin.H{
            "status":  http.StatusNotFound,
            "message": "Entry not found",
        }, http.StatusNotFound
    }

    ambulance.WaitingList = append(ambulance.WaitingList[:entryIndx], ambulance.WaitingList[entryIndx+1:]...)
    return ambulance, nil, http.StatusNoContent
  })
}

// GetWaitingListEntries - Provides the ambulance waiting list
func (this *implInpatientWaitingListAPI) GetWaitingListEntries(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
    result := ambulance.WaitingList
    if result == nil {
        result = []WaitingListEntry{}
    }
    // return nil ambulance - no need to update it in db
    return nil, result, http.StatusOK
  })
}

// GetWaitingListEntry - Provides details about waiting list entry
func (this *implInpatientWaitingListAPI) GetWaitingListEntry(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
    entryId := ctx.Param("entryId")

    if entryId == "" {
        return nil, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Entry ID is required",
        }, http.StatusBadRequest
    }

    entryIndx := slices.IndexFunc(ambulance.WaitingList, func(waiting WaitingListEntry) bool {
        return entryId == waiting.Id
    })

    if entryIndx < 0 {
        return nil, gin.H{
            "status":  http.StatusNotFound,
            "message": "Entry not found",
        }, http.StatusNotFound
    }

    // return nil ambulance - no need to update it in db
    return nil, ambulance.WaitingList[entryIndx], http.StatusOK
})
}

// UpdateWaitingListEntry - Updates specific entry
func (this *implInpatientWaitingListAPI) UpdateWaitingListEntry(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
    var entry WaitingListEntry

    if err := c.ShouldBindJSON(&entry); err != nil {
        return nil, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid request body",
            "error":   err.Error(),
        }, http.StatusBadRequest
    }

    entryId := ctx.Param("entryId")

    if entryId == "" {
        return nil, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Entry ID is required",
        }, http.StatusBadRequest
    }

    entryIndx := slices.IndexFunc(ambulance.WaitingList, func(waiting WaitingListEntry) bool {
        return entryId == waiting.Id
    })

    if entryIndx < 0 {
        return nil, gin.H{
            "status":  http.StatusNotFound,
            "message": "Entry not found",
        }, http.StatusNotFound
    }

    if entry.PatientId != "" {
        ambulance.WaitingList[entryIndx].PatientId = entry.PatientId
    }

    if entry.Id != "" {
        ambulance.WaitingList[entryIndx].Id = entry.Id
    }

    if entry.Room != "" {
      ambulance.WaitingList[entryIndx].Room = entry.Room
    }
    if isInteger(entry.Capacity) {
      ambulance.WaitingList[entryIndx].Capacity = entry.Capacity
    }
    if isInteger(entry.AllocatedCapacity) {
      ambulance.WaitingList[entryIndx].AllocatedCapacity = entry.AllocatedCapacity
    }
    if isInteger(entry.FreeCapacity) {
      ambulance.WaitingList[entryIndx].FreeCapacity = entry.FreeCapacity
    }
    if isInteger(entry.ToPrepareCapacity) {
      ambulance.WaitingList[entryIndx].ToPrepareCapacity = entry.ToPrepareCapacity
    }

    return ambulance, ambulance.WaitingList[entryIndx], http.StatusOK
  })
}

// isInteger checks if the provided value is an integer type.
func isInteger(value interface{}) bool {
  switch value.(type) {
  case int, int8, int16, int32, int64,
      uint, uint8, uint16, uint32, uint64:
      return true
  default:
      return false
  }
}