package main

import (
	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	tools_domains "github.com/mcp-server/mcp-server/tools/domains"
	tools_v1 "github.com/mcp-server/mcp-server/tools/v1"
	tools_notifications "github.com/mcp-server/mcp-server/tools/notifications"
	tools_actions "github.com/mcp-server/mcp-server/tools/actions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferoutacceptTool(cfg),
		tools_v1.CreateGetagreementTool(cfg),
		tools_v1.CreateSuggestTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_register_validateTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_registerTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_renewTool(cfg),
		tools_domains.CreateGet_v2_customers_customerid_domains_domainTool(cfg),
		tools_v1.CreateUpdateTool(cfg),
		tools_v1.CreateCancelTool(cfg),
		tools_v1.CreateGetTool(cfg),
		tools_notifications.CreateGet_v2_customers_customerid_domains_notifications_schemas_typeTool(cfg),
		tools_actions.CreateDelete_v2_customers_customerid_domains_domain_actions_typeTool(cfg),
		tools_actions.CreateGet_v2_customers_customerid_domains_domain_actions_typeTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferoutrejectTool(cfg),
		tools_v1.CreateTransferinTool(cfg),
		tools_v1.CreatePurchaseprivacyTool(cfg),
		tools_notifications.CreateGet_v2_customers_customerid_domains_notifications_optinTool(cfg),
		tools_notifications.CreatePut_v2_customers_customerid_domains_notifications_optinTool(cfg),
		tools_v1.CreateContactsvalidateTool(cfg),
		tools_v1.CreateRenewTool(cfg),
		tools_domains.CreateGet_v2_customers_customerid_domains_register_schema_tldTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferincancelTool(cfg),
		tools_v1.CreateListTool(cfg),
		tools_v1.CreateCancelprivacyTool(cfg),
		tools_v1.CreateRecordreplacetypeTool(cfg),
		tools_v1.CreateSchemaTool(cfg),
		tools_v1.CreateVerifyemailTool(cfg),
		tools_v1.CreatePurchaseTool(cfg),
		tools_actions.CreateGet_v2_customers_customerid_domains_domain_actionsTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferinacceptTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferoutTool(cfg),
		tools_domains.CreateGet_v2_domains_maintenancesTool(cfg),
		tools_notifications.CreateGet_v2_customers_customerid_domains_notificationsTool(cfg),
		tools_v1.CreateValidateTool(cfg),
		tools_notifications.CreatePost_v2_customers_customerid_domains_notifications_notificationid_acknowledgeTool(cfg),
		tools_domains.CreateGet_v2_domains_maintenances_maintenanceidTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferinretryTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_redeemTool(cfg),
		tools_v1.CreateUpdatecontactsTool(cfg),
		tools_v1.CreateTldsTool(cfg),
		tools_v1.CreateRecordaddTool(cfg),
		tools_v1.CreateRecordreplaceTool(cfg),
		tools_domains.CreatePost_v2_customers_customerid_domains_domain_transferinrestartTool(cfg),
		tools_v1.CreateRecorddeletetypenameTool(cfg),
		tools_v1.CreateRecordgetTool(cfg),
		tools_v1.CreateRecordreplacetypenameTool(cfg),
		tools_v1.CreateAvailableTool(cfg),
		tools_v1.CreateAvailablebulkTool(cfg),
		tools_domains.CreateDomainsforwardsdeleteTool(cfg),
		tools_domains.CreateDomainsforwardsgetTool(cfg),
		tools_domains.CreateDomainsforwardspostTool(cfg),
		tools_domains.CreateDomainsforwardsputTool(cfg),
	}
}
