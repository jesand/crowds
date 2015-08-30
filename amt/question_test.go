package amt

import (
	"encoding/xml"
	answerkey "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd_go"
	questionformanswers "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestHTMLQuestion(t *testing.T) {
	Convey("Given XML for an HTMLQuestion", t, func() {
		xml := `<HTMLQuestion xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd">
  <HTMLContent><![CDATA[
<!DOCTYPE html>
<html>
 <head>
  <meta http-equiv='Content-Type' content='text/html; charset=UTF-8'/>
  <script type='text/javascript' src='https://s3.amazonaws.com/mturk-public/externalHIT_v1.js'></script>
 </head>
 <body>
  <form name='mturk_form' method='post' id='mturk_form' action='https://www.mturk.com/mturk/externalSubmit'>
  <input type='hidden' value='' name='assignmentId' id='assignmentId'/>
  <h1>What's up?</h1>
  <p><textarea name='comment' cols='80' rows='3'></textarea></p>
  <p><input type='submit' id='submitButton' value='Submit' /></p></form>
  <script language='Javascript'>turkSetAssignmentID();</script>
 </body>
</html>
]]>
  </HTMLContent>
  <FrameHeight>450</FrameHeight>
</HTMLQuestion>`
		Convey("When I Unmarshal the XML", func() {
			question, error := DecodeQuestion([]byte(xml))
			So(error, ShouldBeNil)
			So(question, ShouldHaveSameTypeAs, &HTMLQuestion{})
			result := question.(*HTMLQuestion)

			Convey("Then I get the expected HTMLQuestion", func() {
				expected := &HTMLQuestion{}
				expected.XMLName.Space = "http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd"
				expected.XMLName.Local = "HTMLQuestion"
				expected.FrameHeight = 450
				expected.HTMLContent = `
<!DOCTYPE html>
<html>
 <head>
  <meta http-equiv='Content-Type' content='text/html; charset=UTF-8'/>
  <script type='text/javascript' src='https://s3.amazonaws.com/mturk-public/externalHIT_v1.js'></script>
 </head>
 <body>
  <form name='mturk_form' method='post' id='mturk_form' action='https://www.mturk.com/mturk/externalSubmit'>
  <input type='hidden' value='' name='assignmentId' id='assignmentId'/>
  <h1>What's up?</h1>
  <p><textarea name='comment' cols='80' rows='3'></textarea></p>
  <p><input type='submit' id='submitButton' value='Submit' /></p></form>
  <script language='Javascript'>turkSetAssignmentID();</script>
 </body>
</html>

  `
				So(result, ShouldResemble, expected)
			})

			Convey("When I Marshal the HTMLQuestion", func() {
				newxml, err := EncodeQuestion(result)

				Convey("Then I get the expected XML", func() {
					So(err, ShouldBeNil)
					expected := `<HTMLQuestion xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd">` +
						`<HTMLContent>&#xA;&lt;!DOCTYPE html&gt;&#xA;` +
						`&lt;html&gt;&#xA; ` +
						`&lt;head&gt;&#xA;  ` +
						`&lt;meta http-equiv=&#39;Content-Type&#39; content=&#39;text/html; charset=UTF-8&#39;/&gt;&#xA;  ` +
						`&lt;script type=&#39;text/javascript&#39; src=&#39;https://s3.amazonaws.com/mturk-public/externalHIT_v1.js&#39;&gt;` +
						`&lt;/script&gt;&#xA; ` +
						`&lt;/head&gt;&#xA; ` +
						`&lt;body&gt;&#xA;  ` +
						`&lt;form name=&#39;mturk_form&#39; method=&#39;post&#39; id=&#39;mturk_form&#39; action=&#39;https://www.mturk.com/mturk/externalSubmit&#39;&gt;&#xA;  ` +
						`&lt;input type=&#39;hidden&#39; value=&#39;&#39; name=&#39;assignmentId&#39; id=&#39;assignmentId&#39;/&gt;&#xA;  ` +
						`&lt;h1&gt;What&#39;s up?` +
						`&lt;/h1&gt;&#xA;  ` +
						`&lt;p&gt;` +
						`&lt;textarea name=&#39;comment&#39; cols=&#39;80&#39; rows=&#39;3&#39;&gt;` +
						`&lt;/textarea&gt;` +
						`&lt;/p&gt;&#xA;  ` +
						`&lt;p&gt;` +
						`&lt;input type=&#39;submit&#39; id=&#39;submitButton&#39; value=&#39;Submit&#39; /&gt;` +
						`&lt;/p&gt;` +
						`&lt;/form&gt;&#xA;  ` +
						`&lt;script language=&#39;Javascript&#39;&gt;turkSetAssignmentID();` +
						`&lt;/script&gt;&#xA; ` +
						`&lt;/body&gt;&#xA;` +
						`&lt;/html&gt;&#xA;&#xA;  </HTMLContent>` +
						`<FrameHeight>450</FrameHeight>` +
						`</HTMLQuestion>`
					So(string(newxml), ShouldEqual, expected)
				})
			})
		})
	})
}

