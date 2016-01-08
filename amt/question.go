package amt

import (
	"encoding/xml"
	"fmt"
	answerkey "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd_go"
	questionform "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd_go"
	questionformanswers "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd_go"
	externalquestion "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd_go"
	htmlquestion "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
	"net/url"
	"regexp"
	"strings"
)

var (
	reNamespace = regexp.MustCompile(`( xmlns=".*?")`)
)

// HITQuestion is used to initialize HIT questions.
type HITQuestion interface{}

// Encode a question to send to Amazon
func EncodeQuestion(question HITQuestion) ([]byte, error) {
	result, err := xml.Marshal(question)
	if err != nil {
		return nil, err
	}
	idx := reNamespace.FindIndex(result)
	if idx != nil {
		ns := string(result[idx[0]:idx[1]])
		nons := strings.Replace(string(result[idx[1]+1:]), ns, "", -1)
		result = []byte(string(result[0:idx[1]+1]) + nons)
	}
	return result, nil
}

// Decode a question in a response from Amazon
func DecodeQuestion(questionXml []byte) (HITQuestion, error) {

	// Choose a type based on the earliest type-identifying XML element
	var (
		minLoc   = len(questionXml)
		question HITQuestion
		xmls     = string(questionXml)
	)
	if idx := strings.Index(xmls, "ExternalQuestion"); idx >= 0 && idx < minLoc {
		minLoc = idx
		question = &ExternalQuestion{}
	}
	if idx := strings.Index(xmls, "HTMLQuestion"); idx >= 0 && idx < minLoc {
		minLoc = idx
		question = &HTMLQuestion{}
	}
	if idx := strings.Index(xmls, "QuestionForm"); idx >= 0 && idx < minLoc {
		minLoc = idx
		question = &QuestionForm{}
	}

	// Unmarshal the question
	err := xml.Unmarshal(questionXml, question)
	return question, err
}

// ExternalQuestion is a HITQuestion that uses the ExternalQuestion XML schema.
// These questions are hosted by your own server.
type ExternalQuestion struct {

	// The name of the wrapper element for an XML representation of the object
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd ExternalQuestion"`

	externalquestion.XsdGoPkgHasElem_ExternalURLsequenceTxsdExternalQuestionExternalQuestionschema_ExternalURL_XsdtAnyURI_
	externalquestion.XsdGoPkgHasElem_FrameHeightsequenceTxsdExternalQuestionExternalQuestionschema_FrameHeight_XsdtInteger_
}

// HTMLQuestion is a HITQuestion that uses the HTMLQuestion XML schema.
// These questions are hosted by Amazon using arbitrary HTML you provide.
type HTMLQuestion struct {

	// The name of the wrapper element for an XML representation of the object
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd HTMLQuestion"`

	// The question. Exposes the following top-level properties:
	// HTMLContent xsdt.String: The HTML for the question.
	// FrameHeight xsdt.Integer: The height of the HTML frame to hold the question.
	htmlquestion.TxsdHTMLQuestion
}

// QuestionForm is a HITQuestion that uses the QuestionForm XML schema.
// These questions are hosted by Amazon using a standard form interface.
//
// Helper methods are provided to construct the form. In particular, Overview
// and Question elements will be marshalled into the output XML in the order
// added, provided you use AddOverview(), AddQuestion(), and EncodeQuestion().
type QuestionForm struct {

	// The name of the wrapper element for an XML representation of the object
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd QuestionForm"`

	// The raw XML from which the form was parsed
	XMLContent string `xml:",innerxml"`

	// The question. Exposes the following top-level properties:
	// Overviews []*questionform.TContentType: Information to display above the questions.
	// Questions []*questionform.TxsdQuestionFormSequenceChoiceQuestion: The questions.
	questionform.TxsdQuestionForm

	// Corresponds to the addition order of overviews and questions
	addedOverviewNext []bool `xml:"-"`
}

// Add a new Overview item, to be populated by subsequent content-adding method
// calls.
func (question *QuestionForm) AddOverview() {
	question.Overviews = append(question.Overviews,
		&questionform.TContentType{})
	question.addedOverviewNext = append(question.addedOverviewNext, true)
}

// Add a new Question item, to be populated by subsequent content-adding method
// calls.
func (question *QuestionForm) AddQuestion(questionIdentifier,
	displayName string, isRequired bool) {

	qq := &questionform.TxsdQuestionFormSequenceChoiceQuestion{}
	qq.QuestionIdentifier = xsdt.String(questionIdentifier)
	qq.DisplayName = xsdt.String(displayName)
	qq.IsRequired = xsdt.Boolean(isRequired)

	question.Questions = append(question.Questions, qq)
	question.addedOverviewNext = append(question.addedOverviewNext, false)
}

