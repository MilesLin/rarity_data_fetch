package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {

	fetchAssets()
}

func fetchAssets() {
	var result []Asset
	// https://api.opensea.io/api/v1/assets?asset_contract_address=0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb&limit=20&token_ids=10000&token_ids=7804&token_ids=7523
	for i := 0; i < 10000; i += 20 {
		var sTokenIds string
		for j := 0; j < 20; j++ {
			sTokenIds = sTokenIds + fmt.Sprintf("&token_ids=%d", i+j)
		}
		fmt.Println(sTokenIds)
		<-time.After(150 * time.Millisecond)
		url := fmt.Sprintf("https://api.opensea.io/api/v1/assets?asset_contract_address=0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb&limit=20%s", sTokenIds)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(strconv.Itoa(i)+" got an error", err.Error())
			continue
		}
		var s Face
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&s)

		result = append(result, s.Assets...)
		resp.Body.Close()
	}

	resp, err := http.Get("https://api.opensea.io/api/v1/assets?asset_contract_address=0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb&limit=1&token_ids=10000")

	if err != nil {
		fmt.Println("10000 got an error", err.Error())
	}

	var s Face
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&s)
	result = append(result, s.Assets...)

	bytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println("10000 marshal error: ", err.Error())
	}
	fmt.Println("count: ", len(result))
	if err := ioutil.WriteFile("result.json", bytes, fs.ModePerm); err != nil {
		fmt.Println("Write file error: ", err.Error())
	}
}

type Asset struct {
	ID                   int         `json:"id"`
	TokenID              string      `json:"token_id"`
	NumSales             int         `json:"num_sales"`
	BackgroundColor      interface{} `json:"background_color"`
	ImageURL             string      `json:"image_url"`
	ImagePreviewURL      string      `json:"image_preview_url"`
	ImageThumbnailURL    interface{} `json:"image_thumbnail_url"`
	ImageOriginalURL     interface{} `json:"image_original_url"`
	AnimationURL         interface{} `json:"animation_url"`
	AnimationOriginalURL interface{} `json:"animation_original_url"`
	Name                 string      `json:"name"`
	Description          interface{} `json:"description"`
	ExternalLink         string      `json:"external_link"`
	AssetContract        struct {
		Address                     string      `json:"address"`
		AssetContractType           string      `json:"asset_contract_type"`
		CreatedDate                 string      `json:"created_date"`
		Name                        string      `json:"name"`
		NftVersion                  string      `json:"nft_version"`
		OpenseaVersion              interface{} `json:"opensea_version"`
		Owner                       interface{} `json:"owner"`
		SchemaName                  string      `json:"schema_name"`
		Symbol                      string      `json:"symbol"`
		TotalSupply                 interface{} `json:"total_supply"`
		Description                 string      `json:"description"`
		ExternalLink                string      `json:"external_link"`
		ImageURL                    string      `json:"image_url"`
		DefaultToFiat               bool        `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
		PayoutAddress               interface{} `json:"payout_address"`
	} `json:"asset_contract"`
	Permalink  string `json:"permalink"`
	Collection struct {
		BannerImageURL          string      `json:"banner_image_url"`
		ChatURL                 interface{} `json:"chat_url"`
		CreatedDate             string      `json:"created_date"`
		DefaultToFiat           bool        `json:"default_to_fiat"`
		Description             string      `json:"description"`
		DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
		DiscordURL              string      `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalURL                 string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageURL            string      `json:"featured_image_url"`
		Hidden                      bool        `json:"hidden"`
		SafelistRequestStatus       string      `json:"safelist_request_status"`
		ImageURL                    string      `json:"image_url"`
		IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
		LargeImageURL               string      `json:"large_image_url"`
		MediumUsername              interface{} `json:"medium_username"`
		Name                        string      `json:"name"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               interface{} `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramURL                 interface{} `json:"telegram_url"`
		TwitterUsername             string      `json:"twitter_username"`
		InstagramUsername           interface{} `json:"instagram_username"`
		WikiURL                     interface{} `json:"wiki_url"`
	} `json:"collection"`
	Decimals      interface{} `json:"decimals"`
	TokenMetadata string      `json:"token_metadata"`
	Owner         struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"owner"`
	SellOrders interface{} `json:"sell_orders"`
	Creator    struct {
		User          interface{} `json:"user"`
		ProfileImgURL string      `json:"profile_img_url"`
		Address       string      `json:"address"`
		Config        string      `json:"config"`
	} `json:"creator"`
	Traits                  []Traits    `json:"traits"`
	LastSale                interface{} `json:"last_sale"`
	TopBid                  interface{} `json:"top_bid"`
	ListingDate             interface{} `json:"listing_date"`
	IsPresale               bool        `json:"is_presale"`
	TransferFeePaymentToken interface{} `json:"transfer_fee_payment_token"`
	TransferFee             interface{} `json:"transfer_fee"`
}

type Face struct {
	Assets []Asset `json:"assets"`
}

type Traits struct {
	TraitType   string      `json:"trait_type"`
	Value       string      `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}