func TestExternalQuestion(t *testing.T) {
	Convey("Given XML for an ExternalQuestion", t, func() {
		xml := `
			<ExternalQuestion xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd">
			  <ExternalURL>https://tictactoe.amazon.com/gamesurvey.cgi?gameid=01523</ExternalURL>
			  <FrameHeight>400</FrameHeight>
			</ExternalQuestion>`
		Convey("When I Unmarshal the XML", func() {
			question, error := DecodeQuestion([]byte(xml))
			So(error, ShouldBeNil)
			So(question, ShouldHaveSameTypeAs, &ExternalQuestion{})
			result := question.(*ExternalQuestion)

			Convey("Then I get the expected ExternalQuestion", func() {
				expected := &ExternalQuestion{}
				expected.XMLName.Space = "http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd"
				expected.XMLName.Local = "ExternalQuestion"
				expected.FrameHeight = 400
				expected.ExternalURL = "https://tictactoe.amazon.com/gamesurvey.cgi?gameid=01523"
				So(result, ShouldResemble, expected)
			})

			Convey("When I Marshal the ExternalQuestion", func() {
				newxml, err := EncodeQuestion(result)

				Convey("Then I get the expected XML", func() {
					So(err, ShouldBeNil)
					expected := `<ExternalQuestion xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd">` +
						`<FrameHeight>400</FrameHeight>` +
						`<ExternalURL>https://tictactoe.amazon.com/gamesurvey.cgi?gameid=01523</ExternalURL>` +
						`</ExternalQuestion>`
					So(string(newxml), ShouldEqual, expected)
				})
			})
		})
	})
}

