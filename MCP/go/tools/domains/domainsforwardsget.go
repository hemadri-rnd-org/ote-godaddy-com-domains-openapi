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

func DomainsforwardsgetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		fqdnVal, ok := args["fqdn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fqdn"), nil
		}
		fqdn, ok := fqdnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fqdn"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["includeSubs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeSubs=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/customers/%s/domains/forwards/%s%s", cfg.BaseURL, customerId, fqdn, queryString)
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
		var result []DomainForwarding
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

func CreateDomainsforwardsgetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_customers_customerId_domains_forwards_fqdn",
		mcp.WithDescription("Retrieve the forwarding information for the given fqdn"),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.")),
		mcp.WithString("fqdn", mcp.Required(), mcp.Description("The fully qualified domain name whose forwarding details are to be retrieved.")),
		mcp.WithBoolean("includeSubs", mcp.Description("Optionally include all sub domains if the fqdn specified is a domain and not a sub domain.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DomainsforwardsgetHandler(cfg),
	}
}
