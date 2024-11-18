// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"k8s.io/utils/ptr"
)

// bool
func BoolValueToBool(input basetypes.BoolValue) bool {
	return input.ValueBool()
}

func BoolValueToPtrBool(input basetypes.BoolValue) *bool {
	return ptr.To(input.ValueBool())
}

func BoolToBoolValue(input bool) basetypes.BoolValue {
	return basetypes.NewBoolValue(input)
}

func PtrBoolToBoolValue(input *bool) basetypes.BoolValue {
	if input == nil {
		return basetypes.NewBoolValue(false)
	}
	return basetypes.NewBoolValue(*input)
}

// string
func StringValueToString(input basetypes.StringValue) string {
	return input.ValueString()
}

func StringToStringValue(input string) basetypes.StringValue {
	return basetypes.NewStringValue(input)
}

func ListValueStringToSliceString(input basetypes.ListValue) ([]string, error) {
	out := []string{}
	diag := input.ElementsAs(context.Background(), &out, true)
	if diag.HasError() {
		return nil, errors.New("could not convert listvalue[string] to string slice")
	}
	return out, nil
}

func SliceStringToListValueString(input []string) (basetypes.ListValue, error) {
	out, diag := basetypes.NewListValueFrom(context.TODO(), types.StringType, input)
	if diag.HasError() {
		return out, errors.New("could not convert string slice to listvalue[string]")
	}
	return out, nil
}

func ListValueStringToPtrSliceString(input basetypes.ListValue) (*[]string, error) {
	out := []string{}
	diag := input.ElementsAs(context.Background(), &out, true)
	if diag.HasError() {
		return nil, errors.New("could not convert listvalue[string] to string slice")
	}
	return &out, nil
}

func PtrSliceStringToListValueString(input *[]string) (basetypes.ListValue, error) {
	out, diag := basetypes.NewListValueFrom(context.TODO(), types.StringType, input)
	if diag.HasError() {
		return out, errors.New("could not convert string slice to listvalue[string]")
	}
	return out, nil
}

// time
func TimeToStringValue(input time.Time) basetypes.StringValue {
	return basetypes.NewStringValue(input.Format(time.RFC3339))
}

// uuid
func UUIDToStringValue(input uuid.UUID) basetypes.StringValue {
	return basetypes.NewStringValue(input.String())
}

func StringValueToUUID(input basetypes.StringValue) uuid.UUID {
	return uuid.MustParse(input.ValueString())
}

// int
func IntToInt64Value(input int) basetypes.Int64Value {
	return basetypes.NewInt64Value(int64(input))
}