func TestQuestionForm(t *testing.T) {
	Convey("Given XML for an QuestionForm", t, func() {
		innerxml := `<Overview>
			    <Title>Game 01523, "X" to play</Title>
			    <Text>
			      You are helping to decide the next move in a game of Tic-Tac-Toe.  The board looks like this:
			    </Text>
			    <Binary>
			      <MimeType>
			        <Type>image</Type>
			        <SubType>gif</SubType>
			      </MimeType>
			      <DataURL>http://tictactoe.amazon.com/game/01523/board.gif</DataURL>
			      <AltText>The game board, with "X" to move.</AltText>
			    </Binary>
			    <Text>
			      Player "X" has the next move.
			    </Text>
			  </Overview>
			  <Question>
			    <QuestionIdentifier>nextmove</QuestionIdentifier>
			    <DisplayName>The Next Move</DisplayName>
			    <IsRequired>true</IsRequired>
			    <QuestionContent>
			      <Text>
			        What are the coordinates of the best move for player "X" in this game?
			      </Text>
			    </QuestionContent>
			    <AnswerSpecification>
			      <FreeTextAnswer>
			        <Constraints>
			          <Length minLength="2" maxLength="2" />
			        </Constraints>
			        <DefaultText>C1</DefaultText>
			      </FreeTextAnswer>
			    </AnswerSpecification>
			  </Question>
			  <Question>
			    <QuestionIdentifier>likelytowin</QuestionIdentifier>
			    <DisplayName>The Next Move</DisplayName>
			    <IsRequired>true</IsRequired>
			    <QuestionContent>
			      <Text>
			        How likely is it that player "X" will win this game?
			      </Text>
			    </QuestionContent>
			    <AnswerSpecification>
			      <SelectionAnswer>
			        <StyleSuggestion>radiobutton</StyleSuggestion>
			        <Selections>
			          <Selection>
			            <SelectionIdentifier>notlikely</SelectionIdentifier>
			            <Text>Not likely</Text>
			          </Selection>
			          <Selection>
			            <SelectionIdentifier>unsure</SelectionIdentifier>
			            <Text>It could go either way</Text>
			          </Selection>
			          <Selection>
			            <SelectionIdentifier>likely</SelectionIdentifier>
			            <Text>Likely</Text>
			          </Selection>
			        </Selections>
			      </SelectionAnswer>
			    </AnswerSpecification>
			  </Question>`
		xml := `<QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd">` +
			innerxml + `</QuestionForm>`
		Convey("When I Unmarshal the XML", func() {
			question, error := DecodeQuestion([]byte(xml))
			So(error, ShouldBeNil)
			So(question, ShouldHaveSameTypeAs, &QuestionForm{})
			result := question.(*QuestionForm)

			Convey("Then I get the expected QuestionForm", func() {
				So(result.XMLContent, ShouldEqual, innerxml)

				var exp QuestionForm
				exp.XMLName.Space = "http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd"
				exp.XMLName.Local = "QuestionForm"
				exp.AddOverview()
				exp.AddTitleContent(`Game 01523, "X" to play`)
				exp.AddTextContent(`
			      You are helping to decide the next move in a game of Tic-Tac-Toe.  The board looks like this:
			    `)
				binUrl, _ := url.Parse("http://tictactoe.amazon.com/game/01523/board.gif")
				exp.AddBinaryContent("image", "gif", binUrl, `The game board, with "X" to move.`)
				exp.AddTextContent(`
			      Player "X" has the next move.
			    `)

				exp.AddQuestion("nextmove", "The Next Move", true)
				exp.AddTextContent(`
			        What are the coordinates of the best move for player "X" in this game?
			      `)
				exp.AddFreeTextAnswerDefaultText("C1")

				// NOTE: The generated code cannot unmarshal MinLength and MaxLength
				// because it places a namespace in the attr tags. Delete the namespace
				// to get it working. See:
				// https://github.com/metaleap/go-xsd/issues/18
				exp.AddFreeTextAnswerLengthConstraints(2, 2)

				exp.AddQuestion("likelytowin", "The Next Move", true)
				exp.AddTextContent(`
			        How likely is it that player "X" will win this game?
			      `)
				exp.AddSelectionAnswerStyle("radiobutton")
				exp.AddSelectionAnswerTextSelection("notlikely", "Not likely")
				exp.AddSelectionAnswerTextSelection("unsure", "It could go either way")
				exp.AddSelectionAnswerTextSelection("likely", "Likely")

				result.XMLContent = ""
				exp.addedOverviewNext = nil
				So(result.Overviews[0], ShouldResemble, exp.Overviews[0])
				So(result.Questions[0], ShouldResemble, exp.Questions[0])
				So(result.Questions[1], ShouldResemble, exp.Questions[1])
				So(result, ShouldResemble, &exp)
			})

			SkipConvey("When I Marshal the QuestionForm", func() {
				// TODO: Skipped because Marshal isn't smart enough to
				// preserve the order of subelements, and order matters here.

				result.XMLContent = ""
				newxml, err := EncodeQuestion(result)

				Convey("Then I get the expected XML", func() {
					So(err, ShouldBeNil)
					So(string(newxml), ShouldEqual,
						`<QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd">`+
							`<Overview>`+
							`<Title>Game 01523, "X" to play</Title>`+
							`<Text>
			      You are helping to decide the next move in a game of Tic-Tac-Toe.  The board looks like this:
			    </Text>`+
							`<Binary>`+
							`<MimeType>`+
							`<SubType>gif</SubType>`+
							`<Type>image</Type>`+
							`</MimeType>`+
							`<DataURL>http://tictactoe.amazon.com/game/01523/board.gif</DataURL>`+
							`<AltText>The game board, with "X" to move.</AltText>`+
							`</Binary>`+
							`<Text>
			      Player "X" has the next move.
			    </Text>`+
							`</Overview>`+
							`<Question>`+
							`<QuestionIdentifier>nextmove</QuestionIdentifier>`+
							`<DisplayName>The Next Move</DisplayName>`+
							`<IsRequired>true</IsRequired>`+
							`<QuestionContent>`+
							`<Text>
			        What are the coordinates of the best move for player "X" in this game?
			      </Text>`+
							`</QuestionContent>`+
							`<AnswerSpecification>`+
							`<FreeTextAnswer>`+
							`<Constraints>`+
							`<Length minLength="2" maxLength="2">`+
							`</Length>`+
							`</Constraints>`+
							`<DefaultText>C1</DefaultText>`+
							`<NumberOfLinesSuggestion>0</NumberOfLinesSuggestion>`+
							`</FreeTextAnswer>`+
							`</AnswerSpecification>`+
							`</Question>`+
							`<Question>`+
							`<QuestionIdentifier>likelytowin</QuestionIdentifier>`+
							`<DisplayName>The Next Move</DisplayName>`+
							`<IsRequired>true</IsRequired>`+
							`<QuestionContent>`+
							`<Text>
			        How likely is it that player "X" will win this game?
			      </Text>`+
							`</QuestionContent>`+
							`<AnswerSpecification>`+
							`<SelectionAnswer>`+
							`<MinSelectionCount>0</MinSelectionCount>`+
							`<MaxSelectionCount>0</MaxSelectionCount>`+
							`<StyleSuggestion>radiobutton</StyleSuggestion>`+
							`<Selections>`+
							`<Selection>`+
							`<FormattedContent>`+
							`</FormattedContent>`+
							`<SelectionIdentifier>notlikely</SelectionIdentifier>`+
							`<Text>Not likely</Text>`+
							`</Selection>`+
							`<Selection>`+
							`<FormattedContent>`+
							`</FormattedContent>`+
							`<SelectionIdentifier>unsure</SelectionIdentifier>`+
							`<Text>It could go either way</Text>`+
							`</Selection>`+
							`<Selection>`+
							`<FormattedContent>`+
							`</FormattedContent>`+
							`<SelectionIdentifier>likely</SelectionIdentifier>`+
							`<Text>Likely</Text>`+
							`</Selection>`+
							`</Selections>`+
							`</SelectionAnswer>`+
							`</AnswerSpecification>`+
							`</Question>`+
							`</QuestionForm>`)
				})
			})
		})
	})
}

