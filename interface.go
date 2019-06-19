/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfig

import "context"

// GetConfig fills passed config struct and returns ok == true, ok == false and err == nil if config doesn't
// exists and ok == false and err != nil if an error occurred
var GetConfig func(ctx context.Context, configName string, config interface{}) (ok bool, err error)

// PutConfig saves config struct with given name or returns error if something went wrong
var PutConfig func(ctx context.Context, configName string, config interface{}) error
