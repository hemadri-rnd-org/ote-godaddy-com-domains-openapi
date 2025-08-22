package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["statuses"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("statuses=%v", val))
		}
		if val, ok := args["statusGroups"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("statusGroups=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["marker"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("marker=%v", val))
		}
		if val, ok := args["includes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includes=%v", val))
		}
		if val, ok := args["modifiedDate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modifiedDate=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/domains%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Shopper-Id"]; ok {
			req.Header.Set("X-Shopper-Id", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []DomainSummary
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateListTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_domains",
		mcp.WithDescription("Retrieve a list of Domains for the specified Shopper"),
		mcp.WithString("X-Shopper-Id", mcp.Description("Shopper ID whose domains are to be retrieved")),
		mcp.WithArray("statuses", mcp.Description("Only include results with `status` value in the specified set")),
		mcp.WithArray("statusGroups", mcp.Description("Only include results with `status` value in any of the specified groups")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of domains to return")),
		mcp.WithString("marker", mcp.Description("Marker Domain to use as the offset in results")),
		mcp.WithArray("includes", mcp.Description("Optional details to be included in the response")),
		mcp.WithString("modifiedDate", mcp.Description("Only include results that have been modified since the specified date")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListHandler(cfg),
	}
}
