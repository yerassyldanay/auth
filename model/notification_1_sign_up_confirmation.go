package model

import (
	"auth/utils/constants"
	"fmt"
)

type NotifySignUpConfirmation struct {
	MailerBasics
	Address					string
	PreparedLink			string
}

var MapNotifySignUpConfirmation = MapMailNotification{
	SubjectLangMap:  MapMailToLang{
		Kaz: "Аккаунты растауыңызды сұраймыз",
		Rus: "Потвердите, пожалуйста, аккаунт",
		Default: "Please, verify your account",
	},
	TextHtmlLangMap: MapMailToLang{
		Kaz:     "Осы сілтеме бойынша өтіңіз: %s \n",
		Rus:     "Пожалуйста, проходите по следующей ссылке: %s \n",
		Default: "Please, verify your account using following link: %s \n",
	},
	TextPageLangMap: MapMailToLang{
		Kaz:     "Осы сілтеме бойынша өтіңіз: %s \n",
		Rus:     "Пожалуйста, проходите по следующей ссылке: %s \n",
		Default: "Please, verify your account using following link: %s \n",
	},
}

func (nsuc *NotifySignUpConfirmation) GetFrom() string {
	return constants.BASIC_MAILER_FROM
}

func (nsuc *NotifySignUpConfirmation) GetToList() []string {
	return []string{nsuc.Address}
}

func (nsuc *NotifySignUpConfirmation) GetSubject() string {
	switch nsuc.Language {
	case constants.LANG_KAZ:
		return MapNotifySignUpConfirmation.SubjectLangMap.Kaz
	case constants.LANG_RUS:
		return MapNotifySignUpConfirmation.SubjectLangMap.Rus
	}
	return MapNotifySignUpConfirmation.SubjectLangMap.Default
}

func (nsuc *NotifySignUpConfirmation) GetHtml() string {
	if nsuc.TextHtml != "" {
		return nsuc.TextHtml
	}

	switch nsuc.Language {
	case constants.LANG_KAZ:
		nsuc.TextHtml = fmt.Sprintf(MapNotifySignUpConfirmation.TextHtmlLangMap.Kaz, nsuc.PreparedLink)
	case constants.LANG_RUS:
		nsuc.TextHtml = fmt.Sprintf(MapNotifySignUpConfirmation.TextHtmlLangMap.Rus, nsuc.PreparedLink)
	default:
		nsuc.TextHtml = fmt.Sprintf(MapNotifySignUpConfirmation.TextHtmlLangMap.Default, nsuc.PreparedLink)
	}

	return nsuc.TextHtml
}

func (nsuc *NotifySignUpConfirmation) GetPlainText() string {
	return nsuc.GetHtml()
}
