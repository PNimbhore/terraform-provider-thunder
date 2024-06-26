package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ValidateDataResourceConfig_Request(in *tfprotov6.ValidateDataResourceConfigRequest) (*tfplugin6.ValidateDataResourceConfig_Request, error) {
	resp := &tfplugin6.ValidateDataResourceConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func ValidateDataResourceConfig_Response(in *tfprotov6.ValidateDataResourceConfigResponse) (*tfplugin6.ValidateDataResourceConfig_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfplugin6.ValidateDataResourceConfig_Response{
		Diagnostics: diags,
	}, nil
}

func ReadDataSource_Request(in *tfprotov6.ReadDataSourceRequest) (*tfplugin6.ReadDataSource_Request, error) {
	resp := &tfplugin6.ReadDataSource_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	if in.ProviderMeta != nil {
		resp.ProviderMeta = DynamicValue(in.ProviderMeta)
	}
	return resp, nil
}

func ReadDataSource_Response(in *tfprotov6.ReadDataSourceResponse) (*tfplugin6.ReadDataSource_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	resp := &tfplugin6.ReadDataSource_Response{
		Diagnostics: diags,
	}
	if in.State != nil {
		resp.State = DynamicValue(in.State)
	}
	return resp, nil
}

// we have to say this next thing to get golint to stop yelling at us about the
// underscores in the function names. We want the function names to match
// actually-generated code, so it feels like fair play. It's just a shame we
// lose golint for the entire file.
//
// This file is not actually generated. You can edit it. Ignore this next line.
// Code generated by hand ignore this next bit DO NOT EDIT.
