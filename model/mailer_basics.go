package model

type MailerBasics struct {
	Language			string			`json:"language"`
	Subject				string			`json:"subject"`
	TextHtml			string			`json:"text_html"`
	TextPlain			string			`json:"text_plain"`
	ToList				[]string		`json:"to_list"`
}

type MapMailToLang struct {
	Kaz					string
	Rus					string
	Default					string
}

type MapMailNotification struct {
	SubjectLangMap					MapMailToLang
	TextHtmlLangMap					MapMailToLang
	TextPageLangMap					MapMailToLang
}
