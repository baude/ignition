package applehv

// Copyright 2023 Red Hat
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

import (
	"context"
	"io"
	"net"
	"net/http"

	"github.com/coreos/ignition/v2/config/v3_5_experimental/types"
	"github.com/coreos/ignition/v2/internal/platform"
	"github.com/coreos/ignition/v2/internal/providers/util"
	"github.com/coreos/ignition/v2/internal/resource"
	"github.com/coreos/vcontext/report"
	"github.com/mdlayher/vsock"
)

func init() {
	platform.Register(platform.Provider{
		Name:  "applehv",
		Fetch: fetchConfig,
	})
}

func fetchConfig(f *resource.Fetcher) (types.Config, report.Report, error) {
	contextID, err := vsock.ContextID()
	if err != nil {
		return types.Config{}, report.Report{}, err
	}
	conn, err := vsock.Dial(contextID, 1024, &vsock.Config{})
	if err != nil {
		return types.Config{}, report.Report{}, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			f.Logger.Err("unable to close vsock connection: %v", err)
		}
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return conn, nil
			},
		},
	}

	r, err := client.Get("http://localhost")
	if err != nil {
		return types.Config{}, report.Report{}, err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			f.Logger.Err("unable to close response body: %v", err)
		}
	}()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return types.Config{}, report.Report{}, err
	}

	return util.ParseConfig(f.Logger, b)
}
