// Code generated by "packer-sdc mapstructure-to-hcl2 "; DO NOT EDIT.

package googlecomputeexport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AccountFile               *string           `mapstructure:"account_file" cty:"account_file" hcl:"account_file"`
	ImpersonateServiceAccount *string           `mapstructure:"impersonate_service_account" required:"false" cty:"impersonate_service_account" hcl:"impersonate_service_account"`
	DiskSizeGb                *int64            `mapstructure:"disk_size" cty:"disk_size" hcl:"disk_size"`
	DiskType                  *string           `mapstructure:"disk_type" cty:"disk_type" hcl:"disk_type"`
	MachineType               *string           `mapstructure:"machine_type" cty:"machine_type" hcl:"machine_type"`
	Network                   *string           `mapstructure:"network" cty:"network" hcl:"network"`
	Paths                     []string          `mapstructure:"paths" required:"true" cty:"paths" hcl:"paths"`
	Subnetwork                *string           `mapstructure:"subnetwork" cty:"subnetwork" hcl:"subnetwork"`
	Zone                      *string           `mapstructure:"zone" cty:"zone" hcl:"zone"`
	VaultGCPOauthEngine       *string           `mapstructure:"vault_gcp_oauth_engine" cty:"vault_gcp_oauth_engine" hcl:"vault_gcp_oauth_engine"`
	ServiceAccountEmail       *string           `mapstructure:"service_account_email" cty:"service_account_email" hcl:"service_account_email"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":           &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":         &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":         &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":             &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":       &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":  &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"account_file":                &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"impersonate_service_account": &hcldec.AttrSpec{Name: "impersonate_service_account", Type: cty.String, Required: false},
		"disk_size":                   &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_type":                   &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"machine_type":                &hcldec.AttrSpec{Name: "machine_type", Type: cty.String, Required: false},
		"network":                     &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"paths":                       &hcldec.AttrSpec{Name: "paths", Type: cty.List(cty.String), Required: false},
		"subnetwork":                  &hcldec.AttrSpec{Name: "subnetwork", Type: cty.String, Required: false},
		"zone":                        &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"vault_gcp_oauth_engine":      &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
		"service_account_email":       &hcldec.AttrSpec{Name: "service_account_email", Type: cty.String, Required: false},
	}
	return s
}
