package cacherepositories

import (
	"encoding/json"
	"fmt"

	cacherepositories "github.com/go-web-templates/api/internal/application/interfaces/cache-repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksCacheRepository,
		fx.As(new(cacherepositories.BooksCacheRepository)),
	),
)

func mustSerializeToJson(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if (err != nil) {
		panic(err)
	}

	return string(jsonData)
}


func mustDeserializeFromJson(data string, target interface{}) {
	err := json.Unmarshal([]byte(data), target)
	if err != nil {
		panic(err)
	}
}

func addPrefix(prefix string, value string) string {
	return fmt.Sprintf("%s:%s", prefix, value)
}
