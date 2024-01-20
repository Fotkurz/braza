// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package formatters_test

import (
	"testing"

	engine "github.com/Fotkurz/braza/internal/engine"
	"github.com/Fotkurz/braza/pkg/entities/analysis"
	"github.com/stretchr/testify/assert"

	"github.com/Fotkurz/braza/internal/services/formatters"
	"github.com/Fotkurz/braza/internal/services/formatters/csharp/horuseccsharp"
	"github.com/Fotkurz/braza/internal/services/formatters/dart/horusecdart"
	"github.com/Fotkurz/braza/internal/services/formatters/java/horusecjava"
	"github.com/Fotkurz/braza/internal/services/formatters/javascript/horusecjavascript"
	"github.com/Fotkurz/braza/internal/services/formatters/kotlin/horuseckotlin"
	"github.com/Fotkurz/braza/internal/services/formatters/leaks/horusecleaks"
	"github.com/Fotkurz/braza/internal/services/formatters/nginx/horusecnginx"
	"github.com/Fotkurz/braza/internal/services/formatters/swift/horusecswift"
	"github.com/Fotkurz/braza/internal/services/formatters/yaml/horuseckubernetes"
	"github.com/Fotkurz/braza/internal/utils/testutil"
)

func TestStartAnalysis(t *testing.T) {
	testcases := []struct {
		engine    string
		formatter func(formatters.IService) formatters.IFormatter
	}{
		{
			engine:    "Csharp",
			formatter: horuseccsharp.NewFormatter,
		},
		{
			engine:    "Dart",
			formatter: horusecdart.NewFormatter,
		},
		{
			engine:    "Java",
			formatter: horusecjava.NewFormatter,
		},
		{
			engine:    "Javascript",
			formatter: horusecjavascript.NewFormatter,
		},
		{
			engine:    "Kotlin",
			formatter: horuseckotlin.NewFormatter,
		},
		{
			engine:    "Leaks",
			formatter: horusecleaks.NewFormatter,
		},
		{
			engine:    "Nginx",
			formatter: horusecnginx.NewFormatter,
		},
		{
			engine:    "Kubernetes",
			formatter: horuseckubernetes.NewFormatter,
		},
		{
			engine:    "Swift",
			formatter: horusecswift.NewFormatter,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.engine, func(t *testing.T) {
			t.Run("should success execute analysis without errors", func(t *testing.T) {
				analysis := &analysis.Analysis{}
				service := testutil.NewFormatterMock()

				service.On("LogDebugWithReplace")
				service.On("SetToolFinishedAnalysis")
				service.On("SetAnalysisError")
				service.On("ToolIsToIgnore").Return(false)
				service.On("GetConfigProjectPath").Return(".")
				service.On("ParseFindingsToVulnerabilities").Return(nil)
				service.On("GetCustomRulesByLanguage").Return([]engine.Rule{})

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})

				assert.Empty(t, len(analysis.Errors))
			})

			t.Run("should return error when getting text unit", func(t *testing.T) {
				service := testutil.NewFormatterMock()

				service.On("LogDebugWithReplace")
				service.On("SetToolFinishedAnalysis")
				service.On("SetAnalysisError")
				service.On("ToolIsToIgnore").Return(false)
				service.On("GetConfigProjectPath").Return(".")
				service.On("ParseFindingsToVulnerabilities").Return(nil)
				service.On("GetCustomRulesByLanguage").Return([]engine.Rule{})

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})
			})

			t.Run("should ignore this tool", func(t *testing.T) {
				service := testutil.NewFormatterMock()

				service.On("ToolIsToIgnore").Return(true)

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})
			})
		})
	}
}
