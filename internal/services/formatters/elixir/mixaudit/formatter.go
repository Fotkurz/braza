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

package mixaudit

import (
	"encoding/json"

	"github.com/Fotkurz/braza/pkg/entities/vulnerability"
	"github.com/Fotkurz/braza/pkg/enums/languages"
	"github.com/Fotkurz/braza/pkg/enums/severities"
	"github.com/Fotkurz/braza/pkg/enums/tools"
	"github.com/Fotkurz/braza/pkg/utils/logger"

	"github.com/Fotkurz/braza/internal/entities/docker"
	"github.com/Fotkurz/braza/internal/enums/images"
	"github.com/Fotkurz/braza/internal/helpers/messages"
	"github.com/Fotkurz/braza/internal/services/formatters"
	vulnhash "github.com/Fotkurz/braza/internal/utils/vuln_hash"
)

type Formatter struct {
	formatters.IService
}

func NewFormatter(service formatters.IService) *Formatter {
	return &Formatter{
		service,
	}
}

func (f *Formatter) StartAnalysis(projectSubPath string) {
	if f.ToolIsToIgnore(tools.MixAudit) || f.IsDockerDisabled() {
		logger.LogDebugWithLevel(messages.MsgDebugToolIgnored + tools.MixAudit.ToString())
		return
	}

	output, err := f.startMixAudit(projectSubPath)
	f.SetAnalysisError(err, tools.MixAudit, output, projectSubPath)
	f.LogDebugWithReplace(messages.MsgDebugToolFinishAnalysis, tools.MixAudit, languages.Elixir)
}

func (f *Formatter) startMixAudit(projectSubPath string) (string, error) {
	f.LogDebugWithReplace(messages.MsgDebugToolStartAnalysis, tools.MixAudit, languages.Elixir)

	output, err := f.ExecuteContainer(f.getConfigData(projectSubPath))
	if err != nil || output == "" {
		return output, err
	}

	return output, f.parseOutput(output)
}

func (f *Formatter) getConfigData(projectSubPath string) *docker.AnalysisData {
	analysisData := &docker.AnalysisData{
		CMD:      f.GetConfigCMDByFileExtension(projectSubPath, CMD, "mix.lock", tools.MixAudit),
		Language: languages.Elixir,
	}

	return analysisData.SetImage(f.GetCustomImageByLanguage(languages.Elixir), images.Elixir)
}

func (f *Formatter) parseOutput(output string) error {
	var result mixAuditResult

	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return err
	}

	for index := range result.Vulnerabilities {
		f.AddNewVulnerabilityIntoAnalysis(f.newVulnerability(&result.Vulnerabilities[index]))
	}

	return nil
}

func (f *Formatter) newVulnerability(mixVuln *mixAuditVulnerability) *vulnerability.Vulnerability {
	vuln := &vulnerability.Vulnerability{
		SecurityTool: tools.MixAudit,
		RuleID:       mixVuln.Advisory.CVE,
		Language:     languages.Elixir,
		Severity:     severities.High,
		Details:      mixVuln.getDetails(),
		Code:         f.GetCodeWithMaxCharacters(mixVuln.Advisory.Package, 0),
		File:         f.RemoveSrcFolderFromPath(mixVuln.Dependency.Lockfile),
	}
	return f.SetCommitAuthor(vulnhash.Bind(vuln))
}
