package main

func (h *HtmlJsxCompiler) Tokenize(content []byte) ([]Token, error) {
	for h.Index < h.Length {

		if h.Index >= h.Length {
			return nil, nil
		}

		if h.timeToGetAttribute || h.timeToGetInnerText {

			if h.timeToGetAttribute {
				h.Index += 1
				if isChar(h.Content[h.Index]) {
					str := h.tokAttribute()
					h.Tokens = append(h.Tokens, Token{
						Symbol: "ATTR",
						Value:  str,
					})
				}
				h.timeToGetAttribute = false
				// lg(h.Tokens)
				continue
			}

			if h.timeToGetInnerText {
				str := ""
				for isChar(h.Content[h.Index]) {
					str += string(h.Content[h.Index])
					h.Index += 1
				}
				h.Tokens = append(h.Tokens, Token{
					Symbol: "INNERTEXT",
					Value:  str,
				})
			}
			h.timeToGetInnerText = false
			// lg(h.Tokens)
			continue
		}

		if isLeft(h.Content[h.Index]) {
			h.Tokens = append(h.Tokens, Token{
				Symbol: "LEFT",
				Value:  "<",
			})
			h.Index += 1
			continue
		}

		if isSlash(h.Content[h.Index]) {
			str := "/"
			h.Index += 1
			for isChar(h.Content[h.Index]) {
				str += string(h.Content[h.Index])
				h.Index += 1
			}
			h.Tokens = append(h.Tokens, Token{
				Symbol: "CLOSETAG",
				Value:  str,
			})
			// lg(h.Tokens)
		}

		if isChar(h.Content[h.Index]) {
			str := ""
			for isChar(h.Content[h.Index]) {
				str += string(h.Content[h.Index])
				h.Index += 1
			}

			if str == "div" {
				h.Tokens = append(h.Tokens, Token{
					Symbol: "TAG",
					Value:  "div",
				})
				// lg(h.Tokens)
				h.timeToGetAttribute = true
				continue
			}
		}

		if isRight(h.Content[h.Index]) {
			h.Tokens = append(h.Tokens, Token{
				Symbol: "RIGHT",
				Value:  ">",
			})
			h.Index += 1
			// lg(h.Tokens)
			h.timeToGetInnerText = true
			continue
		}
		h.Index += 1
	}
	return h.Tokens, nil
}

func isLeft(c byte) bool {
	return c == byte('<')
}

func isRight(c byte) bool {
	return c == byte('>')
}

func isSlash(c byte) bool {
	return c == byte('/')
}

func isChar(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c >= 'Z'
}

func (h *HtmlJsxCompiler) tokAttribute() string {
	str := ""
	for h.Content[h.Index] != '>' {
		str += string(h.Content[h.Index])
		h.Index += 1
	}
	return str
}
