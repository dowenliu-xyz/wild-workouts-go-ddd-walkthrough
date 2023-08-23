// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package main

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// PostTraining defines model for PostTraining.
type PostTraining struct {
	Notes string    `json:"notes"`
	Time  time.Time `json:"time"`
}

// Training defines model for Training.
type Training struct {
	CanBeCancelled     bool       `json:"canBeCancelled"`
	MoveProposedBy     *string    `json:"moveProposedBy,omitempty"`
	MoveRequiresAccept bool       `json:"moveRequiresAccept"`
	Notes              string     `json:"notes"`
	ProposedTime       *time.Time `json:"proposedTime,omitempty"`
	Time               time.Time  `json:"time"`
	User               string     `json:"user"`
	UserUuid           string     `json:"userUuid"`
	Uuid               string     `json:"uuid"`
}

// Trainings defines model for Trainings.
type Trainings struct {
	Trainings []Training `json:"trainings"`
}

// CreateTrainingJSONRequestBody defines body for CreateTraining for application/json ContentType.
type CreateTrainingJSONRequestBody = PostTraining

// RescheduleTrainingJSONRequestBody defines body for RescheduleTraining for application/json ContentType.
type RescheduleTrainingJSONRequestBody = PostTraining
