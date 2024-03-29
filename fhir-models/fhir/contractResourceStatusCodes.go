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

// ContractResourceStatusCodes is documented here http://hl7.org/fhir/ValueSet/contract-status
type ContractResourceStatusCodes int

const (
	ContractResourceStatusCodesAmended ContractResourceStatusCodes = iota
	ContractResourceStatusCodesAppended
	ContractResourceStatusCodesCancelled
	ContractResourceStatusCodesDisputed
	ContractResourceStatusCodesEnteredInError
	ContractResourceStatusCodesExecutable
	ContractResourceStatusCodesExecuted
	ContractResourceStatusCodesNegotiable
	ContractResourceStatusCodesOffered
	ContractResourceStatusCodesPolicy
	ContractResourceStatusCodesRejected
	ContractResourceStatusCodesRenewed
	ContractResourceStatusCodesRevoked
	ContractResourceStatusCodesResolved
	ContractResourceStatusCodesTerminated
)

func (code ContractResourceStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContractResourceStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "amended":
		*code = ContractResourceStatusCodesAmended
	case "appended":
		*code = ContractResourceStatusCodesAppended
	case "cancelled":
		*code = ContractResourceStatusCodesCancelled
	case "disputed":
		*code = ContractResourceStatusCodesDisputed
	case "entered-in-error":
		*code = ContractResourceStatusCodesEnteredInError
	case "executable":
		*code = ContractResourceStatusCodesExecutable
	case "executed":
		*code = ContractResourceStatusCodesExecuted
	case "negotiable":
		*code = ContractResourceStatusCodesNegotiable
	case "offered":
		*code = ContractResourceStatusCodesOffered
	case "policy":
		*code = ContractResourceStatusCodesPolicy
	case "rejected":
		*code = ContractResourceStatusCodesRejected
	case "renewed":
		*code = ContractResourceStatusCodesRenewed
	case "revoked":
		*code = ContractResourceStatusCodesRevoked
	case "resolved":
		*code = ContractResourceStatusCodesResolved
	case "terminated":
		*code = ContractResourceStatusCodesTerminated
	default:
		return fmt.Errorf("unknown ContractResourceStatusCodes code `%s`", s)
	}
	return nil
}
func (code ContractResourceStatusCodes) String() string {
	return code.Code()
}
func (code ContractResourceStatusCodes) Code() string {
	switch code {
	case ContractResourceStatusCodesAmended:
		return "amended"
	case ContractResourceStatusCodesAppended:
		return "appended"
	case ContractResourceStatusCodesCancelled:
		return "cancelled"
	case ContractResourceStatusCodesDisputed:
		return "disputed"
	case ContractResourceStatusCodesEnteredInError:
		return "entered-in-error"
	case ContractResourceStatusCodesExecutable:
		return "executable"
	case ContractResourceStatusCodesExecuted:
		return "executed"
	case ContractResourceStatusCodesNegotiable:
		return "negotiable"
	case ContractResourceStatusCodesOffered:
		return "offered"
	case ContractResourceStatusCodesPolicy:
		return "policy"
	case ContractResourceStatusCodesRejected:
		return "rejected"
	case ContractResourceStatusCodesRenewed:
		return "renewed"
	case ContractResourceStatusCodesRevoked:
		return "revoked"
	case ContractResourceStatusCodesResolved:
		return "resolved"
	case ContractResourceStatusCodesTerminated:
		return "terminated"
	}
	return "<unknown>"
}
func (code ContractResourceStatusCodes) Display() string {
	switch code {
	case ContractResourceStatusCodesAmended:
		return "Amended"
	case ContractResourceStatusCodesAppended:
		return "Appended"
	case ContractResourceStatusCodesCancelled:
		return "Cancelled"
	case ContractResourceStatusCodesDisputed:
		return "Disputed"
	case ContractResourceStatusCodesEnteredInError:
		return "Entered in Error"
	case ContractResourceStatusCodesExecutable:
		return "Executable"
	case ContractResourceStatusCodesExecuted:
		return "Executed"
	case ContractResourceStatusCodesNegotiable:
		return "Negotiable"
	case ContractResourceStatusCodesOffered:
		return "Offered"
	case ContractResourceStatusCodesPolicy:
		return "Policy"
	case ContractResourceStatusCodesRejected:
		return "Rejected"
	case ContractResourceStatusCodesRenewed:
		return "Renewed"
	case ContractResourceStatusCodesRevoked:
		return "Revoked"
	case ContractResourceStatusCodesResolved:
		return "Resolved"
	case ContractResourceStatusCodesTerminated:
		return "Terminated"
	}
	return "<unknown>"
}
func (code ContractResourceStatusCodes) Definition() string {
	switch code {
	case ContractResourceStatusCodesAmended:
		return "Contract is augmented with additional information to correct errors in a predecessor or to updated values in a predecessor. Usage: Contract altered within effective time. Precedence Order = 9. Comparable FHIR and v.3 status codes: revised; replaced."
	case ContractResourceStatusCodesAppended:
		return "Contract is augmented with additional information that was missing from a predecessor Contract. Usage: Contract altered within effective time. Precedence Order = 9. Comparable FHIR and v.3 status codes: updated, replaced."
	case ContractResourceStatusCodesCancelled:
		return "Contract is terminated due to failure of the Grantor and/or the Grantee to fulfil one or more contract provisions. Usage: Abnormal contract termination. Precedence Order = 10. Comparable FHIR and v.3 status codes: stopped; failed; aborted."
	case ContractResourceStatusCodesDisputed:
		return "Contract is pended to rectify failure of the Grantor or the Grantee to fulfil contract provision(s). E.g., Grantee complaint about Grantor's failure to comply with contract provisions. Usage: Contract pended. Precedence Order = 7. Comparable FHIR and v.3 status codes: on hold; pended; suspended."
	case ContractResourceStatusCodesEnteredInError:
		return "Contract was created in error. No Precedence Order.  Status may be applied to a Contract with any status."
	case ContractResourceStatusCodesExecutable:
		return "Contract execution pending; may be executed when either the Grantor or the Grantee accepts the contract provisions by signing. I.e., where either the Grantor or the Grantee has signed, but not both. E.g., when an insurance applicant signs the insurers application, which references the policy.\u00a0Usage: Optional first step of contract execution activity.  May be skipped and contracting activity moves directly to executed state. Precedence Order = 3. Comparable FHIR and v.3 status codes: draft; preliminary; planned; intended; active."
	case ContractResourceStatusCodesExecuted:
		return "Contract is activated for period stipulated when both the Grantor and Grantee have signed it. Usage: Required state for normal completion of contracting activity.  Precedence Order = 6. Comparable FHIR and v.3 status codes: accepted; completed."
	case ContractResourceStatusCodesNegotiable:
		return "Contract execution is suspended while either or both the Grantor and Grantee propose and consider new or revised contract provisions. I.e., where the party which has not signed proposes changes to the terms.  E .g., a life insurer declines to agree to the signed application because the life insurer has evidence that the applicant, who asserted to being younger or a non-smoker to get a lower premium rate - but offers instead to agree to a higher premium based on the applicants actual age or smoking status. Usage: Optional contract activity between executable and executed state. Precedence Order = 4. Comparable FHIR and v.3 status codes: in progress; review; held."
	case ContractResourceStatusCodesOffered:
		return "Contract is a proposal by either the Grantor or the Grantee. Aka - A Contract hard copy or electronic 'template', 'form' or 'application'. E.g., health insurance application; consent directive form. Usage: Beginning of contract negotiation, which may have been completed as a precondition because used for 0..* contracts. Precedence Order = 2. Comparable FHIR and v.3 status codes: requested; new."
	case ContractResourceStatusCodesPolicy:
		return "Contract template is available as the basis for an application or offer by the Grantor or Grantee. E.g., health insurance policy; consent directive policy.  Usage: Required initial contract activity, which may have been completed as a precondition because used for 0..* contracts. Precedence Order = 1. Comparable FHIR and v.3 status codes: proposed; intended."
	case ContractResourceStatusCodesRejected:
		return " Execution of the Contract is not completed because either or both the Grantor and Grantee decline to accept some or all of the contract provisions. Usage: Optional contract activity between executable and abnormal termination. Precedence Order = 5. Comparable FHIR and v.3 status codes:  stopped; cancelled."
	case ContractResourceStatusCodesRenewed:
		return "Beginning of a successor Contract at the termination of predecessor Contract lifecycle. Usage: Follows termination of a preceding Contract that has reached its expiry date. Precedence Order = 13. Comparable FHIR and v.3 status codes: superseded."
	case ContractResourceStatusCodesRevoked:
		return "A Contract that is rescinded.  May be required prior to replacing with an updated Contract. Comparable FHIR and v.3 status codes: nullified."
	case ContractResourceStatusCodesResolved:
		return "Contract is reactivated after being pended because of faulty execution. *E.g., competency of the signer(s), or where the policy is substantially different from and did not accompany the application/form so that the applicant could not compare them. Aka - ''reactivated''. Usage: Optional stage where a pended contract is reactivated. Precedence Order = 8. Comparable FHIR and v.3 status codes: reactivated."
	case ContractResourceStatusCodesTerminated:
		return "Contract reaches its expiry date.\u00a0It might or might not be renewed or renegotiated. Usage: Normal end of contract period. Precedence Order = 12. Comparable FHIR and v.3 status codes: Obsoleted."
	}
	return "<unknown>"
}