func TestAnswerKey(t *testing.T) {
	Convey("Given XML for an AnswerKey", t, func() {
		keyxml := `<AnswerKey xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd">
  <Question>
    <QuestionIdentifier>nextmove</QuestionIdentifier>
    <AnswerOption>
      <SelectionIdentifier>D</SelectionIdentifier>
      <AnswerScore>5</AnswerScore>
    </AnswerOption>
  </Question>
  <Question>
    <QuestionIdentifier>favoritefruit</QuestionIdentifier>
    <AnswerOption>
      <SelectionIdentifier>apples</SelectionIdentifier>
      <AnswerScore>10</AnswerScore>
    </AnswerOption>
  </Question>
  <QualificationValueMapping>
    <PercentageMapping>
      <MaximumSummedScore>15</MaximumSummedScore>
    </PercentageMapping>
  </QualificationValueMapping>
</AnswerKey>`
		Convey("When I Unmarshal the XML", func() {
			var result AnswerKey
			err := xml.Unmarshal([]byte(keyxml), &result)
			So(err, ShouldBeNil)

			Convey("Then I get the expected AnswerKey", func() {
				expected := AnswerKey{}
				expected.XMLName.Space = "http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd"
				expected.XMLName.Local = "AnswerKey"
				expected.Questions = append(expected.Questions,
					&answerkey.TxsdAnswerKeySequenceQuestion{},
					&answerkey.TxsdAnswerKeySequenceQuestion{})

				expected.Questions[0].QuestionIdentifier = xsdt.String("nextmove")
				expected.Questions[0].AnswerOptions = append(expected.Questions[0].AnswerOptions, &answerkey.TxsdAnswerKeySequenceQuestionSequenceAnswerOption{})
				expected.Questions[0].AnswerOptions[0].SelectionIdentifiers = []xsdt.String{"D"}
				expected.Questions[0].AnswerOptions[0].AnswerScore = xsdt.Int(5)

				expected.Questions[1].QuestionIdentifier = xsdt.String("favoritefruit")
				expected.Questions[1].AnswerOptions = append(expected.Questions[1].AnswerOptions, &answerkey.TxsdAnswerKeySequenceQuestionSequenceAnswerOption{})
				expected.Questions[1].AnswerOptions[0].SelectionIdentifiers = []xsdt.String{"apples"}
				expected.Questions[1].AnswerOptions[0].AnswerScore = xsdt.Int(10)

				expected.QualificationValueMapping = &answerkey.TxsdAnswerKeySequenceQualificationValueMapping{}
				expected.QualificationValueMapping.PercentageMapping = &answerkey.TxsdAnswerKeySequenceQualificationValueMappingChoicePercentageMapping{}
				expected.QualificationValueMapping.PercentageMapping.MaximumSummedScore = xsdt.Int(15)

				So(result, ShouldResemble, expected)
			})

			Convey("When I Marshal the AnswerKey", func() {
				newxml, err := EncodeQuestion(result)

				Convey("Then I get the expected XML", func() {
					So(err, ShouldBeNil)
					expected := `<AnswerKey xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd">` +
						`<Question>` +
						`<QuestionIdentifier>nextmove</QuestionIdentifier>` +
						`<AnswerOption>` +
						`<SelectionIdentifier>D</SelectionIdentifier>` +
						`<AnswerScore>5</AnswerScore>` +
						`</AnswerOption>` +
						`<DefaultScore>0</DefaultScore>` +
						`</Question>` +
						`<Question>` +
						`<QuestionIdentifier>favoritefruit</QuestionIdentifier>` +
						`<AnswerOption>` +
						`<SelectionIdentifier>apples</SelectionIdentifier>` +
						`<AnswerScore>10</AnswerScore>` +
						`</AnswerOption>` +
						`<DefaultScore>0</DefaultScore>` +
						`</Question>` +
						`<QualificationValueMapping>` +
						`<PercentageMapping>` +
						`<MaximumSummedScore>15</MaximumSummedScore>` +
						`</PercentageMapping>` +
						`</QualificationValueMapping>` +
						`</AnswerKey>`
					So(string(newxml), ShouldEqual, expected)
				})
			})
		})
	})
}