// Retrieve the content struct for the most recently-added Overview or Question
func (question *QuestionForm) getCurrentContent() *questionform.TContentType {
	if question.addedOverviewNext[len(question.addedOverviewNext)-1] {
		return question.Overviews[len(question.Overviews)-1]
	} else {
		qq := question.Questions[len(question.Questions)-1]
		if qq.QuestionContent == nil {
			qq.QuestionContent = &questionform.TContentType{}
		}
		return qq.QuestionContent
	}
}

// Retrieve the answer specification struct for the most recently-added Question
func (question *QuestionForm) getCurrentAnswer() *questionform.TAnswerSpecificationType {
	qq := question.Questions[len(question.Questions)-1]
	if qq.AnswerSpecification == nil {
		qq.AnswerSpecification = &questionform.TAnswerSpecificationType{}
	}
	return qq.AnswerSpecification
}

// Add a List item to the most recent Question/Overview added.
func (question *QuestionForm) AddListContent(listItems []string) {
	list := &questionform.TxsdContentTypeChoiceList{}
	for _, listItem := range listItems {
		list.ListItems = append(list.ListItems, xsdt.String(listItem))
	}

	content := question.getCurrentContent()
	content.Lists = append(content.Lists, list)
}

// Add a FormattedContent item to the most recent Question/Overview added.
func (question *QuestionForm) AddFormattedContent(formattedContent string) {
	content := question.getCurrentContent()
	content.FormattedContents = append(content.FormattedContents, xsdt.String(formattedContent))
}

// Add a Binary item to the most recent Question/Overview added.
func (question *QuestionForm) AddBinaryContent(mimeType, mimeSubType string,
	dataURL *url.URL, altText string) {
	binary := &questionform.TBinaryContentType{}
	if mimeType != "" || mimeSubType != "" {
		binary.MimeType = &questionform.TMimeType{}
		if mimeType != "" {
			binary.MimeType.Type = questionform.TxsdMimeTypeSequenceType(mimeType)
		}
		if mimeSubType != "" {
			binary.MimeType.SubType = xsdt.String(mimeSubType)
		}
	}
	binary.DataURL = questionform.TURLType(dataURL.String())
	if altText != "" {
		binary.AltText = xsdt.String(altText)
	}

	content := question.getCurrentContent()
	content.Binaries = append(content.Binaries, binary)
}

// Add a Flash Application item to the most recent Question/Overview added.
func (question *QuestionForm) AddFlashApplicationContent(flashMovieURL *url.URL,
	width, height int, applicationParameters map[string]string) {
	application := &questionform.TApplicationContentType{}
	application.Flash = &questionform.TFlashContentType{}
	application.Flash.FlashMovieURL = questionform.TURLType(flashMovieURL.String())
	application.Flash.Width = xsdt.String(fmt.Sprint(width))
	application.Flash.Height = xsdt.String(fmt.Sprint(height))
	for key, value := range applicationParameters {
		param := &questionform.TApplicationParameter{}
		param.Name = xsdt.String(key)
		param.Value = xsdt.String(value)
		application.Flash.ApplicationParameters = append(
			application.Flash.ApplicationParameters, param)
	}

	content := question.getCurrentContent()
	content.Applications = append(content.Applications, application)
}

// Add a JavaApplet Application item to the most recent Question/Overview added.
func (question *QuestionForm) AddJavaAppletApplicationContent(
	appletFilename string, width, height int,
	applicationParameters map[string]string, appletPath *url.URL) {
	application := &questionform.TApplicationContentType{}
	application.JavaApplet = &questionform.TJavaAppletContentType{}
	application.JavaApplet.AppletFilename = xsdt.String(appletFilename)
	application.JavaApplet.Width = xsdt.String(fmt.Sprint(width))
	application.JavaApplet.Height = xsdt.String(fmt.Sprint(height))
	for key, value := range applicationParameters {
		param := &questionform.TApplicationParameter{}
		param.Name = xsdt.String(key)
		param.Value = xsdt.String(value)
		application.JavaApplet.ApplicationParameters = append(
			application.JavaApplet.ApplicationParameters, param)
	}
	if appletPath != nil {
		application.JavaApplet.AppletPath = questionform.TURLType(appletPath.String())
	}

	content := question.getCurrentContent()
	content.Applications = append(content.Applications, application)
}

