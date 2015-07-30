package language_test

import (
	"fmt"
	"github.com/slyrz/language"
	"testing"
)

// Random Wikipedia articles were used to create these test vectors.
var testVectors = []struct {
	Code string
	Text string
}{
	{"de", "Das Herbert-Dröse-Stadion ist ein Sportstadion in Hanau (Wilhelmsbad) im deutschen Bundesland Hessen."},
	{"en", "Route 233 is a provincial highway located in the Montérégie region of Quebec east of Montreal."},
	{"es", "Es un galardón honorífico entregado por la FIFA desde el 2001 a todas aquellas personas o entidades que contribuyen de forma significativa en el fútbol."},
	{"fr", "Baignes est une commune française, située dans le département de la Haute-Saône en région Franche-Comté."},
	{"it", "Il singolare del torneo di tennis Canberra Women's Classic 2005, facente parte del WTA Tour 2005, ha avuto come vincitrice Ana Ivanović che ha battuto in finale Melinda Czink 7-5, 6-1."},
	{"ja", "バレーボールシンガポール男子代表は、バレーボールの国際大会で編成されるシンガポールの男子バレーボールナショナルチームである。"},
	{"pl", "Krytyka Polityczna – lewicowe pismo społeczno-polityczne ukazujące się od 2002. Założycielem i redaktorem naczelnym pisma jest Sławomir Sierakowski."},
	{"th", "นายแพทย์สงวน นิตยารัมภ์พงศ์ เกิดเมื่อวันที่ 18 มีนาคม พ.ศ. 2495 เป็นลูกคนสุดท้องจากพี่น้องทั้งหมด 6 คนในครอบครัวชาวจีนในกรุงเทพมหานคร"},
	{"zh", "牙中站（朝鲜语：송천역；牙中驛）曾为韩国铁路全罗线上的一个车站，位于全罗北道全州市德津區牛牙洞，已于2011年废止。"},
	{"da", "Bose er et nedslagskrater på Månen. Det befinder sig på den sydlige halvkugle på Månens bagside og er opkaldt efter den bengalske polyhistor Chandra Bose (1858 – 1937)."},
	{"pt", "O Mosteiro da Escultura do Senhor é um histórico mosteiro localizado na cidade de Chã Grande, no estado brasileiro de Pernambuco."},
	{"un", ""},
}

func TestDetect(t *testing.T) {
	for _, test := range testVectors {
		lang := language.Detect(test.Text)
		t.Log(lang)
		if lang.Code != test.Code {
			t.Errorf("expected language %q, got %q", test.Code, lang.Code)
		}
	}
}

func TestDetectHTML(t *testing.T) {
	const format = `
<!DOCTYPE html>
<html lang="%[1]s">
<head>
<meta charset="UTF-8">
<title>%.10[2]s...</title>
</head>
<body>
%[2]s
</body>
</html>
`
	for _, test := range testVectors {
		html := fmt.Sprintf(format, test.Code, test.Text)
		lang := language.DetectHTML(html)
		t.Log(lang)
		if lang.Code != test.Code {
			t.Errorf("expected language %q, got %q", test.Code, lang.Code)
		}
	}
}
