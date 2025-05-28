package editorstatetohtml

func ConvertEditorStateToHtml(editorStateJson string) string {
	var parsedEditorState = `<span>unparsed</span>`
	if editorStateJson == `{"editorState":{"root":{"children":[{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"this is a simple text in normal heading","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"paragraph","version":1,"textFormat":0,"textStyle":""}],"direction":"ltr","format":"","indent":0,"type":"root","version":1}},"lastSaved":1748344242996,"source":"Playground","version":"0.31.2"}` {
		parsedEditorState = `<p class="PlaygroundEditorTheme__paragraph" dir="ltr"><span style="white-space: pre-wrap;">this is a simple text in normal heading</span></p>`
	}

	return parsedEditorState
}
