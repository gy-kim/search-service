package data

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

const (
	mapping = `
	{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"_doc":{
				"properties":{
					"id":{
						"type":"keyword"
					},
					"product_type":{
						"type":"keyword"
					},
					"brand":{
						"type":"keyword"
					},
					"name":{
						"type":"keyword"
					}
				}
			}
		}
	}
	`
)

// ref) https://github.com/olivere/elastic/wiki/QueryDSL
func createAndPopulateIndex(client *elastic.Client) error {
	ctx := context.Background()
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		fmt.Println("exists")
		_, err = client.DeleteIndex(indexName).Do(ctx)
		if err != nil {
			return err
		}
	}

	if _, err = client.CreateIndex(indexName).Body(mapping).Do(ctx); err != nil {
		fmt.Printf("failed to CreateIndex. err: (%v)\n", err)
		return err
	}

	products := []Product{
		{ID: "1_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "1_product_1_name"},
		{ID: "2_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "2_product_2_name"},
		{ID: "3_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "3_product_3_name"},
		{ID: "4_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "4_product_4_name"},
		{ID: "5_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "5_product_5_name"},
		{ID: "6_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "6_product_6_name"},
		{ID: "7_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "7_product_7_name"},
		{ID: "8_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "8_product_8_name"},
		{ID: "9_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "9_product_9_name"},
		{ID: "10_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "10_product_10_name"},
		{ID: "11_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "11_product_11_name"},
		{ID: "12_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "12_product_12_name"},

		{ID: "13_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "13_product_name"},
		{ID: "14_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "14_product_name"},
		{ID: "15_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "15_product_name"},
		{ID: "16_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "16_product_name"},
		{ID: "17_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "17_product_name"},
		{ID: "18_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "18_product_name"},
		{ID: "19_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "19_product_name"},
		{ID: "20_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "20_product_name"},
		{ID: "21_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "21_product_name"},
	}

	for _, product := range products {
		_, err := client.Index().
			Index(indexName).
			Type(typeName).
			BodyJson(product).
			Do(ctx)
		if err != nil {
			fmt.Printf("failed to index. err:(%v)", err)
			return err
		}
	}
	_, err = client.Flush(indexName).WaitIfOngoing(true).Do(ctx)

	return nil
}
