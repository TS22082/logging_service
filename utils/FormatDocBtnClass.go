package utils

func FormatDocBtnClass(currentRoute string, focusedSubject string) string {
	style := "docs_btn"

	if currentRoute == focusedSubject {
		style += " docs_btn_active"
	}

	return style
}
