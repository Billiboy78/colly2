package extensions

import (
	"github.com/billiboy/colly2/v2"
)

// URLLengthFilter filters out requests with URLs longer than URLLengthLimit
func URLLengthFilter(c *colly.Collector, URLLengthLimit int) {
	c.OnRequest(func(r *colly.Request) {
		if len(r.URL.String()) > URLLengthLimit {
			r.Abort()
		}
	})
}
