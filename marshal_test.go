package openxml

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ppr struct {
	WrPr struct {
		WrFonts struct {
			Wascii    string `xml:"ascii,attr"`
			WeastAsia string `xml:"eastAsia,attr"`
			WhAnsi    string `xml:"hAnsi,attr"`
			Wcs       string `xml:"cs,attr"`
			Whint     string `xml:"hint,attr"`
		} `xml:"rFonts"`
		Wsz struct {
			Wval string `xml:"val,attr"`
		} `xml:"sz"`
		WszCs struct {
			Wval string `xml:"val,attr"`
		} `xml:"szCs"`
	} `xml:"rPr"`
	Wt string `xml:"t"`
}

type pprm struct {
	WrPr struct {
		WrFonts struct {
			Wascii    string `xml:"w:ascii,attr"`
			WeastAsia string `xml:"w:eastAsia,attr"`
			WhAnsi    string `xml:"w:hAnsi,attr"`
			Wcs       string `xml:"w:cs,attr"`
			Whint     string `xml:"w:hint,attr"`
		} `xml:"w:rFonts"`
		Wsz struct {
			Wval string `xml:"w:val,attr"`
		} `xml:"w:sz"`
		WszCs struct {
			Wval string `xml:"w:val,attr"`
		} `xml:"w:szCs"`
	} `xml:"w:rPr"`
	Wt string `xml:"w:t"`
}

func (*pprm) String() string {
	return "w:r"
}

func TestMarshal(t *testing.T) {
	org := []byte(`<w:r><w:rPr><w:rFonts w:ascii="华文黑体" w:eastAsia="华文黑体" w:hAnsi="华文黑体" w:cs="华文黑体" w:hint="eastAsia"/><w:sz w:val="36"/><w:szCs w:val="44"/></w:rPr><w:t>Test文本원본</w:t></w:r>`)
	m := ppr{}
	err := xml.Unmarshal(org, &m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m)
	data, err := Marshal(&m, (*pprm)(nil))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(org), string(data))
}

func TestMarshalIndent(t *testing.T) {
	org := []byte(`<w:r>
	<w:rPr>
		<w:rFonts w:ascii="华文黑体" w:eastAsia="华文黑体" w:hAnsi="华文黑体" w:cs="华文黑体" w:hint="eastAsia"/>
		<w:sz w:val="36"/>
		<w:szCs w:val="44"/>
	</w:rPr>
	<w:t>Test文本원본</w:t>
</w:r>`)
	m := ppr{}
	err := xml.Unmarshal(org, &m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m)
	data, err := MarshalIndent(&m, (*pprm)(nil), "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(org), string(data))
}
