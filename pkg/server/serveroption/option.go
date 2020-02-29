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

package serveroption

import (
	"context"

	pb "github.com/cilium/hubble/api/v1/flow"
)

// returning `stop: true` from a (pre)processor stops the execution chain, even
// if there was no error encountered (for example, explicitly filtering out
// certain events, or similar).
type stop = bool

// Options ...
type Options struct {
	Processors    []Processor
	Preprocessors []Preprocessor
}

// Option ...
type Option func(o *Options) error

// Preprocessor gets to act on the monitor payload before it's inserted into
// ths ring buffer.
type Preprocessor interface {
	Preprocess(context.Context, *pb.Payload) (error, stop)
}

// PreprocessorFunc defines a stateless preprocessor.
type PreprocessorFunc func(context.Context, *pb.Payload) (error, stop)

func (p PreprocessorFunc) Preprocess(
	ctx context.Context, pl *pb.Payload,
) (error, stop) {
	return p(ctx, pl)
}

// WithPreprocessor ...
func WithPreprocessor(pp Preprocessor) Option {
	return func(o *Options) error {
		o.Preprocessors = append(o.Preprocessors, pp)
		return nil
	}
}

// WithPreprocessorFunc ...
func WithPreprocessorFunc(
	f func(ctx context.Context, pl *pb.Payload) (error, stop),
) Option {
	return WithPreprocessor(PreprocessorFunc(f))
}

// Processor gets to act on the flow after it has been decoded from the payload
// and before insertion into the ring buffer.
type Processor interface {
	Process(context.Context, *pb.Flow) (error, stop)
}

// ProcessorFunc defines a stateless processor.
type ProcessorFunc func(context.Context, *pb.Flow) (error, stop)

func (p ProcessorFunc) Process(ctx context.Context, f *pb.Flow) (error, stop) {
	return p(ctx, f)
}

// WithProcessor ...
func WithProcessor(p Processor) Option {
	return func(o *Options) error {
		o.Processors = append(o.Processors, p)
		return nil
	}
}

// WithProcessorFunc ...
func WithProcessorFunc(
	f func(context.Context, *pb.Flow) (error, stop),
) Option {
	return WithProcessor(ProcessorFunc(f))
}
