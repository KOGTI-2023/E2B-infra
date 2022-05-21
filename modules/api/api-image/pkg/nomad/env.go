package nomad

import (
	"bytes"
	"fmt"
	"path"
  "strings"
	"text/template"

	"github.com/hashicorp/nomad/api"
)

const (
  templatesDir = "templates"
)

func escapeNewLines(input string) string {
  // HCL doesn't allow newlines in strings. We have to escape them.
  return strings.Replace(input, "\n", "\\\\n", -1)
}


//func (n *Nomad) CreateEnvironment(codeSnippetID, dockerfile string) (*api.JobDispatchResponse, *api.WriteMeta, error) {
func (n *Nomad) RegisterFCEnvJob(codeSnippetID, runtime string, deps []string) (string, error) {
	dockerfileName := fmt.Sprintf("%s.Dockerfile", runtime)
	tname := path.Join(templatesDir, "runtimes", dockerfileName)
	dockerfileTemp, err := template.ParseFiles(tname)
	if err != nil {
		return "", fmt.Errorf("Failed to parse template file '%s': %s", tname, err)
	}

	dockerfileTemp = template.Must(dockerfileTemp, err)

	dockerfileVars := struct {
		Deps []string
	}{
		Deps: deps,
	}
	var dockerfile bytes.Buffer
	if err := dockerfileTemp.Execute(&dockerfile, dockerfileVars); err != nil {
		return "", fmt.Errorf("Failed to `dockerfileTemp.Execute()`: %s", err)
	}

	tname = path.Join(templatesDir, "firecracker-envs.hcl")
	envsJobTemp, err := template.New("firecracker-envs.hcl").Funcs(
		template.FuncMap{
			"escapeNewLines": escapeNewLines,
		},
	).ParseFiles(tname)
	if err != nil {
		return "", fmt.Errorf("Failed to parse template file '%s': %s", tname, err)
	}

	envsJobTemp = template.Must(envsJobTemp, err)

	jobVars := struct {
		CodeSnippetID string
		Dockerfile    string
	}{
		CodeSnippetID: codeSnippetID,
		Dockerfile:    dockerfile.String(),
	}
	var jobDef bytes.Buffer
	if err := envsJobTemp.Execute(&jobDef, jobVars); err != nil {
		return "", fmt.Errorf("Failed to `envsJobTemp.Execute()`: %s", err)
	}

	job, err := n.nomadClient.Jobs().ParseHCL(jobDef.String(), false)
	if err != nil {
		return "", fmt.Errorf("Failed to parse the `firecracker-envs` HCL job file: %s", err)
	}

	res, _, err := n.nomadClient.Jobs().Register(job, &api.WriteOptions{})
	if err != nil {
		return "", fmt.Errorf("Failed to register 'firecracker-envs/%s' job: %s", jobVars.CodeSnippetID, err)
	}

	return res.EvalID, nil
}
