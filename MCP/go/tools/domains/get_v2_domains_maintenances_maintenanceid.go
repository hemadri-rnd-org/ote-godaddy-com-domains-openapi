package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_v2_domains_maintenances_maintenanceidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		maintenanceIdVal, ok := args["maintenanceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: maintenanceId"), nil
		}
		maintenanceId, ok := maintenanceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: maintenanceId"), nil
		}
		url := fmt.Sprintf("%s/v2/domains/maintenances/%s", cfg.BaseURL, maintenanceId)
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
		var result models.MaintenanceDetail
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

func CreateGet_v2_domains_maintenances_maintenanceidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_domains_maintenances_maintenanceId",
		mcp.WithDescription("Retrieve the details for an upcoming system Maintenances"),
		mcp.WithString("X-Request-Id", mcp.Description("A client provided identifier for tracking this request.")),
		mcp.WithString("maintenanceId", mcp.Required(), mcp.Description("The identifier for the system maintenance")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_domains_maintenances_maintenanceidHandler(cfg),
	}
}