// Add an EmbeddedBinary item to the most recent Question/Overview added.
func (question *QuestionForm) AddEmbeddedBinaryContent(
	dataURL *url.URL, altText string, width, height int,
	applicationParameters map[string]string, mimeType, mimeSubType string) {
	binary := &questionform.TEmbeddedBinaryContentType{}
	binary.DataURL = questionform.TURLType(dataURL.String())
	binary.AltText = xsdt.String(altText)
	binary.Width = xsdt.String(fmt.Sprint(width))
	binary.Height = xsdt.String(fmt.Sprint(height))
	for key, value := range applicationParameters {
		param := &questionform.TApplicationParameter{}
		param.Name = xsdt.String(key)
		param.Value = xsdt.String(value)
		binary.ApplicationParameters = append(
			binary.ApplicationParameters, param)
	}
	if mimeType != "" || mimeSubType != "" {
		binary.EmbeddedMimeType = &questionform.TEmbeddedMimeType{}
		if mimeType != "" {
			binary.EmbeddedMimeType.Type = xsdt.String(mimeType)
		}
		if mimeSubType != "" {
			binary.EmbeddedMimeType.SubType = xsdt.String(mimeSubType)
		}
	}

	content := question.getCurrentContent()
	content.EmbeddedBinaries = append(content.EmbeddedBinaries, binary)
}

// Add a Title item to the most recent Question/Overview added.
func (question *QuestionForm) AddTitleContent(title string) {
	content := question.getCurrentContent()
	content.Titles = append(content.Titles, xsdt.String(title))
}

// Add a Text item to the most recent Question/Overview added.
func (question *QuestionForm) AddTextContent(text string) {
	content := question.getCurrentContent()
	content.Texts = append(content.Texts, xsdt.String(text))
}

// Add default text for a free text answer specification
func (question *QuestionForm) AddFreeTextAnswerDefaultText(text string) {
	answer := question.getCurrentAnswer()
	if answer.FreeTextAnswer == nil {
		answer.FreeTextAnswer = &questionform.TFreeTextAnswerType{}
	}
	answer.FreeTextAnswer.DefaultText = xsdt.String(text)
}

// Add suggested number of lines for a free text answer specification
func (question *QuestionForm) AddFreeTextAnswerNumberOfLinesSuggestion(lines int) {
	answer := question.getCurrentAnswer()
	if answer.FreeTextAnswer == nil {
		answer.FreeTextAnswer = &questionform.TFreeTextAnswerType{}
	}
	answer.FreeTextAnswer.NumberOfLinesSuggestion = xsdt.PositiveInteger(lines)
}

// Add length constraints for a free text answer specification
func (question *QuestionForm) AddFreeTextAnswerLengthConstraints(min, max int) {
	answer := question.getCurrentAnswer()
	if answer.FreeTextAnswer == nil {
		answer.FreeTextAnswer = &questionform.TFreeTextAnswerType{}
	}
	if answer.FreeTextAnswer.Constraints == nil {
		answer.FreeTextAnswer.Constraints = &questionform.TxsdFreeTextAnswerTypeSequenceConstraints{}
	}
	if answer.FreeTextAnswer.Constraints.Length == nil {
		answer.FreeTextAnswer.Constraints.Length = &questionform.TxsdFreeTextAnswerTypeSequenceConstraintsSequenceLength{}
	}
	answer.FreeTextAnswer.Constraints.Length.MinLength = xsdt.NonNegativeInteger(min)
	answer.FreeTextAnswer.Constraints.Length.MaxLength = xsdt.PositiveInteger(max)
}

// Add numeric value constraints for a free text answer specification
func (question *QuestionForm) AddFreeTextAnswerNumericConstraints(min, max int) {
	answer := question.getCurrentAnswer()
	if answer.FreeTextAnswer == nil {
		answer.FreeTextAnswer = &questionform.TFreeTextAnswerType{}
	}
	if answer.FreeTextAnswer.Constraints == nil {
		answer.FreeTextAnswer.Constraints = &questionform.TxsdFreeTextAnswerTypeSequenceConstraints{}
	}
	if answer.FreeTextAnswer.Constraints.IsNumeric == nil {
		answer.FreeTextAnswer.Constraints.IsNumeric = &questionform.TxsdFreeTextAnswerTypeSequenceConstraintsSequenceIsNumeric{}
	}
	answer.FreeTextAnswer.Constraints.IsNumeric.MinValue = xsdt.Int(min)
	answer.FreeTextAnswer.Constraints.IsNumeric.MaxValue = xsdt.Int(max)
}

