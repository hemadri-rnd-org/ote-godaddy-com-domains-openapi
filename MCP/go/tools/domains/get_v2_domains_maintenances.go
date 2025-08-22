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

func Get_v2_domains_maintenancesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["modifiedAtAfter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modifiedAtAfter=%v", val))
		}
		if val, ok := args["startsAtAfter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startsAtAfter=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/domains/maintenances%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Request-Id"]; ok {
			req.Header.Set("X-Request-Id", fmt.Sprintf("%v", val))
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
		var result models.Maintenance
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

func CreateGet_v2_domains_maintenancesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_domains_maintenances",
		mcp.WithDescription("Retrieve a list of upcoming system Maintenances"),
		mcp.WithString("X-Request-Id", mcp.Description("A client provided identifier for tracking this request.")),
		mcp.WithString("status", mcp.Description("Only include results with the selected `status` value.  Returns all results if omitted<br/><ul><li><strong style='margin-left: 12px;'>ACTIVE</strong> - The upcoming maintenance is active.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - The upcoming maintenance has been cancelled.</li></ul>")),
		mcp.WithString("modifiedAtAfter", mcp.Description("Only include results with `modifiedAt` after the supplied date")),
		mcp.WithString("startsAtAfter", mcp.Description("Only include results with `startsAt` after the supplied date")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of results to return")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_domains_maintenancesHandler(cfg),
	}
}
