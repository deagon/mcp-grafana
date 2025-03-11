package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/server"

	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/models"
	mcpgrafana "github.com/grafana/mcp-grafana"
)

type SearchDashboardsParams struct {
	Query string `json:"query" jsonschema:"description=The query to search for"`
}

func searchDashboards(ctx context.Context, args SearchDashboardsParams) (models.HitList, error) {
	c := mcpgrafana.GrafanaClientFromContext(ctx)
	params := search.NewSearchParamsWithContext(ctx)
	if args.Query != "" {
		params.SetQuery(&args.Query)
	}
	search, err := c.Search.Search(params)
	if err != nil {
		return nil, fmt.Errorf("search dashboards for %+v: %w", c, err)
	}
	return search.Payload, nil
}

var SearchDashboards = mcpgrafana.MustTool(
	"search_dashboards",
	"Search for dashboards",
	searchDashboards,
)

func AddSearchTools(mcp *server.MCPServer) {
	SearchDashboards.Register(mcp)
}
