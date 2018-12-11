// Copyright (C) 2018 Andrew Colin Kissa <andrew@datopdog.io>
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package spoolfile Reads and parses Exim spool files
*/
package spoolfile

import (
	"bytes"
	"fmt"
	"go/build"
	"os"
	"path"
	"testing"
)

var (
	gopath string
)

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
}

func TestBasics(t *testing.T) {
	var p, id, hf, df, did, hid, user, sender string
	var uid, gid int
	var err error
	var msg *Msg
	uid = 93
	gid = 93
	user = "exim"
	sender = "andrew@kudusoft.home.topdog-software.com"
	id = "1eXn2s-0008DG-EX"
	p = path.Join(gopath, "src/github.com/baruwa-enterprise/go-exim/testdata")
	hid = fmt.Sprintf("%s-H", id)
	did = fmt.Sprintf("%s-D", id)
	hf = path.Join(p, hid)
	df = path.Join(p, did)
	if msg, err = NewMsg(p, id); err != nil {
		t.Fatalf("UnExpected error: %s", err)
	}
	defer msg.Close()

	if !bytes.Equal(msg.ID, []byte(id)) {
		t.Errorf("Got %q want %q", msg.ID, id)
	}

	if !bytes.Equal(msg.User, []byte(user)) {
		t.Errorf("Got %q want %q", msg.User, user)
	}

	if msg.UID != uid {
		t.Errorf("Got %d want %d", msg.UID, uid)
	}

	if msg.GID != gid {
		t.Errorf("Got %d want %d", msg.GID, gid)
	}

	if !bytes.Equal(msg.Sender, []byte(sender)) {
		t.Errorf("Got %q want %q", msg.Sender, sender)
	}

	if msg.HdrFile != hf {
		t.Errorf("Got %q want %q", msg.HdrFile, hf)
	}

	if msg.DtaFile != df {
		t.Errorf("Got %q want %q", msg.DtaFile, df)
	}
}
