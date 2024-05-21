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
    "github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
  group := engine.Group("/api")
  
  {
    api := newInpatientWaitingListAPI()
    api.addRoutes(group)
  }
  
}