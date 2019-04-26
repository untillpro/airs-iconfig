/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfig

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var ctx context.Context

// TestImpl s.e.
// Storage must be empty before testing
func TestImpl(actx context.Context, t *testing.T) {
	ctx = actx
	t.Run("TestBasicUsage", testBasicUsage)
}

type testConfig struct {
	Param1 string
	Param2 int
	Param3 bool
	Param4 []string
	Param5 map[string]float64
}

type minTestConfig struct {
	Param1 string
	Param2 int
	Param4 []string
	Param5 map[string]float64
}

type maxTestConfig struct {
	Param1 string
	Param2 int
	Param3 bool
	Param4 []string
	Param5 map[string]float64
	Param6 map[string]interface{}
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func testBasicUsage(t *testing.T) {
	require.NotNil(t, ctx, "Need to provide not nil context.Context to TestImpl(context.Context, *ttesting.T)")
	TestPutGet(t)
	TestNilConfig(t)
	TestNotPointerInGet(t)
	TestGetWrongStruct(t)
	TestPutGetDifferentStructs(t)
}

var testConfig1 = testConfig{"ac", 3, true, []string{"assert", "b", "c"},
	map[string]float64{"assert": 1.1, "b": 2.2}}

func TestPutGet(t *testing.T) {
	prefix := randStringBytes(8)
	err := PutConfig(ctx, prefix, &testConfig1)
	require.Nil(t, err, "Can't put test config to KV! Config: ", err)
	var b testConfig
	err = GetConfig(ctx, prefix, &b)
	require.Nil(t, err, "Can't get test config from KV! Config: ", err)
	require.True(t, cmp.Equal(testConfig1, b), "Structs must be equal! ", testConfig1, b)
	require.False(t, cmp.Equal(&ctx, &b))
}

func TestNilConfig(t *testing.T) {
	var config *testConfig = nil
	err := PutConfig(ctx, "", config)
	require.NotNil(t, err)
}

func TestNotPointerInGet(t *testing.T) {
	var b testConfig
	err := GetConfig(ctx, "", b)
	require.NotNil(t, err)
}

func TestGetWrongStruct(t *testing.T) {
	prefix := randStringBytes(8)
	err := PutConfig(ctx, prefix, &testConfig1)
	require.Nil(t, err, "Can't put test config to KV! Config: ", err)

	//try to unmarshal config to wrong struct
	var b error
	err = GetConfig(ctx, prefix, &b)
	require.Nil(t, b)
	require.NotNil(t, err)
}

func TestPutGetDifferentStructs(t *testing.T) {
	prefix := randStringBytes(8)
	err := PutConfig(ctx, prefix, &testConfig1)
	require.Nil(t, err, "Can't put test config to KV! Config: ", err)

	var b minTestConfig

	err = GetConfig(ctx, prefix, &b)
	require.Nil(t, err, "Can't get test config from KV! Config: ", err)
	require.True(t, !cmp.Equal(testConfig1, b), "Structs must be unequal! ", testConfig1, b)

	//all presented values in minTestConfig are equal
	require.Equal(t, testConfig1.Param1, b.Param1)
	require.Equal(t, testConfig1.Param2, b.Param2)
	require.Equal(t, testConfig1.Param4, b.Param4)
	require.Equal(t, testConfig1.Param5, b.Param5)

	var c maxTestConfig

	err = GetConfig(ctx, prefix, &c)
	require.Nil(t, err, "Can't get test config from KV! Config: ", err)
	require.True(t, !cmp.Equal(testConfig1, c), "Structs must be unequal! ", testConfig1, b)

	//all presented values in minTestConfig are equal
	require.Equal(t, testConfig1.Param1, c.Param1)
	require.Equal(t, testConfig1.Param2, c.Param2)
	require.Equal(t, testConfig1.Param4, c.Param4)
	require.Equal(t, testConfig1.Param5, c.Param5)
	require.True(t, len(c.Param6) == 0)
}
