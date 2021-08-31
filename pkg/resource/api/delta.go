// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package api

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.APIKeySelectionExpression, b.ko.Spec.APIKeySelectionExpression) {
		delta.Add("Spec.APIKeySelectionExpression", a.ko.Spec.APIKeySelectionExpression, b.ko.Spec.APIKeySelectionExpression)
	} else if a.ko.Spec.APIKeySelectionExpression != nil && b.ko.Spec.APIKeySelectionExpression != nil {
		if *a.ko.Spec.APIKeySelectionExpression != *b.ko.Spec.APIKeySelectionExpression {
			delta.Add("Spec.APIKeySelectionExpression", a.ko.Spec.APIKeySelectionExpression, b.ko.Spec.APIKeySelectionExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Basepath, b.ko.Spec.Basepath) {
		delta.Add("Spec.Basepath", a.ko.Spec.Basepath, b.ko.Spec.Basepath)
	} else if a.ko.Spec.Basepath != nil && b.ko.Spec.Basepath != nil {
		if *a.ko.Spec.Basepath != *b.ko.Spec.Basepath {
			delta.Add("Spec.Basepath", a.ko.Spec.Basepath, b.ko.Spec.Basepath)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Body, b.ko.Spec.Body) {
		delta.Add("Spec.Body", a.ko.Spec.Body, b.ko.Spec.Body)
	} else if a.ko.Spec.Body != nil && b.ko.Spec.Body != nil {
		if *a.ko.Spec.Body != *b.ko.Spec.Body {
			delta.Add("Spec.Body", a.ko.Spec.Body, b.ko.Spec.Body)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CORSConfiguration, b.ko.Spec.CORSConfiguration) {
		delta.Add("Spec.CORSConfiguration", a.ko.Spec.CORSConfiguration, b.ko.Spec.CORSConfiguration)
	} else if a.ko.Spec.CORSConfiguration != nil && b.ko.Spec.CORSConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CORSConfiguration.AllowCredentials, b.ko.Spec.CORSConfiguration.AllowCredentials) {
			delta.Add("Spec.CORSConfiguration.AllowCredentials", a.ko.Spec.CORSConfiguration.AllowCredentials, b.ko.Spec.CORSConfiguration.AllowCredentials)
		} else if a.ko.Spec.CORSConfiguration.AllowCredentials != nil && b.ko.Spec.CORSConfiguration.AllowCredentials != nil {
			if *a.ko.Spec.CORSConfiguration.AllowCredentials != *b.ko.Spec.CORSConfiguration.AllowCredentials {
				delta.Add("Spec.CORSConfiguration.AllowCredentials", a.ko.Spec.CORSConfiguration.AllowCredentials, b.ko.Spec.CORSConfiguration.AllowCredentials)
			}
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORSConfiguration.AllowHeaders, b.ko.Spec.CORSConfiguration.AllowHeaders) {
			delta.Add("Spec.CORSConfiguration.AllowHeaders", a.ko.Spec.CORSConfiguration.AllowHeaders, b.ko.Spec.CORSConfiguration.AllowHeaders)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORSConfiguration.AllowMethods, b.ko.Spec.CORSConfiguration.AllowMethods) {
			delta.Add("Spec.CORSConfiguration.AllowMethods", a.ko.Spec.CORSConfiguration.AllowMethods, b.ko.Spec.CORSConfiguration.AllowMethods)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORSConfiguration.AllowOrigins, b.ko.Spec.CORSConfiguration.AllowOrigins) {
			delta.Add("Spec.CORSConfiguration.AllowOrigins", a.ko.Spec.CORSConfiguration.AllowOrigins, b.ko.Spec.CORSConfiguration.AllowOrigins)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORSConfiguration.ExposeHeaders, b.ko.Spec.CORSConfiguration.ExposeHeaders) {
			delta.Add("Spec.CORSConfiguration.ExposeHeaders", a.ko.Spec.CORSConfiguration.ExposeHeaders, b.ko.Spec.CORSConfiguration.ExposeHeaders)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CORSConfiguration.MaxAge, b.ko.Spec.CORSConfiguration.MaxAge) {
			delta.Add("Spec.CORSConfiguration.MaxAge", a.ko.Spec.CORSConfiguration.MaxAge, b.ko.Spec.CORSConfiguration.MaxAge)
		} else if a.ko.Spec.CORSConfiguration.MaxAge != nil && b.ko.Spec.CORSConfiguration.MaxAge != nil {
			if *a.ko.Spec.CORSConfiguration.MaxAge != *b.ko.Spec.CORSConfiguration.MaxAge {
				delta.Add("Spec.CORSConfiguration.MaxAge", a.ko.Spec.CORSConfiguration.MaxAge, b.ko.Spec.CORSConfiguration.MaxAge)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN) {
		delta.Add("Spec.CredentialsARN", a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN)
	} else if a.ko.Spec.CredentialsARN != nil && b.ko.Spec.CredentialsARN != nil {
		if *a.ko.Spec.CredentialsARN != *b.ko.Spec.CredentialsARN {
			delta.Add("Spec.CredentialsARN", a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DisableExecuteAPIEndpoint, b.ko.Spec.DisableExecuteAPIEndpoint) {
		delta.Add("Spec.DisableExecuteAPIEndpoint", a.ko.Spec.DisableExecuteAPIEndpoint, b.ko.Spec.DisableExecuteAPIEndpoint)
	} else if a.ko.Spec.DisableExecuteAPIEndpoint != nil && b.ko.Spec.DisableExecuteAPIEndpoint != nil {
		if *a.ko.Spec.DisableExecuteAPIEndpoint != *b.ko.Spec.DisableExecuteAPIEndpoint {
			delta.Add("Spec.DisableExecuteAPIEndpoint", a.ko.Spec.DisableExecuteAPIEndpoint, b.ko.Spec.DisableExecuteAPIEndpoint)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DisableSchemaValidation, b.ko.Spec.DisableSchemaValidation) {
		delta.Add("Spec.DisableSchemaValidation", a.ko.Spec.DisableSchemaValidation, b.ko.Spec.DisableSchemaValidation)
	} else if a.ko.Spec.DisableSchemaValidation != nil && b.ko.Spec.DisableSchemaValidation != nil {
		if *a.ko.Spec.DisableSchemaValidation != *b.ko.Spec.DisableSchemaValidation {
			delta.Add("Spec.DisableSchemaValidation", a.ko.Spec.DisableSchemaValidation, b.ko.Spec.DisableSchemaValidation)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FailOnWarnings, b.ko.Spec.FailOnWarnings) {
		delta.Add("Spec.FailOnWarnings", a.ko.Spec.FailOnWarnings, b.ko.Spec.FailOnWarnings)
	} else if a.ko.Spec.FailOnWarnings != nil && b.ko.Spec.FailOnWarnings != nil {
		if *a.ko.Spec.FailOnWarnings != *b.ko.Spec.FailOnWarnings {
			delta.Add("Spec.FailOnWarnings", a.ko.Spec.FailOnWarnings, b.ko.Spec.FailOnWarnings)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProtocolType, b.ko.Spec.ProtocolType) {
		delta.Add("Spec.ProtocolType", a.ko.Spec.ProtocolType, b.ko.Spec.ProtocolType)
	} else if a.ko.Spec.ProtocolType != nil && b.ko.Spec.ProtocolType != nil {
		if *a.ko.Spec.ProtocolType != *b.ko.Spec.ProtocolType {
			delta.Add("Spec.ProtocolType", a.ko.Spec.ProtocolType, b.ko.Spec.ProtocolType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RouteKey, b.ko.Spec.RouteKey) {
		delta.Add("Spec.RouteKey", a.ko.Spec.RouteKey, b.ko.Spec.RouteKey)
	} else if a.ko.Spec.RouteKey != nil && b.ko.Spec.RouteKey != nil {
		if *a.ko.Spec.RouteKey != *b.ko.Spec.RouteKey {
			delta.Add("Spec.RouteKey", a.ko.Spec.RouteKey, b.ko.Spec.RouteKey)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RouteSelectionExpression, b.ko.Spec.RouteSelectionExpression) {
		delta.Add("Spec.RouteSelectionExpression", a.ko.Spec.RouteSelectionExpression, b.ko.Spec.RouteSelectionExpression)
	} else if a.ko.Spec.RouteSelectionExpression != nil && b.ko.Spec.RouteSelectionExpression != nil {
		if *a.ko.Spec.RouteSelectionExpression != *b.ko.Spec.RouteSelectionExpression {
			delta.Add("Spec.RouteSelectionExpression", a.ko.Spec.RouteSelectionExpression, b.ko.Spec.RouteSelectionExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	} else if a.ko.Spec.Tags != nil && b.ko.Spec.Tags != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
			delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Target, b.ko.Spec.Target) {
		delta.Add("Spec.Target", a.ko.Spec.Target, b.ko.Spec.Target)
	} else if a.ko.Spec.Target != nil && b.ko.Spec.Target != nil {
		if *a.ko.Spec.Target != *b.ko.Spec.Target {
			delta.Add("Spec.Target", a.ko.Spec.Target, b.ko.Spec.Target)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Version, b.ko.Spec.Version) {
		delta.Add("Spec.Version", a.ko.Spec.Version, b.ko.Spec.Version)
	} else if a.ko.Spec.Version != nil && b.ko.Spec.Version != nil {
		if *a.ko.Spec.Version != *b.ko.Spec.Version {
			delta.Add("Spec.Version", a.ko.Spec.Version, b.ko.Spec.Version)
		}
	}

	return delta
}