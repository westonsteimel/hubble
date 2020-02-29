// Copyright 2020 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"context"

	pb "github.com/cilium/hubble/api/v1/flow"
	"github.com/cilium/hubble/pkg/metrics"
)

// ProcessMetrics ...
type ProcessMetrics struct{}

func (m *ProcessMetrics) Process(ctx context.Context, f *pb.Flow) (error, bool) {
	// TODO: refactor metrics to no longer be global; maybe also add
	// `WithoutMetrics() Option` to turn off the default.
	metrics.ProcessFlow(f)

	return nil, false
}
