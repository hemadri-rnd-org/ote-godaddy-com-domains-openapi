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

func RecorddeletetypenameHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		domainVal, ok := args["domain"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: domain"), nil
		}
		domain, ok := domainVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: domain"), nil
		}
		typeVal, ok := args["type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: type"), nil
		}
		type, ok := typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: type"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		url := fmt.Sprintf("%s/v1/domains/%s/records/%s/%s", cfg.BaseURL, domain, type, name)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result map[string]interface{}
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

func CreateRecorddeletetypenameTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v1_domains_domain_records_type_name",
		mcp.WithDescription("Delete all DNS Records for the specified Domain with the specified Type and Name"),
		mcp.WithString("X-Shopper-Id", mcp.Description("Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com")),
		mcp.WithString("domain", mcp.Required(), mcp.Description("Domain whose DNS Records are to be deleted")),
		mcp.WithString("type", mcp.Required(), mcp.Description("DNS Record Type for which DNS Records are to be deleted")),
		mcp.WithString("name", mcp.Required(), mcp.Description("DNS Record Name for which DNS Records are to be deleted")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RecorddeletetypenameHandler(cfg),
	}
}
