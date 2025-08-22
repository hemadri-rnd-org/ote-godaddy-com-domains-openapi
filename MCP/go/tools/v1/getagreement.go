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

func GetagreementHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["tlds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tlds=%v", val))
		}
		if val, ok := args["privacy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("privacy=%v", val))
		}
		if val, ok := args["forTransfer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("forTransfer=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/domains/agreements%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Market-Id"]; ok {
			req.Header.Set("X-Market-Id", fmt.Sprintf("%v", val))
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
		var result []LegalAgreement
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

func CreateGetagreementTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_domains_agreements",
		mcp.WithDescription("Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons"),
		mcp.WithString("X-Market-Id", mcp.Description("Unique identifier of the Market used to retrieve/translate Legal Agreements")),
		mcp.WithArray("tlds", mcp.Required(), mcp.Description("list of TLDs whose legal agreements are to be retrieved")),
		mcp.WithBoolean("privacy", mcp.Required(), mcp.Description("Whether or not privacy has been requested")),
		mcp.WithBoolean("forTransfer", mcp.Description("Whether or not domain tranfer has been requested")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetagreementHandler(cfg),
	}
}
