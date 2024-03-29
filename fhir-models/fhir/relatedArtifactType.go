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

// RelatedArtifactType is documented here http://hl7.org/fhir/ValueSet/related-artifact-type
type RelatedArtifactType int

const (
	RelatedArtifactTypeDocumentation RelatedArtifactType = iota
	RelatedArtifactTypeJustification
	RelatedArtifactTypeCitation
	RelatedArtifactTypePredecessor
	RelatedArtifactTypeSuccessor
	RelatedArtifactTypeDerivedFrom
	RelatedArtifactTypeDependsOn
	RelatedArtifactTypeComposedOf
)

func (code RelatedArtifactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RelatedArtifactType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "documentation":
		*code = RelatedArtifactTypeDocumentation
	case "justification":
		*code = RelatedArtifactTypeJustification
	case "citation":
		*code = RelatedArtifactTypeCitation
	case "predecessor":
		*code = RelatedArtifactTypePredecessor
	case "successor":
		*code = RelatedArtifactTypeSuccessor
	case "derived-from":
		*code = RelatedArtifactTypeDerivedFrom
	case "depends-on":
		*code = RelatedArtifactTypeDependsOn
	case "composed-of":
		*code = RelatedArtifactTypeComposedOf
	default:
		return fmt.Errorf("unknown RelatedArtifactType code `%s`", s)
	}
	return nil
}
func (code RelatedArtifactType) String() string {
	return code.Code()
}
func (code RelatedArtifactType) Code() string {
	switch code {
	case RelatedArtifactTypeDocumentation:
		return "documentation"
	case RelatedArtifactTypeJustification:
		return "justification"
	case RelatedArtifactTypeCitation:
		return "citation"
	case RelatedArtifactTypePredecessor:
		return "predecessor"
	case RelatedArtifactTypeSuccessor:
		return "successor"
	case RelatedArtifactTypeDerivedFrom:
		return "derived-from"
	case RelatedArtifactTypeDependsOn:
		return "depends-on"
	case RelatedArtifactTypeComposedOf:
		return "composed-of"
	}
	return "<unknown>"
}
func (code RelatedArtifactType) Display() string {
	switch code {
	case RelatedArtifactTypeDocumentation:
		return "Documentation"
	case RelatedArtifactTypeJustification:
		return "Justification"
	case RelatedArtifactTypeCitation:
		return "Citation"
	case RelatedArtifactTypePredecessor:
		return "Predecessor"
	case RelatedArtifactTypeSuccessor:
		return "Successor"
	case RelatedArtifactTypeDerivedFrom:
		return "Derived From"
	case RelatedArtifactTypeDependsOn:
		return "Depends On"
	case RelatedArtifactTypeComposedOf:
		return "Composed Of"
	}
	return "<unknown>"
}
func (code RelatedArtifactType) Definition() string {
	switch code {
	case RelatedArtifactTypeDocumentation:
		return "Additional documentation for the knowledge resource. This would include additional instructions on usage as well as additional information on clinical context or appropriateness."
	case RelatedArtifactTypeJustification:
		return "A summary of the justification for the knowledge resource including supporting evidence, relevant guidelines, or other clinically important information. This information is intended to provide a way to make the justification for the knowledge resource available to the consumer of interventions or results produced by the knowledge resource."
	case RelatedArtifactTypeCitation:
		return "Bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource."
	case RelatedArtifactTypePredecessor:
		return "The previous version of the knowledge resource."
	case RelatedArtifactTypeSuccessor:
		return "The next version of the knowledge resource."
	case RelatedArtifactTypeDerivedFrom:
		return "The knowledge resource is derived from the related artifact. This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but is modified to capture either a different set of overall requirements, or a more specific set of requirements such as those involved in a particular institution or clinical setting."
	case RelatedArtifactTypeDependsOn:
		return "The knowledge resource depends on the given related artifact."
	case RelatedArtifactTypeComposedOf:
		return "The knowledge resource is composed of the given related artifact."
	}
	return "<unknown>"
}
