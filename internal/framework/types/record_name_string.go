// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*recordNameStringType)(nil)
)

type recordNameStringType struct {
	basetypes.StringType
}

var (
	RecordNameStringType = recordNameStringType{}
)

func (t recordNameStringType) Equal(o attr.Type) bool {
	other, ok := o.(recordNameStringType)
	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (recordNameStringType) String() string {
	return "RecordNameStringType"
}

func (t recordNameStringType) ValueFromString(_ context.Context, in types.String) (basetypes.StringValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return RecordNameStringNull(), diags
	}
	if in.IsUnknown() {
		return RecordNameStringUnknown(), diags
	}

	return RecordNameStringValue(in.ValueString()), diags
}

func (t recordNameStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (recordNameStringType) ValueType(context.Context) attr.Value {
	return RecordNameString{}
}

var (
	_ basetypes.StringValuable                   = (*RecordNameString)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*RecordNameString)(nil)
)

type RecordNameString struct {
	basetypes.StringValue
}

func RecordNameStringNull() RecordNameString {
	return RecordNameString{StringValue: basetypes.NewStringNull()}
}

func RecordNameStringUnknown() RecordNameString {
	return RecordNameString{StringValue: basetypes.NewStringUnknown()}
}

func RecordNameStringValue(value string) RecordNameString {
	return RecordNameString{StringValue: basetypes.NewStringValue(value)}
}

func (v RecordNameString) Equal(o attr.Value) bool {
	other, ok := o.(RecordNameString)
	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (RecordNameString) Type(context.Context) attr.Type {
	return RecordNameStringType
}

func (v RecordNameString) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(RecordNameString)
	if !ok {
		return false, diags
	}

	old, d := v.ToStringValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	new, d := newValue.ToStringValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	return RecordNamesSemanticallyEqual(old.ValueString(), new.ValueString()), diags
}

// RecordNamesSemanticallyEqual determines whether two record names are semantically equal
//
// This is done by:
// - Remove trailing periods (if present) from both values.
// - Applying a case insensitive comparison.
func RecordNamesSemanticallyEqual(s1, s2 string) bool {
	return strings.EqualFold(strings.TrimSuffix(s1, "."), strings.TrimSuffix(s2, "."))
}
