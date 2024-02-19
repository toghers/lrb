package share

import (
	"encoding/json"

	"github.com/astaxie/beego/httplib"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/config"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"
)

func InitEs(c config.Config, resp []*goods.GoodsResp) error {

	for _, v := range resp {
		res := httplib.Post(c.Esurl + "inx_goons" + "/_doc/" + v.GoodsId)
		res.JSONBody(v)
		_, err := res.String()
		if err != nil {
			return err
		}
	}
	return nil

}

type TypeHg struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				GoodsName  string `json:"GoodsName"`
				GoodsPrice string `json:"GoodsPrice"`
				GoodsNum   string `json:"GoodsNum"`
				GoodsRef   string `json:"GoodsRef"`
				GoodsId    string `json:"GoodsId"`
				GoodsImage string `json:"GoodsImage"`
			} `json:"_source"`
			Highlight struct {
				GoodsName []string `json:"GoodsName"`
			} `json:"highlight"`
		} `json:"hits"`
	} `json:"hits"`
}

type Source struct {
	GoodsName  string `json:"GoodsName"`
	GoodsPrice string `json:"GoodsPrice"`
	GoodsNum   string `json:"GoodsNum"`
	GoodsRef   string `json:"GoodsRef"`
	GoodsId    string `json:"GoodsId"`
	GoodsImage string `json:"GoodsImage"`
}

type Hg struct {
	Query struct {
		Match struct {
			GoodsName string `json:"GoodsName"`
		} `json:"match"`
	} `json:"query"`
	Highlight struct {
		Fields struct {
			GoodsName struct {
				PreTags  string `json:"pre_tags"`
				PostTags string `json:"post_tags"`
			} `json:"GoodsName"`
		} `json:"fields"`
	} `json:"highlight"`
	From int `json:"from"`
	Size int `json:"size"`
}

type Participle struct { //分词
	Settings struct {
		Index struct {
			Analysis struct {
				Analyzer struct {
					Default struct {
						Type string `json:"type"`
					} `json:"default"`
				} `json:"analyzer"`
			} `json:"analysis"`
		} `json:"index"`
	} `json:"settings"`
	Mappings struct {
		Properties struct {
			Description struct {
				Type     string `json:"type"`
				Analyzer string `json:"analyzer"`
			} `json:"description"`
		} `json:"properties"`
	} `json:"mappings"`
}

func HighSearch(c config.Config, context string, page, limit int) (error, []types.Goodinfo) {

	//var p Participle
	//p.Settings.Index.Analysis.Analyzer.Default.Type = "ik_max_word"
	//p.Mappings.Properties.Description.Analyzer = "ik_max_word"
	//p.Mappings.Properties.Description.Type = "text"
	//

	var h Hg
	h.Query.Match.GoodsName = context
	h.Highlight.Fields.GoodsName.PreTags = "<span color='red'>"
	h.Highlight.Fields.GoodsName.PostTags = "</span>"
	h.From = (page - 1) * limit
	h.Size = limit

	var g []types.Goodinfo
	res := httplib.Get(c.Esurl + "inx_goons" + "/_search")
	res.JSONBody(h)
	s, err := res.String()
	if err != nil {
		return err, []types.Goodinfo{}
	}

	var t TypeHg
	json.Unmarshal([]byte(s), &t)

	for _, v := range t.Hits.Hits {
		g = append(g, types.Goodinfo{
			GoodsName:  v.Highlight.GoodsName[0],
			GoodsPrice: v.Source.GoodsPrice,
			GoodsNum:   v.Source.GoodsNum,
			GoodsRef:   v.Source.GoodsRef,
			GoodsId:    v.Source.GoodsId,
			GoodsImage: v.Source.GoodsImage,
		})
	}

	return nil, g
}
