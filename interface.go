/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfig

import "context"

// GetConfig returns config with given name
// Config never will be nil
// If config does not exist empty map is returned
var GetConfig func(ctx context.Context, configName string, config interface{}) error

// PutConfig saves config with given name
var PutConfig func(ctx context.Context, configName string, config interface{}) error
