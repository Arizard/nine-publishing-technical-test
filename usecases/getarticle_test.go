package usecases

import (
	"github.com/arizard/nine-publishing-technical-test/infrastructure"
	"github.com/icrowley/fake"
	"fmt"
	"testing"

	"github.com/arizard/nine-publishing-technical-test/entities"
)

func TestGetArticle_Execute(t *testing.T) {
	type fields struct {
		ArticleRepository entities.ArticleRepository
		ArticleData       map[string]interface{}
		Response          *ResponseCollector
	}

	tests := []struct {
		name   string
		fields fields
	}{}

	for index := 0; index < 25; index++ {
		tempField := struct {
				name   string
				fields fields
			}{
			name: fmt.Sprintf("faked input %d", index),
			fields: fields{
				infrastructure.NewInMemoryArticleRepository(),
				map[string]interface{}{
					"id": string(index),
					"title": fake.Paragraph(),
					"date" : fmt.Sprintf("%d-%02d-%02d", fake.Year(2005, 2019), fake.MonthNum(), fake.Day()),
					"body" : fake.Paragraphs(),
					"tags" : []interface{}{
						fake.Word(), 
						fake.Word(), 
						fake.Word(),
					},
				},
				&(ResponseCollector{}),
			},
		}
		tests = append(tests, tempField)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := SubmitArticle{
				ArticleRepository: tt.fields.ArticleRepository,
				ArticleData:       tt.fields.ArticleData,
				Response:          tt.fields.Response,
			}
			uc.Execute()
			if uc.Response.Error != nil {
				
				t.Error(
					fmt.Sprintf("error occured in test: %s (%s)",
						uc.Response.Error.Name,
						uc.Response.Error.Description,
					),
				)
			}

			uc2 := GetArticle{
				ArticleRepository: tt.fields.ArticleRepository,
				ArticleID: tt.fields.ArticleData["id"].(string),
				Response: &(ResponseCollector{}),
			}
			uc2.Execute()

			if uc2.Response.Error != nil {
				t.Error(
					fmt.Sprintf("error occured in test: %s (%s)",
						uc.Response.Error.Name,
						uc.Response.Error.Description,
					),
				)
			}

			output := uc2.Response.Response.Body

			if output["article"].(entities.Article).ID != tt.fields.ArticleData["id"].(string) {
				t.Error(
					fmt.Sprintf("error occured in test: mismatched ID returned"),
				)
			}
			
		})
	}
}
