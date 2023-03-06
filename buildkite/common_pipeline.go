package buildkite

import "fmt"

func TrivyStepGenerator(trivyPlugin, shellPlugin string, w io.Writer) error {
	trivyStepFormat := `
steps:
- command: ls
  plugins:
    - equinixmetal-buildkite/trivy#%s
- label: ":sparkles: SHELL CHECK"
  plugins:
    - shellcheck#%s:
        files: script.sh
`

	trivyStep := fmt.Sprintf(trivyStepFormat, trivyPlugin, shellPlugin)
	fmt.Println(trivyStep)
}
