package data

import (
	"net/http"
	"os"

	"golang.frontdoorhome.com/software/protos/go/coreproductpb"
	"google.golang.org/protobuf/proto"
)

type Environment int

const (
	DEV Environment = iota
	TEST
	LOCAL
)

func GetEnvName(d Environment) string {
	return [...]string{"Dev", "Test", "Local"}[d]
}

func GetEnvURL(env Environment) string {
	var URL string
	switch env {
	case DEV:
		URL = os.Getenv("DEV_URL")
	case TEST:
		URL = os.Getenv("TEST_URL")

	case LOCAL:
		URL = os.Getenv("LOCAL_URL")
	}
	return URL
}

type APIDetails struct {
	Name       string
	Path       string
	Request    proto.Message
	Response   proto.Message
	Ignore     bool
	Env        Environment
	MethodType string
}

var APIInfo = []APIDetails{
	{
		Name: "Get all Templates",
		Path: "core-product/product-workbench/product-templates",
		Request: &coreproductpb.ProductTemplatesRequest{
			ProductTemplateID:     "2e9c6091-2eae-432a-b000-de975d318d69",
			IncludeAttributeRules: false,
			IncludeFeatureRules:   true,
			IncludeVersionHistory: true,
			IncludeAuditTrail:     false,
		},
		Response:   &coreproductpb.ProductTemplatesResponse{},
		Ignore:     true,
		Env:        LOCAL,
		MethodType: http.MethodGet,
	},
	{
		Path: "core-product/skus",
		Request: &coreproductpb.SKUByIDRequest{
			ProductID: "236a1d58-600d-11ed-9b6a-0242ac120002",
		},
		Response:   &coreproductpb.SKUByIDResponse{},
		Ignore:     true,
		Env:        DEV,
		MethodType: http.MethodGet,
	},
	{
		Name: "Get Products",
		Path: "core-product/product-workbench/product-templates/products",
		Request: &coreproductpb.GetProductsByTemplateIDRequest{
			ProductTemplateID: "a5599d54-9011-4da8-9a02-47cf5846e8af",
		},
		Response:   &coreproductpb.GetProductsByTemplateIDResponse{},
		Ignore:     true,
		Env:        DEV,
		MethodType: http.MethodGet,
	},
	{
		Name: "Get Products V2",
		Path: "core-product/v2/products",
		Request: &coreproductpb.GetProductsRequest{
			Id: &coreproductpb.GetProductsRequest_StarPVID{
				StarPVID: "96148",
			},
		},
		Response:   &coreproductpb.GetProductsResponse{},
		Ignore:     true,
		Env:        DEV,
		MethodType: http.MethodGet,
	},
	{
		Name: "Get Products V1",
		Path: "core-product/products",
		Request: &coreproductpb.ProductsRequest{
			ProductID: "96068",
		},
		Response:   &coreproductpb.ProductsResponse{},
		Ignore:     true,
		Env:        TEST,
		MethodType: http.MethodGet,
	},
	{
		Name: "Get Attribute-Values",
		Path: "core-product/product-workbench/attribute-values",
		Request: &coreproductpb.AttributeValuesRequest{
			AttributeName: "PRODUCT_TYPE",
		},
		Response:   &coreproductpb.AttributeValuesResponse{},
		Ignore:     true,
		Env:        DEV,
		MethodType: http.MethodGet,
	},
	{
		Name:    "GenerateProduct With Put Template",
		Path:    "core-product/product-workbench/product-templates/versions",
		Request: &coreproductpb.ReplaceProductTemplateVersionRequest{},
		Ignore:  true,
	},
	{
		Name: "GenerateProduct With Put Template",
		Path: "core-product/product-availability-by-location",
		Request: &coreproductpb.ProductAvailabilityRequest{
			SalesSourceId: "153612",
			CustomerType:  1,
			ProductOptions: &coreproductpb.ProductOptions{
				SalesChannel: 1,
			},
			PropertyFeatures: &coreproductpb.PropertyFeatures{
				Zip:   "38125",
				State: "TN",
				SquareFootageRange: &coreproductpb.SquareFootageRange{
					Min: 12,
					Max: 23,
				},
				DwellingType: 1,
			},
		},
		Ignore: false,
	},
}
