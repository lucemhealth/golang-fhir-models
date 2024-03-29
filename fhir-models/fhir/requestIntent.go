// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/lucemhealth/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// RequestIntent is documented here http://hl7.org/fhir/ValueSet/request-intent
type RequestIntent int

const (
	RequestIntentProposal RequestIntent = iota
	RequestIntentPlan
	RequestIntentDirective
	RequestIntentOrder
	RequestIntentOriginalOrder
	RequestIntentReflexOrder
	RequestIntentFillerOrder
	RequestIntentInstanceOrder
	RequestIntentOption
)

func (code RequestIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposal":
		*code = RequestIntentProposal
	case "plan":
		*code = RequestIntentPlan
	case "directive":
		*code = RequestIntentDirective
	case "order":
		*code = RequestIntentOrder
	case "original-order":
		*code = RequestIntentOriginalOrder
	case "reflex-order":
		*code = RequestIntentReflexOrder
	case "filler-order":
		*code = RequestIntentFillerOrder
	case "instance-order":
		*code = RequestIntentInstanceOrder
	case "option":
		*code = RequestIntentOption
	default:
		return fmt.Errorf("unknown RequestIntent code `%s`", s)
	}
	return nil
}
func (code RequestIntent) String() string {
	return code.Code()
}
func (code RequestIntent) Code() string {
	switch code {
	case RequestIntentProposal:
		return "proposal"
	case RequestIntentPlan:
		return "plan"
	case RequestIntentDirective:
		return "directive"
	case RequestIntentOrder:
		return "order"
	case RequestIntentOriginalOrder:
		return "original-order"
	case RequestIntentReflexOrder:
		return "reflex-order"
	case RequestIntentFillerOrder:
		return "filler-order"
	case RequestIntentInstanceOrder:
		return "instance-order"
	case RequestIntentOption:
		return "option"
	}
	return "<unknown>"
}
func (code RequestIntent) Display() string {
	switch code {
	case RequestIntentProposal:
		return "Proposal"
	case RequestIntentPlan:
		return "Plan"
	case RequestIntentDirective:
		return "Directive"
	case RequestIntentOrder:
		return "Order"
	case RequestIntentOriginalOrder:
		return "Original Order"
	case RequestIntentReflexOrder:
		return "Reflex Order"
	case RequestIntentFillerOrder:
		return "Filler Order"
	case RequestIntentInstanceOrder:
		return "Instance Order"
	case RequestIntentOption:
		return "Option"
	}
	return "<unknown>"
}
func (code RequestIntent) Definition() string {
	switch code {
	case RequestIntentProposal:
		return "The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act."
	case RequestIntentPlan:
		return "The request represents an intention to ensure something occurs without providing an authorization for others to act."
	case RequestIntentDirective:
		return "The request represents a legally binding instruction authored by a Patient or RelatedPerson."
	case RequestIntentOrder:
		return "The request represents a request/demand and authorization for action by a Practitioner."
	case RequestIntentOriginalOrder:
		return "The request represents an original authorization for action."
	case RequestIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization."
	case RequestIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order."
	case RequestIntentInstanceOrder:
		return "An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug."
	case RequestIntentOption:
		return "The request represents a component or option for a RequestGroup that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestGroup]]] for additional information on how this status is used."
	}
	return "<unknown>"
}
