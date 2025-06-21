package main

import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {
	var m = make(map[string]int)
	for i := 0; i < len(text); i++ {
		m[string(text[i])]++
	}
	chab := m["b"]
	chaa := m["a"]
	chal := m["l"] / 2
	chao := m["o"] / 2
	charn := m["n"]

	return min(chab, chaa, chal, chao, chab, charn)
}

func main() {
	var s = "krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw"
	fmt.Println(maxNumberOfBalloons(s))
}
