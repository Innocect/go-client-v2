package main

import (
	"clientv1/httpclientutils"
	"log"

	"github.com/joho/godotenv"
	"golang.frontdoorhome.com/software/protos/go/coreproductpb"
)

var (
	data = []*coreproductpb.ProductTemplateAttributeRule{
		{
			Name: "SalesChannel",
			ValueRules: []*coreproductpb.ValueRule{
				{
					Value: &coreproductpb.TypedValue{
						TypedValue: &coreproductpb.TypedValue_StringValue{
							StringValue: "RE",
						},
					},
					InclusionConditions: []*coreproductpb.AttributeCondition{
						{
							Name:     "Brand",
							Operator: coreproductpb.ConditionalOperator_CO_IS_IN_LIST,
							Value: &coreproductpb.TypedValue{
								TypedValue: &coreproductpb.TypedValue_StringList_{
									StringList: &coreproductpb.TypedValue_StringList{
										List: []string{"AHS", "HSA"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "Brand",
			ValueRules: []*coreproductpb.ValueRule{
				{
					Value: &coreproductpb.TypedValue{
						TypedValue: &coreproductpb.TypedValue_StringValue{
							StringValue: "AHS",
						},
					},
				},
				{
					Value: &coreproductpb.TypedValue{
						TypedValue: &coreproductpb.TypedValue_StringValue{
							StringValue: "HSA",
						},
					},
				},
			},
		},
	}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// csvutils.FillAPIScenarios()
	httpclientutils.TestGetApiCall()

	// fmt.Print(protojson.Format(data[0]))
	// fmt.Print(protojson.Format(data[1]))
}
