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

func AvailableHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["domain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("domain=%v", val))
		}
		if val, ok := args["checkType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("checkType=%v", val))
		}
		if val, ok := args["forTransfer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("forTransfer=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/domains/available%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

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
		var result models.DomainAvailableResponse
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

func CreateAvailableTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_domains_available",
		mcp.WithDescription("Determine whether or not the specified domain is available for purchase"),
		mcp.WithString("domain", mcp.Required(), mcp.Description("Domain name whose availability is to be checked")),
		mcp.WithString("checkType", mcp.Description("Optimize for time ('FAST') or accuracy ('FULL')")),
		mcp.WithBoolean("forTransfer", mcp.Description("Whether or not to include domains available for transfer. If set to True, checkType is ignored")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AvailableHandler(cfg),
	}
}
