package structs

import (
	"time"
)

// Product is riachuelo json response
type Product struct {
	Score          float64   `json:"_score"`
	ChMaxPrice01   string    `json:"ch_max_price_0_1"`
	MinPrice01     string    `json:"min_price_0_1"`
	SmallImage     string    `json:"small_image"`
	ActivationDate time.Time `json:"activation_date"`
	ClasseCatalogo string    `json:"classe_catalogo"`
	PlpHoverImages string    `json:"plp_hover_images"`
	MaxPrice01     string    `json:"max_price_0_1"`
	Marca          string    `json:"marca"`
	Dco            string    `json:"dco"`
	Tamanho        []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"tamanho"`
	AttributeSetID            string `json:"attribute_set_id"`
	Manga                     string `json:"manga"`
	LinhaCatalogo             string `json:"linha_catalogo"`
	ShowVisenzeSimilarProduct string `json:"show_visenze_similar_product"`
	CategoriaRaizValue        string `json:"categoria_raiz_value"`
	Sku                       string `json:"sku"`
	Padronagem                string `json:"padronagem"`
	PadronagemValue           string `json:"padronagem_value"`
	TypeID                    string `json:"type_id"`
	Gm                        string `json:"gm"`
	CorSimples                []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"cor_simples"`
	MarcaValue   string `json:"marca_value"`
	CreationDate string `json:"creation_date"`
	EntityID     string `json:"entity_id"`
	Gola         string `json:"gola"`
	URLKey       string `json:"url_key"`
	ChMinPrice01 string `json:"ch_min_price_0_1"`
	GeneralColor []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"general_color"`
	Name               string `json:"name"`
	EnableProductBadge string `json:"enable_product_badge"`
}
