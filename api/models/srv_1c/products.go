package models

type Response struct {
	Metadata string      `json:"odata.metadata"`
	Product  []Product1c `json:"value"`
}

type Product1c struct {
	UUID        string `json:"Ref_Key" gorm:"primaryKey"`
	Title       string `json:"НаименованиеПолное"`
	Code        string `json:"Code"`
	Unit        string `json:"БазоваяЕдиницаИзмерения@navigationLinkUrl"`
	Description string `json:"Description"`
}
