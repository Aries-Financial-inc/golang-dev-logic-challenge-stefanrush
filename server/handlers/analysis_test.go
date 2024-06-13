package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	requestBody          = `[{"type": "call", "long_short": "long", "strike_price": 250, "expiration_date": "2021-01-01T00:00:00Z", "bid": 5.54, "ask": 6.61}]`
	staticUUID           = `1f680f4d-e173-4b18-8cda-289c12e983ba`
	expectedResponseBody = `[{"id":"1f680f4d-e173-4b18-8cda-289c12e983ba","type":"call","long_short":"long","strike_price":250,"expiration_date":"2021-01-01T00:00:00Z","bid":5.54,"ask":6.61,"is_valid":true,"analysis_result":{"bid_graph_data":[{"x":125,"y":-554},{"x":126,"y":-554},{"x":127,"y":-554},{"x":128,"y":-554},{"x":129,"y":-554},{"x":130,"y":-554},{"x":131,"y":-554},{"x":132,"y":-554},{"x":133,"y":-554},{"x":134,"y":-554},{"x":135,"y":-554},{"x":136,"y":-554},{"x":137,"y":-554},{"x":138,"y":-554},{"x":139,"y":-554},{"x":140,"y":-554},{"x":141,"y":-554},{"x":142,"y":-554},{"x":143,"y":-554},{"x":144,"y":-554},{"x":145,"y":-554},{"x":146,"y":-554},{"x":147,"y":-554},{"x":148,"y":-554},{"x":149,"y":-554},{"x":150,"y":-554},{"x":151,"y":-554},{"x":152,"y":-554},{"x":153,"y":-554},{"x":154,"y":-554},{"x":155,"y":-554},{"x":156,"y":-554},{"x":157,"y":-554},{"x":158,"y":-554},{"x":159,"y":-554},{"x":160,"y":-554},{"x":161,"y":-554},{"x":162,"y":-554},{"x":163,"y":-554},{"x":164,"y":-554},{"x":165,"y":-554},{"x":166,"y":-554},{"x":167,"y":-554},{"x":168,"y":-554},{"x":169,"y":-554},{"x":170,"y":-554},{"x":171,"y":-554},{"x":172,"y":-554},{"x":173,"y":-554},{"x":174,"y":-554},{"x":175,"y":-554},{"x":176,"y":-554},{"x":177,"y":-554},{"x":178,"y":-554},{"x":179,"y":-554},{"x":180,"y":-554},{"x":181,"y":-554},{"x":182,"y":-554},{"x":183,"y":-554},{"x":184,"y":-554},{"x":185,"y":-554},{"x":186,"y":-554},{"x":187,"y":-554},{"x":188,"y":-554},{"x":189,"y":-554},{"x":190,"y":-554},{"x":191,"y":-554},{"x":192,"y":-554},{"x":193,"y":-554},{"x":194,"y":-554},{"x":195,"y":-554},{"x":196,"y":-554},{"x":197,"y":-554},{"x":198,"y":-554},{"x":199,"y":-554},{"x":200,"y":-554},{"x":201,"y":-554},{"x":202,"y":-554},{"x":203,"y":-554},{"x":204,"y":-554},{"x":205,"y":-554},{"x":206,"y":-554},{"x":207,"y":-554},{"x":208,"y":-554},{"x":209,"y":-554},{"x":210,"y":-554},{"x":211,"y":-554},{"x":212,"y":-554},{"x":213,"y":-554},{"x":214,"y":-554},{"x":215,"y":-554},{"x":216,"y":-554},{"x":217,"y":-554},{"x":218,"y":-554},{"x":219,"y":-554},{"x":220,"y":-554},{"x":221,"y":-554},{"x":222,"y":-554},{"x":223,"y":-554},{"x":224,"y":-554},{"x":225,"y":-554},{"x":226,"y":-554},{"x":227,"y":-554},{"x":228,"y":-554},{"x":229,"y":-554},{"x":230,"y":-554},{"x":231,"y":-554},{"x":232,"y":-554},{"x":233,"y":-554},{"x":234,"y":-554},{"x":235,"y":-554},{"x":236,"y":-554},{"x":237,"y":-554},{"x":238,"y":-554},{"x":239,"y":-554},{"x":240,"y":-554},{"x":241,"y":-554},{"x":242,"y":-554},{"x":243,"y":-554},{"x":244,"y":-554},{"x":245,"y":-554},{"x":246,"y":-554},{"x":247,"y":-554},{"x":248,"y":-554},{"x":249,"y":-554},{"x":250,"y":-554},{"x":251,"y":-454},{"x":252,"y":-354},{"x":253,"y":-254},{"x":254,"y":-154},{"x":255,"y":-54},{"x":256,"y":46},{"x":257,"y":146},{"x":258,"y":246},{"x":259,"y":346},{"x":260,"y":446},{"x":261,"y":546},{"x":262,"y":646},{"x":263,"y":746},{"x":264,"y":846},{"x":265,"y":946},{"x":266,"y":1046},{"x":267,"y":1146},{"x":268,"y":1246},{"x":269,"y":1346},{"x":270,"y":1446},{"x":271,"y":1546},{"x":272,"y":1646},{"x":273,"y":1746},{"x":274,"y":1846},{"x":275,"y":1946},{"x":276,"y":2046},{"x":277,"y":2146},{"x":278,"y":2246},{"x":279,"y":2346},{"x":280,"y":2446},{"x":281,"y":2546},{"x":282,"y":2646},{"x":283,"y":2746},{"x":284,"y":2846},{"x":285,"y":2946},{"x":286,"y":3046},{"x":287,"y":3146},{"x":288,"y":3246},{"x":289,"y":3346},{"x":290,"y":3446},{"x":291,"y":3546},{"x":292,"y":3646},{"x":293,"y":3746},{"x":294,"y":3846},{"x":295,"y":3946},{"x":296,"y":4046},{"x":297,"y":4146},{"x":298,"y":4246},{"x":299,"y":4346},{"x":300,"y":4446},{"x":301,"y":4546},{"x":302,"y":4646},{"x":303,"y":4746},{"x":304,"y":4846},{"x":305,"y":4946},{"x":306,"y":5046},{"x":307,"y":5146},{"x":308,"y":5246},{"x":309,"y":5346},{"x":310,"y":5446},{"x":311,"y":5546},{"x":312,"y":5646},{"x":313,"y":5746},{"x":314,"y":5846},{"x":315,"y":5946},{"x":316,"y":6046},{"x":317,"y":6146},{"x":318,"y":6246},{"x":319,"y":6346},{"x":320,"y":6446},{"x":321,"y":6546},{"x":322,"y":6646},{"x":323,"y":6746},{"x":324,"y":6846},{"x":325,"y":6946},{"x":326,"y":7046},{"x":327,"y":7146},{"x":328,"y":7246},{"x":329,"y":7346},{"x":330,"y":7446},{"x":331,"y":7546},{"x":332,"y":7646},{"x":333,"y":7746},{"x":334,"y":7846},{"x":335,"y":7946},{"x":336,"y":8046},{"x":337,"y":8146},{"x":338,"y":8246},{"x":339,"y":8346},{"x":340,"y":8446},{"x":341,"y":8546},{"x":342,"y":8646},{"x":343,"y":8746},{"x":344,"y":8846},{"x":345,"y":8946},{"x":346,"y":9046},{"x":347,"y":9146},{"x":348,"y":9246},{"x":349,"y":9346},{"x":350,"y":9446},{"x":351,"y":9546},{"x":352,"y":9646},{"x":353,"y":9746},{"x":354,"y":9846},{"x":355,"y":9946},{"x":356,"y":10046},{"x":357,"y":10146},{"x":358,"y":10246},{"x":359,"y":10346},{"x":360,"y":10446},{"x":361,"y":10546},{"x":362,"y":10646},{"x":363,"y":10746},{"x":364,"y":10846},{"x":365,"y":10946},{"x":366,"y":11046},{"x":367,"y":11146},{"x":368,"y":11246},{"x":369,"y":11346},{"x":370,"y":11446},{"x":371,"y":11546},{"x":372,"y":11646},{"x":373,"y":11746},{"x":374,"y":11846},{"x":375,"y":11946},{"x":376,"y":12046},{"x":377,"y":12146},{"x":378,"y":12246},{"x":379,"y":12346},{"x":380,"y":12446},{"x":381,"y":12546},{"x":382,"y":12646},{"x":383,"y":12746},{"x":384,"y":12846},{"x":385,"y":12946},{"x":386,"y":13046},{"x":387,"y":13146},{"x":388,"y":13246},{"x":389,"y":13346},{"x":390,"y":13446},{"x":391,"y":13546},{"x":392,"y":13646},{"x":393,"y":13746},{"x":394,"y":13846},{"x":395,"y":13946},{"x":396,"y":14046},{"x":397,"y":14146},{"x":398,"y":14246},{"x":399,"y":14346},{"x":400,"y":14446},{"x":401,"y":14546},{"x":402,"y":14646},{"x":403,"y":14746},{"x":404,"y":14846},{"x":405,"y":14946},{"x":406,"y":15046},{"x":407,"y":15146},{"x":408,"y":15246},{"x":409,"y":15346},{"x":410,"y":15446},{"x":411,"y":15546},{"x":412,"y":15646},{"x":413,"y":15746},{"x":414,"y":15846},{"x":415,"y":15946},{"x":416,"y":16046},{"x":417,"y":16146},{"x":418,"y":16246},{"x":419,"y":16346},{"x":420,"y":16446},{"x":421,"y":16546},{"x":422,"y":16646},{"x":423,"y":16746},{"x":424,"y":16846},{"x":425,"y":16946},{"x":426,"y":17046},{"x":427,"y":17146},{"x":428,"y":17246},{"x":429,"y":17346},{"x":430,"y":17446},{"x":431,"y":17546},{"x":432,"y":17646},{"x":433,"y":17746},{"x":434,"y":17846},{"x":435,"y":17946},{"x":436,"y":18046},{"x":437,"y":18146},{"x":438,"y":18246},{"x":439,"y":18346},{"x":440,"y":18446},{"x":441,"y":18546},{"x":442,"y":18646},{"x":443,"y":18746},{"x":444,"y":18846},{"x":445,"y":18946},{"x":446,"y":19046},{"x":447,"y":19146},{"x":448,"y":19246},{"x":449,"y":19346},{"x":450,"y":19446},{"x":451,"y":19546},{"x":452,"y":19646},{"x":453,"y":19746},{"x":454,"y":19846},{"x":455,"y":19946},{"x":456,"y":20046},{"x":457,"y":20146},{"x":458,"y":20246},{"x":459,"y":20346},{"x":460,"y":20446},{"x":461,"y":20546},{"x":462,"y":20646},{"x":463,"y":20746},{"x":464,"y":20846},{"x":465,"y":20946},{"x":466,"y":21046},{"x":467,"y":21146},{"x":468,"y":21246},{"x":469,"y":21346},{"x":470,"y":21446},{"x":471,"y":21546},{"x":472,"y":21646},{"x":473,"y":21746},{"x":474,"y":21846},{"x":475,"y":21946},{"x":476,"y":22046},{"x":477,"y":22146},{"x":478,"y":22246},{"x":479,"y":22346},{"x":480,"y":22446},{"x":481,"y":22546},{"x":482,"y":22646},{"x":483,"y":22746},{"x":484,"y":22846},{"x":485,"y":22946},{"x":486,"y":23046},{"x":487,"y":23146},{"x":488,"y":23246},{"x":489,"y":23346},{"x":490,"y":23446},{"x":491,"y":23546},{"x":492,"y":23646},{"x":493,"y":23746},{"x":494,"y":23846},{"x":495,"y":23946},{"x":496,"y":24046},{"x":497,"y":24146},{"x":498,"y":24246},{"x":499,"y":24346},{"x":500,"y":24446}],"bid_pl_data":{"max_profit":1.7976931348623157e+308,"max_loss":554,"break_even_point":255.54},"ask_graph_data":[{"x":125,"y":-661},{"x":126,"y":-661},{"x":127,"y":-661},{"x":128,"y":-661},{"x":129,"y":-661},{"x":130,"y":-661},{"x":131,"y":-661},{"x":132,"y":-661},{"x":133,"y":-661},{"x":134,"y":-661},{"x":135,"y":-661},{"x":136,"y":-661},{"x":137,"y":-661},{"x":138,"y":-661},{"x":139,"y":-661},{"x":140,"y":-661},{"x":141,"y":-661},{"x":142,"y":-661},{"x":143,"y":-661},{"x":144,"y":-661},{"x":145,"y":-661},{"x":146,"y":-661},{"x":147,"y":-661},{"x":148,"y":-661},{"x":149,"y":-661},{"x":150,"y":-661},{"x":151,"y":-661},{"x":152,"y":-661},{"x":153,"y":-661},{"x":154,"y":-661},{"x":155,"y":-661},{"x":156,"y":-661},{"x":157,"y":-661},{"x":158,"y":-661},{"x":159,"y":-661},{"x":160,"y":-661},{"x":161,"y":-661},{"x":162,"y":-661},{"x":163,"y":-661},{"x":164,"y":-661},{"x":165,"y":-661},{"x":166,"y":-661},{"x":167,"y":-661},{"x":168,"y":-661},{"x":169,"y":-661},{"x":170,"y":-661},{"x":171,"y":-661},{"x":172,"y":-661},{"x":173,"y":-661},{"x":174,"y":-661},{"x":175,"y":-661},{"x":176,"y":-661},{"x":177,"y":-661},{"x":178,"y":-661},{"x":179,"y":-661},{"x":180,"y":-661},{"x":181,"y":-661},{"x":182,"y":-661},{"x":183,"y":-661},{"x":184,"y":-661},{"x":185,"y":-661},{"x":186,"y":-661},{"x":187,"y":-661},{"x":188,"y":-661},{"x":189,"y":-661},{"x":190,"y":-661},{"x":191,"y":-661},{"x":192,"y":-661},{"x":193,"y":-661},{"x":194,"y":-661},{"x":195,"y":-661},{"x":196,"y":-661},{"x":197,"y":-661},{"x":198,"y":-661},{"x":199,"y":-661},{"x":200,"y":-661},{"x":201,"y":-661},{"x":202,"y":-661},{"x":203,"y":-661},{"x":204,"y":-661},{"x":205,"y":-661},{"x":206,"y":-661},{"x":207,"y":-661},{"x":208,"y":-661},{"x":209,"y":-661},{"x":210,"y":-661},{"x":211,"y":-661},{"x":212,"y":-661},{"x":213,"y":-661},{"x":214,"y":-661},{"x":215,"y":-661},{"x":216,"y":-661},{"x":217,"y":-661},{"x":218,"y":-661},{"x":219,"y":-661},{"x":220,"y":-661},{"x":221,"y":-661},{"x":222,"y":-661},{"x":223,"y":-661},{"x":224,"y":-661},{"x":225,"y":-661},{"x":226,"y":-661},{"x":227,"y":-661},{"x":228,"y":-661},{"x":229,"y":-661},{"x":230,"y":-661},{"x":231,"y":-661},{"x":232,"y":-661},{"x":233,"y":-661},{"x":234,"y":-661},{"x":235,"y":-661},{"x":236,"y":-661},{"x":237,"y":-661},{"x":238,"y":-661},{"x":239,"y":-661},{"x":240,"y":-661},{"x":241,"y":-661},{"x":242,"y":-661},{"x":243,"y":-661},{"x":244,"y":-661},{"x":245,"y":-661},{"x":246,"y":-661},{"x":247,"y":-661},{"x":248,"y":-661},{"x":249,"y":-661},{"x":250,"y":-661},{"x":251,"y":-561},{"x":252,"y":-461},{"x":253,"y":-361},{"x":254,"y":-261},{"x":255,"y":-161},{"x":256,"y":-61},{"x":257,"y":39},{"x":258,"y":139},{"x":259,"y":239},{"x":260,"y":339},{"x":261,"y":439},{"x":262,"y":539},{"x":263,"y":639},{"x":264,"y":739},{"x":265,"y":839},{"x":266,"y":939},{"x":267,"y":1039},{"x":268,"y":1139},{"x":269,"y":1239},{"x":270,"y":1339},{"x":271,"y":1439},{"x":272,"y":1539},{"x":273,"y":1639},{"x":274,"y":1739},{"x":275,"y":1839},{"x":276,"y":1939},{"x":277,"y":2039},{"x":278,"y":2139},{"x":279,"y":2239},{"x":280,"y":2339},{"x":281,"y":2439},{"x":282,"y":2539},{"x":283,"y":2639},{"x":284,"y":2739},{"x":285,"y":2839},{"x":286,"y":2939},{"x":287,"y":3039},{"x":288,"y":3139},{"x":289,"y":3239},{"x":290,"y":3339},{"x":291,"y":3439},{"x":292,"y":3539},{"x":293,"y":3639},{"x":294,"y":3739},{"x":295,"y":3839},{"x":296,"y":3939},{"x":297,"y":4039},{"x":298,"y":4139},{"x":299,"y":4239},{"x":300,"y":4339},{"x":301,"y":4439},{"x":302,"y":4539},{"x":303,"y":4639},{"x":304,"y":4739},{"x":305,"y":4839},{"x":306,"y":4939},{"x":307,"y":5039},{"x":308,"y":5139},{"x":309,"y":5239},{"x":310,"y":5339},{"x":311,"y":5439},{"x":312,"y":5539},{"x":313,"y":5639},{"x":314,"y":5739},{"x":315,"y":5839},{"x":316,"y":5939},{"x":317,"y":6039},{"x":318,"y":6139},{"x":319,"y":6239},{"x":320,"y":6339},{"x":321,"y":6439},{"x":322,"y":6539},{"x":323,"y":6639},{"x":324,"y":6739},{"x":325,"y":6839},{"x":326,"y":6939},{"x":327,"y":7039},{"x":328,"y":7139},{"x":329,"y":7239},{"x":330,"y":7339},{"x":331,"y":7439},{"x":332,"y":7539},{"x":333,"y":7639},{"x":334,"y":7739},{"x":335,"y":7839},{"x":336,"y":7939},{"x":337,"y":8039},{"x":338,"y":8139},{"x":339,"y":8239},{"x":340,"y":8339},{"x":341,"y":8439},{"x":342,"y":8539},{"x":343,"y":8639},{"x":344,"y":8739},{"x":345,"y":8839},{"x":346,"y":8939},{"x":347,"y":9039},{"x":348,"y":9139},{"x":349,"y":9239},{"x":350,"y":9339},{"x":351,"y":9439},{"x":352,"y":9539},{"x":353,"y":9639},{"x":354,"y":9739},{"x":355,"y":9839},{"x":356,"y":9939},{"x":357,"y":10039},{"x":358,"y":10139},{"x":359,"y":10239},{"x":360,"y":10339},{"x":361,"y":10439},{"x":362,"y":10539},{"x":363,"y":10639},{"x":364,"y":10739},{"x":365,"y":10839},{"x":366,"y":10939},{"x":367,"y":11039},{"x":368,"y":11139},{"x":369,"y":11239},{"x":370,"y":11339},{"x":371,"y":11439},{"x":372,"y":11539},{"x":373,"y":11639},{"x":374,"y":11739},{"x":375,"y":11839},{"x":376,"y":11939},{"x":377,"y":12039},{"x":378,"y":12139},{"x":379,"y":12239},{"x":380,"y":12339},{"x":381,"y":12439},{"x":382,"y":12539},{"x":383,"y":12639},{"x":384,"y":12739},{"x":385,"y":12839},{"x":386,"y":12939},{"x":387,"y":13039},{"x":388,"y":13139},{"x":389,"y":13239},{"x":390,"y":13339},{"x":391,"y":13439},{"x":392,"y":13539},{"x":393,"y":13639},{"x":394,"y":13739},{"x":395,"y":13839},{"x":396,"y":13939},{"x":397,"y":14039},{"x":398,"y":14139},{"x":399,"y":14239},{"x":400,"y":14339},{"x":401,"y":14439},{"x":402,"y":14539},{"x":403,"y":14639},{"x":404,"y":14739},{"x":405,"y":14839},{"x":406,"y":14939},{"x":407,"y":15039},{"x":408,"y":15139},{"x":409,"y":15239},{"x":410,"y":15339},{"x":411,"y":15439},{"x":412,"y":15539},{"x":413,"y":15639},{"x":414,"y":15739},{"x":415,"y":15839},{"x":416,"y":15939},{"x":417,"y":16039},{"x":418,"y":16139},{"x":419,"y":16239},{"x":420,"y":16339},{"x":421,"y":16439},{"x":422,"y":16539},{"x":423,"y":16639},{"x":424,"y":16739},{"x":425,"y":16839},{"x":426,"y":16939},{"x":427,"y":17039},{"x":428,"y":17139},{"x":429,"y":17239},{"x":430,"y":17339},{"x":431,"y":17439},{"x":432,"y":17539},{"x":433,"y":17639},{"x":434,"y":17739},{"x":435,"y":17839},{"x":436,"y":17939},{"x":437,"y":18039},{"x":438,"y":18139},{"x":439,"y":18239},{"x":440,"y":18339},{"x":441,"y":18439},{"x":442,"y":18539},{"x":443,"y":18639},{"x":444,"y":18739},{"x":445,"y":18839},{"x":446,"y":18939},{"x":447,"y":19039},{"x":448,"y":19139},{"x":449,"y":19239},{"x":450,"y":19339},{"x":451,"y":19439},{"x":452,"y":19539},{"x":453,"y":19639},{"x":454,"y":19739},{"x":455,"y":19839},{"x":456,"y":19939},{"x":457,"y":20039},{"x":458,"y":20139},{"x":459,"y":20239},{"x":460,"y":20339},{"x":461,"y":20439},{"x":462,"y":20539},{"x":463,"y":20639},{"x":464,"y":20739},{"x":465,"y":20839},{"x":466,"y":20939},{"x":467,"y":21039},{"x":468,"y":21139},{"x":469,"y":21239},{"x":470,"y":21339},{"x":471,"y":21439},{"x":472,"y":21539},{"x":473,"y":21639},{"x":474,"y":21739},{"x":475,"y":21839},{"x":476,"y":21939},{"x":477,"y":22039},{"x":478,"y":22139},{"x":479,"y":22239},{"x":480,"y":22339},{"x":481,"y":22439},{"x":482,"y":22539},{"x":483,"y":22639},{"x":484,"y":22739},{"x":485,"y":22839},{"x":486,"y":22939},{"x":487,"y":23039},{"x":488,"y":23139},{"x":489,"y":23239},{"x":490,"y":23339},{"x":491,"y":23439},{"x":492,"y":23539},{"x":493,"y":23639},{"x":494,"y":23739},{"x":495,"y":23839},{"x":496,"y":23939},{"x":497,"y":24039},{"x":498,"y":24139},{"x":499,"y":24239},{"x":500,"y":24339}],"ask_pl_data":{"max_profit":1.7976931348623157e+308,"max_loss":661,"break_even_point":256.61}}}]`
)

func TestOnPOSTAnalysis(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/analysis", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	OnPOSTAnalysis(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}

	responseBody := strings.TrimSpace(w.Body.String())

	// replace random uuid with static UUID for testing purposes
	responseBody = strings.Replace(responseBody, responseBody[8:44], staticUUID, 1)

	if responseBody != expectedResponseBody {
		t.Errorf("Expected response body %v, got %v", expectedResponseBody, responseBody)
	}
}
