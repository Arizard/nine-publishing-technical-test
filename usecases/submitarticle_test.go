package usecases

import (
	"log"
	"fmt"
	"github.com/arizard/nine-publishing-technical-test/infrastructure"
	
	"testing"

	"github.com/arizard/nine-publishing-technical-test/entities"
)

func TestSubmitArticle_Execute(t *testing.T) {
	type fields struct {
		ArticleRepository entities.ArticleRepository
		ArticleData       map[string]interface{}
		Response          *ResponseCollector
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "correct input",
			fields: fields{
				infrastructure.NewInMemoryArticleRepository(),
				map[string]interface{}{
					"id": "1",
					"title": "latest science shows that potato chips are better for you than sugar",
					"date" : "2016-09-22",
					"body" : "some text, potentially containing simple markup about how potato chips are great",
					"tags" : []interface{}{
						"health", 
						"fitness", 
						"science",
					},
				},
				&(ResponseCollector{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("%v", tt.fields.ArticleData["tags"])
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
			
		})
	}
}
