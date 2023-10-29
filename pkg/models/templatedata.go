package models

// TemplateData holds data send from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // when you are not sure about the type you use interface
	CSRFToken string
	Flash     string
	Waringin  string
	Error     string
}
