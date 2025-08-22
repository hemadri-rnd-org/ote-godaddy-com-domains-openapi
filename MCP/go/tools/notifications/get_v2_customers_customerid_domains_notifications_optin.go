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

func Get_v2_customers_customerid_domains_notifications_optinHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerIdVal, ok := args["customerId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customerId"), nil
		}
		customerId, ok := customerIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customerId"), nil
		}
		url := fmt.Sprintf("%s/v2/customers/%s/domains/notifications/optIn", cfg.BaseURL, customerId)
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
		var result []DomainNotification
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

func CreateGet_v2_customers_customerid_domains_notifications_optinTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_customers_customerId_domains_notifications_optIn",
		mcp.WithDescription("Retrieve a list of notification types that are opted in"),
		mcp.WithString("X-Request-Id", mcp.Description("A client provided identifier for tracking this request.")),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_customers_customerid_domains_notifications_optinHandler(cfg),
	}
}
