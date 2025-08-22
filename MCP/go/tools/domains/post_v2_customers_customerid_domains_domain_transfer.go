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

func Post_v2_customers_customerid_domains_domain_transferHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		domainVal, ok := args["domain"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: domain"), nil
		}
		domain, ok := domainVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: domain"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DomainTransferInV2
		
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
		url := fmt.Sprintf("%s/v2/customers/%s/domains/%s/transfer", cfg.BaseURL, customerId, domain)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreatePost_v2_customers_customerid_domains_domain_transferTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v2_customers_customerId_domains_domain_transfer",
		mcp.WithDescription("Purchase and start or restart transfer process"),
		mcp.WithString("X-Request-Id", mcp.Description("A client provided identifier for tracking this request.")),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.")),
		mcp.WithString("domain", mcp.Required(), mcp.Description("Domain to transfer in")),
		mcp.WithNumber("period", mcp.Description("Input parameter: Can be more than 1 but no more than 10 years total including current registration length")),
		mcp.WithBoolean("privacy", mcp.Description("Input parameter: Whether or not privacy has been requested")),
		mcp.WithBoolean("renewAuto", mcp.Description("Input parameter: Whether or not the domain should be configured to automatically renew")),
		mcp.WithString("authCode", mcp.Required(), mcp.Description("Input parameter: Authorization code from registrar for transferring a domain")),
		mcp.WithObject("consent", mcp.Required(), mcp.Description("")),
		mcp.WithObject("contacts", mcp.Description("")),
		mcp.WithString("identityDocumentId", mcp.Description("Input parameter: Unique identifier of the identify document that the user wants to associate with the domain being transferred in. This is required only if the gaining registry has a requirement for an approved identity document")),
		mcp.WithObject("metadata", mcp.Description("Input parameter: The domain eligibility data fields as specified by GET /v2/customers/{customerId}/domains/register/schema/{tld}")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_v2_customers_customerid_domains_domain_transferHandler(cfg),
	}
}
