package utils

func FormatDocTextClass(currentRoute string, focusedSubject string) string {
	style := "docs_container"

	if currentRoute != focusedSubject {
		style += " remove-display"
	}

	return style
}
