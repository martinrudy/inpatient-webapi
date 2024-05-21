/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: rudolfmato@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package inpatient_wl

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type InpatientWaitingListAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateWaitingListEntry - Saves new entry into waiting list
   CreateWaitingListEntry(ctx *gin.Context)

    // DeleteWaitingListEntry - Deletes specific entry
   DeleteWaitingListEntry(ctx *gin.Context)

    // GetWaitingListEntries - Provides the inpatient waiting list
   GetWaitingListEntries(ctx *gin.Context)

    // GetWaitingListEntry - Provides details about waiting list entry
   GetWaitingListEntry(ctx *gin.Context)

    // UpdateWaitingListEntry - Updates specific entry
   UpdateWaitingListEntry(ctx *gin.Context)

 }

 // partial implementation of InpatientWaitingListAPI - all functions must be implemented in add on files
type implInpatientWaitingListAPI struct {

}

func newInpatientWaitingListAPI() InpatientWaitingListAPI {
  return &implInpatientWaitingListAPI{}
}

func (this *implInpatientWaitingListAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/waiting-list/:ambulanceId/entries", this.CreateWaitingListEntry)
  routerGroup.Handle( http.MethodDelete, "/waiting-list/:ambulanceId/entries/:entryId", this.DeleteWaitingListEntry)
  routerGroup.Handle( http.MethodGet, "/waiting-list/:ambulanceId/entries", this.GetWaitingListEntries)
  routerGroup.Handle( http.MethodGet, "/waiting-list/:ambulanceId/entries/:entryId", this.GetWaitingListEntry)
  routerGroup.Handle( http.MethodPut, "/waiting-list/:ambulanceId/entries/:entryId", this.UpdateWaitingListEntry)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateWaitingListEntry - Saves new entry into waiting list
// func (this *implInpatientWaitingListAPI) CreateWaitingListEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteWaitingListEntry - Deletes specific entry
// func (this *implInpatientWaitingListAPI) DeleteWaitingListEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetWaitingListEntries - Provides the inpatient waiting list
// func (this *implInpatientWaitingListAPI) GetWaitingListEntries(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetWaitingListEntry - Provides details about waiting list entry
// func (this *implInpatientWaitingListAPI) GetWaitingListEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateWaitingListEntry - Updates specific entry
// func (this *implInpatientWaitingListAPI) UpdateWaitingListEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

