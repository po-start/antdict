package entity

import (
//	"encoding/json"
)

type Translation struct {
	Pos				string							`json:"pos"`
	V				string							`json:"v"`
}

type PronunciationEntity struct {
	Symbol			string							`json:"symbol"`
	Url				string							`json:"url"`
}

type Pronunciation struct {
	BrE				PronunciationEntity				`json:"BrE"`
	AmE				PronunciationEntity				`json:"AmE"`
}

type ChartSensesEntity struct {
	Percent			int64							`json:"percent"`
	Sense			string							`json:"sense"`
}

type SVEntity struct {
	Text			string							`json:"text"`
	Trans			string							`json:"trans"`
}

type SentenceEntity struct {
	Pos				string							`json:"pos"`
	V				[]SVEntity						`json:"v"`
}

type MoreDetailsEntity struct {
	Phrase			string							`json:"phrase"`
	Notes			[]string						`json:"notes"`
}

type DetailsEntity struct {
	Explanation		string							`json:"explanation"`
	MoreDetails		[]MoreDetailsEntity				`json:"moreDetails"`
}

type OVEntity struct {
	Phrase			string							`json:"phrase"`
	Details			DetailsEntity					`json:"details"`
}

type OphraseEntity struct {
	Pos				string							`json:"pos"`
	C				[]OVEntity						`json:"v"`
}

type LanguageEntity struct {
	En				string							`json:"en"`
	Cn				string							`json:"cn"`
}

type VVEntity struct {
	Explanation		string							`json:"explanation"`
	List			[]LanguageEntity				`json:"list"`
}

type VocabularyEntity struct {
	Pos				string							`json:"pos"`
	V				[]VVEntity						`json:"v"`
}

type WVEntity struct {
	Items			string							`json:"items"`
	Notes			[]string						`json:"notes"`
}

type WordDiffEntity struct {
	Pos				string							`json:"pos"`
	V				[]WVEntity						`json:"v"`
}

type SynonymEntity struct {
	V				[]LanguageEntity				`json:"v"`
}

type AntonymEntity struct {
	V				[]LanguageEntity				`json:"v"`
}

type SenEntity struct {
	Pic				string							`json:"pic"`
	En				map[string]string				`json:"en"`
	Cn				map[string]string				`json:"cn"`
}

type KtEntity struct {
	Text			string							`json:"text"`
	Pic				string							`json:"pic"`
}

type BodyEntity struct {
	Pos				string							`json:"pos"`
	Sense			string							`json:"sense"`
	Audio			map[string]string				`json:"audio"`
	Phrase			map[string]LanguageEntity		`json:"phrase"`
	Sen				SenEntity						`json:"sen"`
	Kt				map[string]KtEntity				`json:"kt"`
}

type XXEntity struct {
	Item			string							`json:"item"`
	Body			map[string]BodyEntity			`json:"body"`
}

type RateEntity struct {
	Cloze			string							`json:"cloze"`
	Reading			string							`json:"reading"`
	Translation		string							`json:"translation"`
	Danxuan			string							`json:"danxuan"`
	Title			string							`json:"title"`
}

type TransFormEntity struct {
	Pos				string							`json:"pos"`
	Text			string							`json:"text"`
}

type ZhentiEntity struct {
	En				string							`json:"en"`
	Src				string							`json:"src"`
}

type BasicEntity struct {
	Pos				string							`json:"pos"`
	Sense			string							`json:"sense"`
	Typical			map[string]LanguageEntity		`json:"typical"`
	Zhenti			map[string]ZhentiEntity			`json:"zhenti"`
}

type ChuzhongEntity struct {
	Item			string							`json:"item"`
	Rate			RateEntity						`json:"rate"`
	Transform		map[string]TransFormEntity		`json:"transform"`
	Basic			map[string]BasicEntity			`json:"basic"`
	Ss				[]string						`json:"ss"`
}

type Dict struct {
    Item			string							`json:"item"`
    Translation		[]Translation					`json:"translation"`
	Pronunciation	Pronunciation					`json:"pronunciation"`
	ChartSenses		map[string]ChartSensesEntity	`json:"chart_senses"`
	Sentence		[]SentenceEntity				`json:"sentence"`
	Ophrase			[]OphraseEntity					`json:"ophrase"`
	Vocabulary		[]VocabularyEntity				`json:"vocabulary"`
	WordDiff		[]WordDiffEntity				`json:"word_diff"`
	Synonym			[]SynonymEntity					`json:"synonym"`
	Antonym			[]AntonymEntity					`json:"antonym"`
	Xx				XXEntity						`json:"xx"`
	Chuzhong		ChuzhongEntity					`json:"chuzhong"`
}
