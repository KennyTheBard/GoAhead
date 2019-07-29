package html

type HTMLContent struct {
	Text string
}

func (content *HTMLContent) Tag(tag string) *HTMLContent {
	content.Text = "<" + tag + ">" + content.Text + "</" + tag + ">"
	return content
}

func (content1 *HTMLContent) Merge(content2 *HTMLContent) *HTMLContent {
	content1.Text += content2.Text
	return content1
}

func CreateContent(text string) HTMLContent {
	return HTMLContent{text}
}
