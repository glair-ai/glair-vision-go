package ocr

// GeneralDocument stores OCR result of all-purpose general model\
// from the given input
type GeneralDocument = OCRResult[GeneralDocumentData]

type GeneralDocumentData struct {
	AllTexts []OCRField `json:"all_texts,omitempty"`
}
