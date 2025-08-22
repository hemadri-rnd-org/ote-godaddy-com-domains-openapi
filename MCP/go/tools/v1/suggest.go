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

func SuggestHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["sources"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sources=%v", val))
		}
		if val, ok := args["tlds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tlds=%v", val))
		}
		if val, ok := args["lengthMax"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lengthMax=%v", val))
		}
		if val, ok := args["lengthMin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lengthMin=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["waitMs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("waitMs=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/domains/suggest%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result []DomainSuggestion
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

func CreateSuggestTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_domains_suggest",
		mcp.WithDescription("Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper's purchase history"),
		mcp.WithString("X-Shopper-Id", mcp.Description("Shopper ID for which the suggestions are being generated")),
		mcp.WithString("query", mcp.Description("Domain name or set of keywords for which alternative domain names will be suggested")),
		mcp.WithString("country", mcp.Description("Two-letter ISO country code to be used as a hint for target region<br/><br/>\nNOTE: These are sample values, there are many\n<a href=\"http://www.iso.org/iso/country_codes.htm\">more</a>")),
		mcp.WithString("city", mcp.Description("Name of city to be used as a hint for target region")),
		mcp.WithArray("sources", mcp.Description("Sources to be queried<br/><br/><ul>\n<li><strong>CC_TLD</strong> - Varies the TLD using Country Codes</li>\n<li><strong>EXTENSION</strong> - Varies the TLD</li>\n<li><strong>KEYWORD_SPIN</strong> - Identifies keywords and then rotates each one</li>\n<li><strong>PREMIUM</strong> - Includes variations with premium prices</li></ul>")),
		mcp.WithArray("tlds", mcp.Description("Top-level domains to be included in suggestions<br/><br/>\nNOTE: These are sample values, there are many\n<a href=\"http://www.godaddy.com/tlds/gtld.aspx#domain_search_form\">more</a>")),
		mcp.WithNumber("lengthMax", mcp.Description("Maximum length of second-level domain")),
		mcp.WithNumber("lengthMin", mcp.Description("Minimum length of second-level domain")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of suggestions to return")),
		mcp.WithNumber("waitMs", mcp.Description("Maximum amount of time, in milliseconds, to wait for responses\nIf elapses, return the results compiled up to that point")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SuggestHandler(cfg),
	}
}
