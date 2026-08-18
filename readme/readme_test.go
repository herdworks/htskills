package readme

import (
	"bytes"
	"strings"
	"testing"
)

func TestREADMEExecutesWithCommand(t *testing.T) {
	template, err := README()
	if err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err = template.Execute(&output, map[string]string{"Command": "htskills"}); err != nil {
		t.Fatal(err)
	}
	for _, phrase := range []string{`"htskills up"`, "every discovered project-local skill root", "Use project-local HyperTexting Skills only; remove globally installed copies."} {
		if !strings.Contains(output.String(), phrase) {
			t.Fatalf("README output is missing %q: %q", phrase, output.String())
		}
	}
}
