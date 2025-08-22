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

func TransferinHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.DomainTransferIn
		
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
		url := fmt.Sprintf("%s/v1/domains/%s/transfer", cfg.BaseURL, domain)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.DomainPurchaseResponse
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

func CreateTransferinTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_domains_domain_transfer",
		mcp.WithDescription("Purchase and start or restart transfer process"),
		mcp.WithString("X-Shopper-Id", mcp.Description("The Shopper to whom the domain should be transfered")),
		mcp.WithString("domain", mcp.Required(), mcp.Description("Domain to transfer in")),
		mcp.WithObject("contactAdmin", mcp.Description("")),
		mcp.WithObject("contactBilling", mcp.Description("")),
		mcp.WithObject("contactTech", mcp.Description("")),
		mcp.WithBoolean("privacy", mcp.Description("Input parameter: Whether or not privacy has been requested")),
		mcp.WithString("authCode", mcp.Required(), mcp.Description("Input parameter: Authorization code from registrar for transferring a domain")),
		mcp.WithObject("contactRegistrant", mcp.Description("")),
		mcp.WithNumber("period", mcp.Description("Input parameter: Can be more than 1 but no more than 10 years total including current registration length")),
		mcp.WithBoolean("renewAuto", mcp.Description("Input parameter: Whether or not the domain should be configured to automatically renew")),
		mcp.WithObject("consent", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TransferinHandler(cfg),
	}
}
