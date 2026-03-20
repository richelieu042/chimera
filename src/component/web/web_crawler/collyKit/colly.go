package collyKit

import (
	"github.com/gocolly/colly/v2"
)

func NewCollector(options ...colly.CollectorOption) *colly.Collector {
	return colly.NewCollector(options...)
}