// Add minimum number of selections for a selection answer
func (question *QuestionForm) AddSelectionAnswerMinSelections(selections int) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	answer.SelectionAnswer.MinSelectionCount = xsdt.NonNegativeInteger(selections)
}

// Add maximum number of selections for a selection answer
func (question *QuestionForm) AddSelectionAnswerMaxSelections(selections int) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	answer.SelectionAnswer.MaxSelectionCount = xsdt.PositiveInteger(selections)
}

// Add style suggestion for a selection answer
func (question *QuestionForm) AddSelectionAnswerStyle(style string) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	answer.SelectionAnswer.StyleSuggestion = questionform.TxsdSelectionAnswerTypeSequenceStyleSuggestion(style)
}

// Add a text option for a selection answer
func (question *QuestionForm) AddSelectionAnswerTextSelection(
	selectionIdentifier, text string) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	if answer.SelectionAnswer.Selections == nil {
		answer.SelectionAnswer.Selections = &questionform.TxsdSelectionAnswerTypeSequenceSelections{}
	}
	selection := &questionform.TxsdSelectionAnswerTypeSequenceSelectionsSequenceSelection{}
	selection.SelectionIdentifier = xsdt.String(selectionIdentifier)
	selection.Text = xsdt.String(text)
	answer.SelectionAnswer.Selections.Selections = append(
		answer.SelectionAnswer.Selections.Selections, selection)
}

// Add a formatted content option for a selection answer
func (question *QuestionForm) AddSelectionAnswerFormattedContentSelection(
	selectionIdentifier, content string) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	if answer.SelectionAnswer.Selections == nil {
		answer.SelectionAnswer.Selections = &questionform.TxsdSelectionAnswerTypeSequenceSelections{}
	}
	selection := &questionform.TxsdSelectionAnswerTypeSequenceSelectionsSequenceSelection{}
	selection.SelectionIdentifier = xsdt.String(selectionIdentifier)
	selection.FormattedContent = xsdt.String(content)
	answer.SelectionAnswer.Selections.Selections = append(
		answer.SelectionAnswer.Selections.Selections, selection)
}

// Add a binary option for a selection answer
func (question *QuestionForm) AddSelectionAnswerBinarySelection(
	selectionIdentifier, mimeType, mimeSubType string,
	dataURL *url.URL, altText string) {
	answer := question.getCurrentAnswer()
	if answer.SelectionAnswer == nil {
		answer.SelectionAnswer = &questionform.TSelectionAnswerType{}
	}
	if answer.SelectionAnswer.Selections == nil {
		answer.SelectionAnswer.Selections = &questionform.TxsdSelectionAnswerTypeSequenceSelections{}
	}
	selection := &questionform.TxsdSelectionAnswerTypeSequenceSelectionsSequenceSelection{}
	selection.SelectionIdentifier = xsdt.String(selectionIdentifier)
	selection.Binary = &questionform.TBinaryContentType{}
	if mimeType != "" || mimeSubType != "" {
		selection.Binary.MimeType = &questionform.TMimeType{}
		if mimeType != "" {
			selection.Binary.MimeType.Type = questionform.TxsdMimeTypeSequenceType(mimeType)
		}
		if mimeSubType != "" {
			selection.Binary.MimeType.SubType = xsdt.String(mimeSubType)
		}
	}
	selection.Binary.DataURL = questionform.TURLType(dataURL.String())
	if altText != "" {
		selection.Binary.AltText = xsdt.String(altText)
	}
	answer.SelectionAnswer.Selections.Selections = append(
		answer.SelectionAnswer.Selections.Selections, selection)
}

// Add a file upload answer
func (question *QuestionForm) AddFileUploadAnswer(minSize, maxSize int) {
	answer := question.getCurrentAnswer()
	if answer.FileUploadAnswer == nil {
		answer.FileUploadAnswer = &questionform.TFileUploadAnswerType{}
	}
	answer.FileUploadAnswer.MinFileSizeInBytes = questionform.TMinFileSizeType(minSize)
	answer.FileUploadAnswer.MaxFileSizeInBytes = questionform.TMaxFileSizeType(maxSize)
}

// AnswerKey allows you to provide known answers against which workers can be
// graded. This is used, for example, for qualifications.
type AnswerKey struct {

	// The name of the wrapper element for an XML representation of the object
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd AnswerKey"`

	// The answer key.
	answerkey.TxsdAnswerKey
}

type QuestionFormAnswers struct {

	// The name of the wrapper element for an XML representation of the object
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd QuestionFormAnswers"`

	// The answers
	questionformanswers.TxsdQuestionFormAnswers
}
