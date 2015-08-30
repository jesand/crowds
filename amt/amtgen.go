package amt

//go:generate go get github.com/metaleap/go-xsd/xsd-makepkg
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurk/2014-08-15/AWSMechanicalTurkRequester.xsd
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2011-11-11/HTMLQuestion.xsd
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/AnswerKey.xsd
//go:generate $GOPATH/bin/xsd-makepkg -goinst=false -basepath=github.com/jesand/crowds/amt/gen -uri=http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd
