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

type WaitingListEntry struct {

	// Unique id of the entry in this waiting list
	Id string `json:"id"`

	// Name of inpatient room in waiting list
	Room string `json:"room,omitempty"`

	// Unique identifier of the patient known to Web-In-Cloud system
	PatientId string `json:"patientId,omitempty"`

	// Estimated duration of inpatient visit. If not provided then it will be computed based on condition and inpatient settings
	Capacity int32 `json:"capacity"`

	// Estimated duration of inpatient visit. If not provided then it will be computed based on condition and inpatient settings
	AllocatedCapacity int32 `json:"allocatedCapacity"`

	// Estimated duration of inpatient visit. If not provided then it will be computed based on condition and inpatient settings
	FreeCapacity int32 `json:"freeCapacity,omitempty"`

	// Estimated duration of inpatient visit. If not provided then it will be computed based on condition and inpatient settings
	ToPrepareCapacity int32 `json:"toPrepareCapacity,omitempty"`
}
