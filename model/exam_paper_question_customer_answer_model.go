package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func ExamPaperQuestionCustomerAnswerAllCount() int64 {
	var count int64
	global.Db.Model(&entity.ExamPaperQuestionCustomerAnswer{}).Count(&count)
	return count
}

func SelectCustomerAnswerByStartEnd(start, end string) []entity.ExamPaperQuestionCustomerAnswerCount {
	var res []entity.ExamPaperQuestionCustomerAnswerCount
	err := global.Db.Raw("SELECT create_time AS `name`, COUNT( create_time ) AS `value` FROM ( SELECT DATE_FORMAT( create_time, '%Y-%m-%d' ) AS create_time FROM t_exam_paper_question_customer_answer WHERE create_time BETWEEN ? and ?) a GROUP BY create_time", start, end).Scan(&res).Error
	if err != nil {
		zap.L().Error("SelectEventLogByStartEnd", zap.Error(err))
	}
	return res
}
