// Copyright 2016 goweb Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memcache

import (
	_ "github.com/bradfitz/gomemcache/memcache"

	"strconv"
	"testing"
	"time"

	"github.com/cooleo/goweb/cache"
)

func TestMemcacheCache(t *testing.T) {
	bm, err := cache.NewCache("memcache", `{"conn": "127.0.0.1:11211"}`)
	if err != nil {
		t.Error("init err")
	}
	timeoutDuration := 10 * time.Second
	if err = bm.Put("cooleo", "1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("cooleo") {
		t.Error("check err")
	}

	time.Sleep(11 * time.Second)

	if bm.IsExist("cooleo") {
		t.Error("check err")
	}
	if err = bm.Put("cooleo", "1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("cooleo").(string)); err != nil || v != 1 {
		t.Error("get err")
	}

	if err = bm.Incr("cooleo"); err != nil {
		t.Error("Incr Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("cooleo").(string)); err != nil || v != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("cooleo"); err != nil {
		t.Error("Decr Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("cooleo").(string)); err != nil || v != 1 {
		t.Error("get err")
	}
	bm.Delete("cooleo")
	if bm.IsExist("cooleo") {
		t.Error("delete err")
	}

	//test string
	if err = bm.Put("cooleo", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("cooleo") {
		t.Error("check err")
	}

	if v := bm.Get("cooleo").(string); v != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err = bm.Put("cooleo1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("cooleo1") {
		t.Error("check err")
	}

	vv := bm.GetMulti([]string{"cooleo", "cooleo1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if vv[0].(string) != "author" && vv[0].(string) != "author1" {
		t.Error("GetMulti ERROR")
	}
	if vv[1].(string) != "author1" && vv[1].(string) != "author" {
		t.Error("GetMulti ERROR")
	}

	// test clear all
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}
}
