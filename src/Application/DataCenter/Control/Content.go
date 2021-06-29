package Control

//var min = 0
//var _ = MessageLib.AddAction("Content.UpdateJob", func(fieldList []DataSourceLib.FieldModel) MessageLib.QueueItem {
//	var contentList []DataCenterModel.List
//	err := Manage.DataBase.Select([]DataSourceLib.FieldModel{
//		Manage.DataBase.Greater("Id", min),
//		Manage.DataBase.Page(1, 100),
//	}, &contentList, DataCenterModel.List{})
//	for _, v := range contentList {
//		status := true
//		content := DataCenterModel.Content{}
//		c := colly.NewCollector()
//
//		c.OnHTML(".title-article", func(element *colly.HTMLElement) {
//			html, err := element.DOM.Html()
//			if err != nil {
//				status = false
//				return
//			}
//			content.Title = html
//		})
//
//		c.OnHTML("#content_views", func(element *colly.HTMLElement) {
//			html, err := element.DOM.Html()
//			if err != nil {
//				status = false
//				return
//			}
//			content.Content = html
//		})
//
//		c.OnHTML("#artic-tag-box", func(element *colly.HTMLElement) {
//			html, err := element.DOM.Html()
//			if err != nil {
//				status = false
//				return
//			}
//			content.Tag = html
//		})
//
//		err := c.Visit(v.Url)
//		if err != nil {
//			return MessageLib.QueueItem{}
//		}
//	}
//
//	if err != nil {
//		return MessageLib.NewQueueItem(false, contentList)
//	}
//	return MessageLib.NewQueueItem(true, contentList)
//})
