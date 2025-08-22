package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ContactsvalidateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["marketId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("marketId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DomainsContactsBulk
		
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
		url := fmt.Sprintf("%s/v1/domains/contacts/validate%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Private-Label-Id"]; ok {
			req.Header.Set("X-Private-Label-Id", fmt.Sprintf("%v", val))
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

func CreateContactsvalidateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_domains_contacts_validate",
		mcp.WithDescription("Validate the request body using the Domain Contact Validation Schema for specified domains."),
		mcp.WithNumber("X-Private-Label-Id", mcp.Description("PrivateLabelId to operate as, if different from JWT")),
		mcp.WithString("marketId", mcp.Description("MarketId in which the request is being made, and for which responses should be localized")),
		mcp.WithObject("contactAdmin", mcp.Description("")),
		mcp.WithObject("contactBilling", mcp.Description("")),
		mcp.WithObject("contactPresence", mcp.Description("")),
		mcp.WithObject("contactRegistrant", mcp.Description("")),
		mcp.WithObject("contactTech", mcp.Description("")),
		mcp.WithArray("domains", mcp.Required(), mcp.Description("Input parameter: An array of domain names to be validated against. Alternatively, you can specify the extracted tlds. However, full domain names are required if the tld is `uk`")),
		mcp.WithString("entityType", mcp.Description("Input parameter: Canadian Presence Requirement (CA)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ContactsvalidateHandler(cfg),
	}
}
