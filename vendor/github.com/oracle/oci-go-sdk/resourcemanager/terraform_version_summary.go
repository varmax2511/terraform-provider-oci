// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TerraformVersionSummary A Terraform version supported for use with stacks.
type TerraformVersionSummary struct {

	// A supported Terraform version. Example: `0.12.x`
	Name *string `mandatory:"false" json:"name"`
}

func (m TerraformVersionSummary) String() string {
	return common.PointerString(m)
}
