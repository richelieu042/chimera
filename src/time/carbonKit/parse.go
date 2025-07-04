package carbonKit

import "github.com/dromara/carbon/v2"

var (
	// Parse 将标准时间字符串解析成 Carbon 实例.
	Parse = carbon.Parse

	// ParseByLayout 通过布局模板将时间字符串解析成 Carbon 实例.
	/*
		e.g.
		carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() 						// 2020-08-05 13:14:15
		carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() 			// 2020-08-05 13:14:15
		carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
		carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() 		// 2020-08-05 14:14:15
	*/
	ParseByLayout = carbon.ParseByLayout

	// ParseByFormat 通过格式模板将时间字符串解析成 Carbon 实例.
	/*
		e.g.
		carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() 						// 2020-08-05 13:14:15
		carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() 			// 2020-08-05 13:14:15
		carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
		carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() 		// 2020-08-05 14:14:15
	*/
	ParseByFormat = carbon.ParseByFormat
)
