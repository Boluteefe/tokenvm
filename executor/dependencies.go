// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

type Metrics interface {
	RecordBlocked()
	RecordExecutable()
}