func TestQuestionFormAnswers(t *testing.T) {
	Convey("Given XML for an QuestionFormAnswers", t, func() {
		keyxml := `<QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd">
  <Answer>
    <QuestionIdentifier>nextmove</QuestionIdentifier>
    <FreeText>C3</FreeText>
  </Answer>
  <Answer>
    <QuestionIdentifier>likelytowin</QuestionIdentifier>
    <SelectionIdentifier>notlikely</SelectionIdentifier>
  </Answer>
</QuestionFormAnswers>`
		Convey("When I Unmarshal the XML", func() {
			var result QuestionFormAnswers
			err := xml.Unmarshal([]byte(keyxml), &result)
			So(err, ShouldBeNil)

			Convey("Then I get the expected QuestionFormAnswers", func() {
				expected := QuestionFormAnswers{}
				expected.XMLName.Space = "http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd"
				expected.XMLName.Local = "QuestionFormAnswers"

				expected.Answers = append(expected.Answers,
					&questionformanswers.TxsdQuestionFormAnswersSequenceAnswer{},
					&questionformanswers.TxsdQuestionFormAnswersSequenceAnswer{})

				expected.Answers[0].QuestionIdentifier = "nextmove"
				expected.Answers[0].FreeText = "C3"

				expected.Answers[1].QuestionIdentifier = "likelytowin"
				expected.Answers[1].SelectionIdentifiers = []xsdt.String{"notlikely"}

				So(result.Answers[0], ShouldResemble, expected.Answers[0])
				So(result.Answers[1], ShouldResemble, expected.Answers[1])
				So(result, ShouldResemble, expected)
			})

			SkipConvey("When I Marshal the QuestionFormAnswers", func() {
				// TODO: Skipped because it marshals various empty fields

				newxml, err := EncodeQuestion(result)

				Convey("Then I get the expected XML", func() {
					So(err, ShouldBeNil)
					expected := `<QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd">` +
						`<Answer>` +
						`<FreeText>C3</FreeText>` +
						`<QuestionIdentifier>nextmove</QuestionIdentifier>` +
						`</Answer>` +
						`<Answer>` +
						`<QuestionIdentifier>likelytowin</QuestionIdentifier>` +
						`<SelectionIdentifier>notlikely</SelectionIdentifier>` +
						`</Answer>` +
						`</QuestionFormAnswers>`
					So(string(newxml), ShouldEqual, expected)
				})
			})
		})
	})
}
