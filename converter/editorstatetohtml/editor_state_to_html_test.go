package editorstatetohtml

import "testing"

func TestConvertEditorStateToHtml(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// {"empty string", "", ""},
		// {"simple json", `{"key": "value"}`, `{"key": "value"}`},
		// {"lexical like structure", `{"root":{"children":[{"type":"paragraph","children":[{"type":"text","text":"Hello"}]}]}}`, `{"root":{"children":[{"type":"paragraph","children":[{"type":"text","text":"Hello"}]}]}}`},
		{"normal heading",
			`{"editorState":{"root":{"children":[{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"this is a simple text in normal heading","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"paragraph","version":1,"textFormat":0,"textStyle":""}],"direction":"ltr","format":"","indent":0,"type":"root","version":1}},"lastSaved":1748344242996,"source":"Playground","version":"0.31.2"}`,
			`<p class="PlaygroundEditorTheme__paragraph" dir="ltr"><span style="white-space: pre-wrap;">this is a simple text in normal heading</span></p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertEditorStateToHtml(tt.input); got != tt.want {
				t.Errorf("ConvertEditorStateToHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
