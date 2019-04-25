/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfig

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	gc "github.com/untillpro/gochips"
)

var ctx context.Context

// TestImpl s.e.
// Storage must be empty before testing
func TestImpl(actx context.Context, t *testing.T) {

	ctx = actx

	t.Run("TestBasicUsage", testBasicUsage)
}

func testBasicUsage(t *testing.T) {
	
}
