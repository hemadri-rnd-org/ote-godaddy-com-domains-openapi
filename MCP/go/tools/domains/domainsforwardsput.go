package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DomainsforwardsputHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.DomainForwardingCreate
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v2/customers/%s/domains/forwards/%s", cfg.BaseURL, customerId, fqdn)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateDomainsforwardsputTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_v2_customers_customerId_domains_forwards_fqdn",
		mcp.WithDescription("Modify the forwarding information for the given fqdn"),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.")),
		mcp.WithString("fqdn", mcp.Required(), mcp.Description("The fully qualified domain name whose forwarding details are to be modified.")),
		mcp.WithString("url", mcp.Required(), mcp.Description("Input parameter: Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)")),
		mcp.WithObject("mask", mcp.Description("")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DomainsforwardsputHandler(cfg),
	}
}
