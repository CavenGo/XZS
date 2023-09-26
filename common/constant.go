package common

const (
	UserEnable  = 1
	UserDisable = 2

	SingleChoice   = 1
	MultipleChoice = 2
	TrueFalse      = 3
	GapFilling     = 4
	ShortAnswer    = 5

	Fixed     = 1
	TimeLimit = 4
	Task      = 6
)

var QuestionTypeMap = map[int]string{
	1: "单选题",
	2: "多选题",
	3: "判断题",
	4: "填空题",
	5: "简答题",
}
