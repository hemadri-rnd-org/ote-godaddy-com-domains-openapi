package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DomainContactsCreateV2 represents the DomainContactsCreateV2 schema from the OpenAPI specification
type DomainContactsCreateV2 struct {
	Registrant ContactDomainCreate `json:"registrant,omitempty"`
	Registrantid string `json:"registrantId,omitempty"` // Unique identifier of the contact that the user wants to use for the domain registrant contact. This can be specified instead of the `registrant` property.
	Tech ContactDomainCreate `json:"tech,omitempty"`
	Techid string `json:"techId,omitempty"` // Unique identifier of the contact that the user wants to use for the domain tech contact. This can be specified instead of the `tech` property.
	Admin ContactDomainCreate `json:"admin,omitempty"`
	Adminid string `json:"adminId,omitempty"` // Unique identifier of the contact that the user wants to use for the domain admin contact. This can be specified instead of the `admin` property.
	Billing ContactDomainCreate `json:"billing,omitempty"`
	Billingid string `json:"billingId,omitempty"` // Unique identifier of the contact that the user wants to use for the domain billing contact. This can be specified instead of the `billing` property.
}

// ActionReason represents the ActionReason schema from the OpenAPI specification
type ActionReason struct {
	Message string `json:"message,omitempty"` // Human-readable, English description of the code
	Code string `json:"code"` // Short identifier, suitable for indicating the reason for the current status and how to handle within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
}

// JsonDataType represents the JsonDataType schema from the OpenAPI specification
type JsonDataType struct {
	Format string `json:"format,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	TypeField string `json:"type"`
}

// Maintenance represents the Maintenance schema from the OpenAPI specification
type Maintenance struct {
	Environment string `json:"environment"` // The environment on which the maintenance will be performed<br/><ul><li><strong style='margin-left: 12px;'>OTE</strong> - The Operational Testing Environment.</li><li><strong style='margin-left: 12px;'>PRODUCTION</strong> - The Live Production Environment.</li></ul>
	Summary string `json:"summary"` // A brief description of what is being performed
	TypeField string `json:"type"` // The type of maintenance being performed<br/><ul><li><strong style='margin-left: 12px;'>API</strong> - Programmatic Api components.</li><li><strong style='margin-left: 12px;'>REGISTRY</strong> - The underlying Registry providing the tld(s).</li><li><strong style='margin-left: 12px;'>UI</strong> - User Interface components.</li></ul>
	Endsat string `json:"endsAt"` // Date and time (UTC) when this maintenance will complete
	Modifiedat string `json:"modifiedAt"` // Date and time (UTC) when this maintenance was last modified
	Startsat string `json:"startsAt"` // Date and time (UTC) when this maintenance will start
	Status string `json:"status"` // The status of maintenance<br/><ul><li><strong style='margin-left: 12px;'>ACTIVE</strong> - The upcoming maintenance is active.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - The upcoming maintenance has been cancelled.</li></ul>
	Createdat string `json:"createdAt"` // Date and time (UTC) when this maintenance was created
	Maintenanceid string `json:"maintenanceId"` // The identifier for the system maintenance
	Reason string `json:"reason"` // The reason for the maintenance being performed<br/><ul><li><strong style='margin-left: 12px;'>EMERGENCY</strong> - Unexpected Emergency maintenance.</li><li><strong style='margin-left: 12px;'>PLANNED</strong> - Planned system maintenance.</li></ul>
	Tlds []string `json:"tlds,omitempty"` // List of tlds that are in maintenance. Generally only applies when `type` is REGISTRY
}

// ErrorDomainContactsValidate represents the ErrorDomainContactsValidate schema from the OpenAPI specification
type ErrorDomainContactsValidate struct {
	Fields []ErrorFieldDomainContactsValidate `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Stack []string `json:"stack,omitempty"` // Stack trace indicating where the error occurred.<br/>NOTE: This attribute <strong>MAY</strong> be included for Development and Test environments. However, it <strong>MUST NOT</strong> be exposed from OTE nor Production systems
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
}

// LegalAgreement represents the LegalAgreement schema from the OpenAPI specification
type LegalAgreement struct {
	Url string `json:"url,omitempty"` // URL to a page containing the legal agreement
	Agreementkey string `json:"agreementKey"` // Unique identifier for the legal agreement
	Content string `json:"content"` // Contents of the legal agreement, suitable for embedding
	Title string `json:"title"` // Title of the legal agreement
}

// DomainDnssec represents the DomainDnssec schema from the OpenAPI specification
type DomainDnssec struct {
	Maxsignaturelife int `json:"maxSignatureLife,omitempty"` // This specifies the validity period for the signature. The value is expressed in seconds. You can use any integer value larger than zero
	Publickey string `json:"publicKey,omitempty"` // Registries use this value to encrypt DS records. Decryption requires a matching public key
	Algorithm string `json:"algorithm"` // This identifies the cryptographic algorithm used to generate the signature<br/><ul><li><strong style='margin-left: 12px;'>RSAMD5</strong> - [01] DRSA/MD5 </li><li><strong style='margin-left: 12px;'>DSA</strong> - [03] DSA/SHA1</li><li><strong style='margin-left: 12px;'>RSASHA1</strong> - [05] RSA/SHA-1</li><li><strong style='margin-left: 12px;'>DSA_NSEC3_SHA1</strong> - [06] DSA-NSEC3-SHA1</li><li><strong style='margin-left: 12px;'>RSASHA1_NSEC3_SHA1</strong> - [07] RSASHA1-NSEC3-SHA1</li><li><strong style='margin-left: 12px;'>RSASHA256</strong> - [08] RSA/SHA-256</li><li><strong style='margin-left: 12px;'>RSASHA512</strong> - [10] RSA/SHA-512</li><li><strong style='margin-left: 12px;'>ECC_GOST</strong> - [12] GOST R 34.10-2001</li><li><strong style='margin-left: 12px;'>ECDSAP256SHA256</strong> - [13] ECDSA Curve P-256 with SHA-256</li><li><strong style='margin-left: 12px;'>ECDSAP384SHA384</strong> - [14] ECDSA Curve P-384 with SHA-384</li><li><strong style='margin-left: 12px;'>ED25519</strong> - [15] Ed25519</li><li><strong style='margin-left: 12px;'>ED448</strong> - [16] Ed448</li></ul>
	Digest string `json:"digest,omitempty"` // The digest is an alpha-numeric value
	Digesttype string `json:"digestType,omitempty"` // This identifies the algorithm used to construct the digest<br/><ul><li><strong style='margin-left: 12px;'>SHA1</strong> - [01] SHA-1</li><li><strong style='margin-left: 12px;'>SHA256</strong> - [02] SHA-256</li><li><strong style='margin-left: 12px;'>GOST</strong> - [03] GOST R 34.11-94</li><li><strong style='margin-left: 12px;'>SHA384</strong> - [04] SHA-384</li></ul>
	Flags string `json:"flags,omitempty"` // This identifies the key type; either a Zone-Signing Key or a Key-Signing Key<br/><ul><li><strong style='margin-left: 12px;'>ZSK</strong> - [256] Zone-Signing Key</li><li><strong style='margin-left: 12px;'>KSK</strong> - [257] Key-Signing Key</li></ul>
	Keytag int `json:"keyTag,omitempty"` // This is an integer value less than 65536 used to identify the DNSSEC record for the domain name.
}

// DomainAvailableResponse represents the DomainAvailableResponse schema from the OpenAPI specification
type DomainAvailableResponse struct {
	Available bool `json:"available"` // Whether or not the domain name is available
	Currency string `json:"currency,omitempty"` // Currency in which the `price` is listed. Only returned if tld is offered
	Definitive bool `json:"definitive"` // Whether or not the `available` answer has been definitively verified with the registry
	Domain string `json:"domain"` // Domain name
	Period int `json:"period,omitempty"` // Number of years included in the price. Only returned if tld is offered
	Price int `json:"price,omitempty"` // Price of the domain excluding taxes or fees. Only returned if tld is offered
}

// DomainTransferAuthCode represents the DomainTransferAuthCode schema from the OpenAPI specification
type DomainTransferAuthCode struct {
	Authcode string `json:"authCode"` // Authorization code for transferring the Domain
}

// JsonProperty represents the JsonProperty schema from the OpenAPI specification
type JsonProperty struct {
	TypeField string `json:"type"`
	Defaultvalue string `json:"defaultValue,omitempty"`
	Format string `json:"format,omitempty"`
	Minitems int `json:"minItems,omitempty"`
	Items map[string]interface{} `json:"items,omitempty"`
	Maxitems int `json:"maxItems,omitempty"`
	Maximum int `json:"maximum,omitempty"`
	Minimum int `json:"minimum,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Required bool `json:"required"`
}

// DomainAvailableBulk represents the DomainAvailableBulk schema from the OpenAPI specification
type DomainAvailableBulk struct {
	Domains []DomainAvailableResponse `json:"domains"` // Domain available response array
}

// DomainsContactsBulk represents the DomainsContactsBulk schema from the OpenAPI specification
type DomainsContactsBulk struct {
	Contactpresence Contact `json:"contactPresence,omitempty"`
	Contactregistrant Contact `json:"contactRegistrant,omitempty"`
	Contacttech Contact `json:"contactTech,omitempty"`
	Domains []string `json:"domains"` // An array of domain names to be validated against. Alternatively, you can specify the extracted tlds. However, full domain names are required if the tld is `uk`
	Entitytype string `json:"entityType,omitempty"` // Canadian Presence Requirement (CA)
	Contactadmin Contact `json:"contactAdmin,omitempty"`
	Contactbilling Contact `json:"contactBilling,omitempty"`
}

// DomainUpdate represents the DomainUpdate schema from the OpenAPI specification
type DomainUpdate struct {
	Locked bool `json:"locked,omitempty"` // Whether or not the domain should be locked to prevent transfers
	Nameservers []interface{} `json:"nameServers,omitempty"` // Fully-qualified domain names for Name Servers to associate with the domain
	Renewauto bool `json:"renewAuto,omitempty"` // Whether or not the domain should be configured to automatically renew
	Subaccountid string `json:"subaccountId,omitempty"` // Reseller subaccount shopperid who can manage the domain
	Consent ConsentDomainUpdate `json:"consent,omitempty"`
	Exposewhois bool `json:"exposeWhois,omitempty"` // Whether or not the domain contact details should be shown in the WHOIS
}

// ConsentV2 represents the ConsentV2 schema from the OpenAPI specification
type ConsentV2 struct {
	Registrypremiumpricing bool `json:"registryPremiumPricing,omitempty"` // Only required for hosted registrar if domain is premium. If true indicates that the `price` and `currency` listed are the registry premium price and currency for the domain
	Agreedat string `json:"agreedAt"` // Timestamp indicating when the end-user consented to these legal agreements
	Agreedby string `json:"agreedBy"` // Originating client IP address of the end-user's computer when they consented to these legal agreements
	Agreementkeys []string `json:"agreementKeys"` // Unique identifiers of the legal agreements to which the end-user has agreed, as returned from the/domains/agreements endpoint
	Claimtoken string `json:"claimToken,omitempty"` // The trademark claim token, only needed if the domain has an active trademark claim
	Currency string `json:"currency"` // Currency in which the `price` is listed
	Price int `json:"price"` // Price of the domain excluding taxes or fees. Please use GET /v1/domains/available to retrieve the price and currency for the domain
}

// TldSummary represents the TldSummary schema from the OpenAPI specification
type TldSummary struct {
	TypeField string `json:"type"` // Type of the top-level domain
	Name string `json:"name"` // Name of the top-level domain
}

// ConsentRenew represents the ConsentRenew schema from the OpenAPI specification
type ConsentRenew struct {
	Agreedat string `json:"agreedAt"` // Timestamp indicating when the end-user consented to these legal agreements
	Agreedby string `json:"agreedBy"` // Originating client IP address of the end-user's computer when they consented to these legal agreements
	Currency string `json:"currency"` // Currency in which the `price` is listed
	Price int `json:"price"` // Price of the domain excluding taxes or fees. Please use GET /v2/customers/{customerId}/domains/{domain} to retrieve the renewal price and currency for the domain
	Registrypremiumpricing bool `json:"registryPremiumPricing,omitempty"` // Only required for hosted registrar if domain is premium. If true indicates that the `price` and `currency` listed are the registry premium price and currency for the domain
}

// ContactDomain represents the ContactDomain schema from the OpenAPI specification
type ContactDomain struct {
	Exposewhois bool `json:"exposeWhois"` // Whether or not the contact details should be shown in the WHOIS
	Revision int `json:"_revision,omitempty"` // The current revision number of the contact.
	Addressmailing Address `json:"addressMailing"`
	Organization string `json:"organization,omitempty"`
	Encoding string `json:"encoding,omitempty"` // The encoding of the contact data<br/><ul><li><strong style='margin-left: 12px;'>ASCII</strong> - Data contains only ASCII characters that are not region or language specific</li><li><strong style='margin-left: 12px;'>UTF-8</strong> - Data contains characters that are specific to a region or language</li></ul>
	Deleted bool `json:"_deleted,omitempty"` // Flag indicating if the contact has been logically deleted in the system
	Fax string `json:"fax,omitempty"`
	Jobtitle string `json:"jobTitle,omitempty"`
	Phone string `json:"phone"`
	Namemiddle string `json:"nameMiddle,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The contact eligibility data fields as specified by GET /v2/customers/{customerId}/domains/contacts/schema/{tld}
	Tlds []string `json:"tlds,omitempty"` // The tlds that this contact can be assigned to
	Createdat string `json:"_createdAt,omitempty"` // Timestamp indicating when the contact was created
	Modifiedat string `json:"_modifiedAt,omitempty"` // Timestamp indicating when the contact was last modified
	Contactid string `json:"contactId,omitempty"` // Unique identifier for this Contact
	Namelast string `json:"nameLast"`
	Namefirst string `json:"nameFirst"`
	Email string `json:"email"`
}

// DomainContacts represents the DomainContacts schema from the OpenAPI specification
type DomainContacts struct {
	Contactbilling Contact `json:"contactBilling,omitempty"`
	Contactregistrant Contact `json:"contactRegistrant"`
	Contacttech Contact `json:"contactTech,omitempty"`
	Contactadmin Contact `json:"contactAdmin,omitempty"`
}

// DomainSuggestion represents the DomainSuggestion schema from the OpenAPI specification
type DomainSuggestion struct {
	Domain string `json:"domain"` // Suggested domain name
}

// DomainDetailV2 represents the DomainDetailV2 schema from the OpenAPI specification
type DomainDetailV2 struct {
	Domain string `json:"domain"` // Name of the domain
	Domainid string `json:"domainId"` // Unique identifier for this Domain
	Renewdeadline string `json:"renewDeadline"` // Date the domain must renew on
	Status string `json:"status"` // The current status of the domain<br/><ul><li><strong style='margin-left: 12px;'>ACTIVE</strong> - Domain has been registered and is active.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - Domain has been cancelled by the user or system, and is not be reclaimable.</li><li><strong style='margin-left: 12px;'>DELETED_REDEEMABLE</strong> - Domain is deleted but is redeemable.</li><li><strong style='margin-left: 12px;'>EXPIRED</strong> - Domain has expired.</li><li><strong style='margin-left: 12px;'>FAILED</strong> - Domain registration or transfer error.</li><li><strong style='margin-left: 12px;'>LOCKED_REGISTRAR</strong> - Domain is locked at the registrar - this is usually the result of a spam, abuse, etc.</li><li><strong style='margin-left: 12px;'>PARKED</strong> - Domain has been parked.</li><li><strong style='margin-left: 12px;'>HELD_REGISTRAR</strong> - Domain is held at the registrar and cannot be transferred or modified - this is usually the result of a dispute.</li><li><strong style='margin-left: 12px;'>OWNERSHIP_CHANGED</strong> - Domain has been moved to another account.</li><li><strong style='margin-left: 12px;'>PENDING_TRANSFER</strong> - Domain transfer has been requested and is pending the transfer process.</li><li><strong style='margin-left: 12px;'>PENDING_REGISTRATION</strong> - Domain is pending setup at the registry.</li><li><strong style='margin-left: 12px;'>REPOSSESSED</strong> - Domain has been confiscated - this is usually the result of a chargeback, fraud, abuse, etc.).</li><li><strong style='margin-left: 12px;'>SUSPENDED</strong> - Domain is in violation and has been suspended.</li><li><strong style='margin-left: 12px;'>TRANSFERRED</strong> - Domain has been transferred to another registrar.</li></ul>
	Locked bool `json:"locked"` // Whether or not the domain is locked to prevent transfers
	Deletedat string `json:"deletedAt,omitempty"` // Date and time when this domain was deleted
	Dnssecrecords []DomainDnssec `json:"dnssecRecords,omitempty"` // List of active DNSSEC records for this domain
	Createdat string `json:"createdAt"` // Date and time when this domain was created
	Holdregistrar bool `json:"holdRegistrar"` // Whether or not the domain is on-hold by the registrar
	Renewauto bool `json:"renewAuto"` // Whether or not the domain is configured to automatically renew
	Transferawayeligibleat string `json:"transferAwayEligibleAt,omitempty"` // Date and time when this domain is eligible to transfer
	Expiresat string `json:"expiresAt,omitempty"` // Date and time when this domain will expire
	Hostnames []string `json:"hostnames,omitempty"` // Hostnames owned by the domain
	Registrarcreatedat string `json:"registrarCreatedAt,omitempty"` // Date and time when this domain was created by the registrar
	Transferprotected bool `json:"transferProtected"` // Whether or not the domain is protected from transfer
	Privacy bool `json:"privacy"` // Whether or not the domain has privacy protection
	Renewal RenewalDetails `json:"renewal,omitempty"`
	Subaccountid string `json:"subaccountId,omitempty"` // Reseller subaccount shopperid who can manage the domain
	Actions []Action `json:"actions,omitempty"` // List of current actions in progress for this domain
	Verifications VerificationsDomainV2 `json:"verifications,omitempty"`
	Nameservers []string `json:"nameServers"` // Fully-qualified domain names for DNS servers
	Registrystatuscodes []string `json:"registryStatusCodes,omitempty"` // The current registry status codes of the domain<br/><ul><li><strong style='margin-left: 12px;'>ADD_PERIOD</strong> - This grace period is provided after the initial registration of a domain name.</li><li><strong style='margin-left: 12px;'>AUTO_RENEW_PERIOD</strong> - This grace period is provided after a domain name registration period expires and is extended (renewed) automatically by the registry.</li><li><strong style='margin-left: 12px;'>CLIENT_DELETE_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to delete the domain.</li><li><strong style='margin-left: 12px;'>CLIENT_HOLD</strong> - This status code tells your domain's registry to not activate your domain in the DNS and as a consequence, it will not resolve.</li><li><strong style='margin-left: 12px;'>CLIENT_RENEW_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to renew your domain.</li><li><strong style='margin-left: 12px;'>CLIENT_TRANSFER_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to transfer the domain from your current registrar to another.</li><li><strong style='margin-left: 12px;'>CLIENT_UPDATE_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to update the domain.</li><li><strong style='margin-left: 12px;'>INACTIVE</strong> - This status code indicates that delegation information (name servers) has not been associated with your domain.</li><li><strong style='margin-left: 12px;'>OK</strong> - This is the standard status for a domain, meaning it has no pending operations or prohibitions.</li><li><strong style='margin-left: 12px;'>PENDING_CREATE</strong> - This status code indicates that a request to create your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_DELETE</strong> - This status code indicates that the domain is either in a redemption period if combined with either REDEMPTION_PERIOD or PENDING_RESTORE, if not combined with these, then indicates that the redemption period for the domain has ended and domain will be be purged and dropped from the registry database.</li><li><strong style='margin-left: 12px;'>PENDING_RENEW</strong> - This status code indicates that a request to renew your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_RESTORE</strong> - This status code indicates that your registrar has asked the registry to restore your domain that was in REDEMPTION_PERIOD status</li><li><strong style='margin-left: 12px;'>PENDING_TRANSFER</strong> - This status code indicates that a request to transfer your domain to a new registrar has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_UPDATE</strong> - This status code indicates that a request to update your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>REDEMPTION_PERIOD</strong> - This status code indicates that your registrar has asked the registry to delete your domain.</li><li><strong style='margin-left: 12px;'>RENEW_PERIOD</strong> - This grace period is provided after a domain name registration period is explicitly extended (renewed) by the registrar.</li><li><strong style='margin-left: 12px;'>SERVER_DELETE_PROHIBITED</strong> - This status code prevents your domain from being deleted. </li><li><strong style='margin-left: 12px;'>SERVER_HOLD</strong> - This status code is set by your domain's Registry Operator. Your domain is not activated in the DNS.</li><li><strong style='margin-left: 12px;'>SERVER_RENEW_PROHIBITED</strong> - This status code indicates your domain's Registry Operator will not allow your registrar to renew your domain.</li><li><strong style='margin-left: 12px;'>SERVER_TRANSFER_PROHIBITED</strong> - This status code prevents your domain from being transferred from your current registrar to another. </li><li><strong style='margin-left: 12px;'>SERVER_UPDATE_PROHIBITED</strong> - This status code locks your domain preventing it from being updated.</li><li><strong style='margin-left: 12px;'>TRANSFER_PERIOD</strong> - This grace period is provided after the successful transfer of a domain name from one registrar to another. </li></ul>
	Expirationprotected bool `json:"expirationProtected"` // Whether or not the domain is protected from expiration
	Modifiedat string `json:"modifiedAt,omitempty"` // Date and time when this domain was last modified
	Authcode string `json:"authCode"` // Authorization code for transferring the Domain
	Contacts DomainContactsV2 `json:"contacts"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Modifiedat string `json:"modifiedAt,omitempty"` // Timestamp indicating when the action was last modified
	Requestid string `json:"requestId,omitempty"` // A client provided identifier (via X-Request-Id header) used for tracking individual requests
	Startedat string `json:"startedAt,omitempty"` // Timestamp indicating when the action was started
	TypeField string `json:"type"` // The type of action being performed<br/><ul><li><strong style='margin-left: 12px;'>AUTH_CODE_PURCHASE</strong> - Request for an auth code for a .de domain via POST /v2/customers/{customerId}/domains/{domain}/purchaseAuthCode.</li><li><strong style='margin-left: 12px;'>AUTH_CODE_REGENERATE</strong> - Request to regenerate the authCode for a domain via POST /v2/customers/{customerId}/domains/{domain}/regenerateAuthCode</li><li><strong style='margin-left: 12px;'>AUTO_RENEWAL</strong> - A Domain Auto Renew is in progress.</li><li><strong style='margin-left: 12px;'>BACKORDER_PURCHASE</strong> - Request to purchase a domain backorder via POST /v2/customers/{customerId}/domains/backorders/purchase.</li><li><strong style='margin-left: 12px;'>BACKORDER_DELETE</strong> - Request to cancel the current domain backorder via DELETE /v2/customers/{customerId}/domains/backorders/{domain}.</li><li><strong style='margin-left: 12px;'>BACKORDER_UPDATE</strong> - Request update the current domain backorder via PATCH /v2/customers/{customerId}/domains/backorders/{domain}.</li><li><strong style='margin-left: 12px;'>CONTACT_CREATE</strong> - Request to create a contact via POST /v2/customers/{customerId}/domains/contacts.</li><li><strong style='margin-left: 12px;'>CONTACT_DELETE</strong> - Request to delete a contact via DELETE /v2/customers/{customerId}/domains/contacts/{contactId}</li><li><strong style='margin-left: 12px;'>CONTACT_UPDATE</strong> - Request to update a contact via PATCH /v2/customers/{customerId}/domains/contacts/{contactId}</li><li><strong style='margin-left: 12px;'>DNS_VERIFICATION</strong> - Domain requires zone file setup.</li><li><strong style='margin-left: 12px;'>DNSSEC_CREATE</strong> - Request to create DNSSEC record for the domain via PATCH /v2/customers/{customerId}/domains/{domain}/dnssecRecords.</li><li><strong style='margin-left: 12px;'>DNSSEC_DELETE</strong> - Request to delete DNSSEC record for the domain via DELETE /v2/customers/{customerId}/domains/{domain}/dnssecRecords.</li><li><strong style='margin-left: 12px;'>DOMAIN_DELETE</strong> - Request to delete the domain via DELETE /v2/customers/{customerId}/domains/{domain}</li><li><strong style='margin-left: 12px;'>DOMAIN_UPDATE</strong> - Request to update the domain via PATCH /v2/customers/{customerId}/domains/{domain}</li><li><strong style='margin-left: 12px;'>DOMAIN_UPDATE_CONTACTS</strong> -Request to update the domain contacts via PATCH /v2/customers/{customerId}/domains/{domain}/contacts</li><li><strong style='margin-left: 12px;'>DOMAIN_UPDATE_NAME_SERVERS</strong> - Request to update the domain name servers via PUT /v2/customers/{customerId}/domains/{domain}/nameServers</li><li><strong style='margin-left: 12px;'>EXPIRY</strong> - A Domain Expiration is in progress.</li><li><strong style='margin-left: 12px;'>HOST_CREATE</strong> - Request to create a hostname via PUT /v2/customers/{customerId}/domains/{domain}/hosts/{hostname}</li><li><strong style='margin-left: 12px;'>HOST_DELETE</strong> - Request to delete a hostname via DELETE /v2/customers/{customerId}/domains/{domain}/hosts/{hostname}</li><li><strong style='margin-left: 12px;'>ICANN_VERIFICATION</strong> - Domain requires registrant verification for Icann.</li><li><strong style='margin-left: 12px;'>PREMIUM</strong> - Premium Domain domain sale is in progress.</li><li><strong style='margin-left: 12px;'>PRIVACY_PURCHASE</strong> - Request to purchase privacy for a domain via POST /v2/customers/{customerId}/domains/{domain}/privacy/purchase</li><li><strong style='margin-left: 12px;'>PRIVACY_DELETE</strong> - Request to remove privacy from a domain via DELETE /v2/customers/{customerId}/domains/{domain}/privacy</li><li><strong style='margin-left: 12px;'>REDEEM</strong> - Request to redeem a domain via POST /v2/customers/{customerId}/domains/{domain}/redeem</li><li><strong style='margin-left: 12px;'>REGISTER</strong> - Request to register a domain via POST /v2/customers/{customerId}/domains/{domain}/register</li><li><strong style='margin-left: 12px;'>RENEW</strong> - Request to renew a domain via POST /v2/customers/{customerId}/domains/{domain}/renew</li><li><strong style='margin-left: 12px;'>RENEW_UNDO</strong> - Request to undo a renewal for a uk domain via POST /v2/customers/{customerId}/domains/{domain}/undoRenew</li><li><strong style='margin-left: 12px;'>TRADE</strong> - A domain trade request is in progress</li><li><strong style='margin-left: 12px;'>TRADE_CANCEL</strong> - Request to cancel a trade for a domain via POST /v2/customers/{customerId}/domains/{domain}/tradeCancel</li><li><strong style='margin-left: 12px;'>TRADE_PURCHASE</strong> - Request to purchase a trade for a domain via POST /v2/customers/{customerId}/domains/{domain}/tradePurchase</li><li><strong style='margin-left: 12px;'>TRADE_PURCHASE_AUTH_TEXT_MESSAGE</strong> - Request for a trade purchase text message for a domain via POST /v2/customers/{customerId}/domains/{domain}/tradePurchaseAuthorizationTextMessage</li><li><strong style='margin-left: 12px;'>TRADE_RESEND_AUTH_EMAIL</strong> - Request to resend the trade auth email message for a domain via POST /v2/customers/{customerId}/domains/{domain}/tradeResendAuthorizationEmail</li><li><strong style='margin-left: 12px;'>TRANSFER</strong> - Request to transfer a domain via POST /v2/customers/{customerId}/domains/{domain}/transfer</li><li><strong style='margin-left: 12px;'>TRANSFER_IN</strong> - A domain transfer in request is in progress.</li><li><strong style='margin-left: 12px;'>TRANSFER_IN_ACCEPT</strong> - Request to accept a domain transfer in via POST /v2/customers/{customerId}/domains/{domain}/transferInAccept</li><li><strong style='margin-left: 12px;'>TRANSFER_IN_CANCEL</strong> - Request to cancel a domain transfer via POST /v2/customers/{customerId}/domains/{domain}/transferInCancel</li><li><strong style='margin-left: 12px;'>TRANSFER_IN_RESTART</strong> - Request to restart a domain transfer in via POST /v2/customers/{customerId}/domains/{domain}/transferInRestart</li><li><strong style='margin-left: 12px;'>TRANSFER_IN_RETRY</strong> - Request to retry a domain transfer in via POST /v2/customers/{customerId}/domains/{domain}/transferInRetry</li><li><strong style='margin-left: 12px;'>TRANSFER_OUT</strong> - A domain transfer out request is in progress.</li><li><strong style='margin-left: 12px;'>TRANSFER_OUT_ACCEPT</strong> - Request to accept a transfer out request for a domain via POST /v2/customers/{customerId}/domains/{domain}/transferOutAccept</li><li><strong style='margin-left: 12px;'>TRANSFER_OUT_REJECT</strong> - Request to reject a transfer out request for a domain via POST /v2/customers/{customerId}/domains/{domain}/transferOutReject</li><li><strong style='margin-left: 12px;'>TRANSFER_OUT_REQUESTED</strong> - Request to transfer out for a domain (.de) via POST /v2/customers/{customerId}/domains/{domain}/transferOut</li><li><strong style='margin-left: 12px;'>TRANSIT</strong> - Request to transit a de or at domain at the registry via POST /v2/customers/{customerId}/domains/{domain}/transit</li></ul>
	Createdat string `json:"createdAt"` // Timestamp indicating when the action was created
	Origination string `json:"origination"` // The origination of the action<br/><ul><li><strong style='margin-left: 12px;'>USER</strong> - These are user requests.</li><li><strong style='margin-left: 12px;'>SYSTEM</strong> - These are system processing actions.</li></ul>
	Status string `json:"status"` // The current status of the action<br/><ul><li><strong style='margin-left: 12px;'>ACCEPTED</strong> - The action has been queued, processing has not started.</li><li><strong style='margin-left: 12px;'>AWAITING</strong> - The action is waiting on a user input.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - The action has been cancelled by the user.</li><li><strong style='margin-left: 12px;'>FAILED</strong> - An error occurred while the action was processing, no more processing will be performed.</li><li><strong style='margin-left: 12px;'>PENDING</strong> - The action is being processed.</li><li><strong style='margin-left: 12px;'>SUCCESS</strong> - The action has completed, no additional processing is required.</li></ul>
	Reason ActionReason `json:"reason,omitempty"`
	Completedat string `json:"completedAt,omitempty"` // Timestamp indicating when the action was completed
}

// DNSRecordCreateType represents the DNSRecordCreateType schema from the OpenAPI specification
type DNSRecordCreateType struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Port int `json:"port,omitempty"` // Service port (SRV only)
	Priority int `json:"priority,omitempty"` // Record priority (MX and SRV only)
	Protocol string `json:"protocol,omitempty"` // Service protocol (SRV only)
	Service string `json:"service,omitempty"` // Service type (SRV only)
	Ttl int `json:"ttl,omitempty"`
	Weight int `json:"weight,omitempty"` // Record weight (SRV only)
}

// DomainNotification represents the DomainNotification schema from the OpenAPI specification
type DomainNotification struct {
	Resourcetype string `json:"resourceType"` // The type of resource the notification relates to
	Status string `json:"status"` // The resulting status of the action.
	TypeField string `json:"type"` // The type of action the notification relates to
	Addedat string `json:"addedAt"` // The date the notification was added
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The notification data for the given type as specifed by GET /v2/customers/{customerId}/domains/notifications/schema
	Notificationid string `json:"notificationId"` // The notification ID to be used in POST /v2/customers/{customerId}/domains/notifications to acknowledge the notification
	Requestid string `json:"requestId,omitempty"` // A client provided identifier (via X-Request-Id header) indicating the request this notification is for
	Resource string `json:"resource"` // The resource the notification pertains to.
}

// DomainNotificationType represents the DomainNotificationType schema from the OpenAPI specification
type DomainNotificationType struct {
	TypeField string `json:"type"` // The notification type
}

// DomainTransferInV2 represents the DomainTransferInV2 schema from the OpenAPI specification
type DomainTransferInV2 struct {
	Identitydocumentid string `json:"identityDocumentId,omitempty"` // Unique identifier of the identify document that the user wants to associate with the domain being transferred in. This is required only if the gaining registry has a requirement for an approved identity document
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The domain eligibility data fields as specified by GET /v2/customers/{customerId}/domains/register/schema/{tld}
	Period int `json:"period,omitempty"` // Can be more than 1 but no more than 10 years total including current registration length
	Privacy bool `json:"privacy,omitempty"` // Whether or not privacy has been requested
	Renewauto bool `json:"renewAuto,omitempty"` // Whether or not the domain should be configured to automatically renew
	Authcode string `json:"authCode"` // Authorization code from registrar for transferring a domain
	Consent ConsentV2 `json:"consent"`
	Contacts DomainContactsCreateV2 `json:"contacts,omitempty"`
}

// ErrorField represents the ErrorField schema from the OpenAPI specification
type ErrorField struct {
	Message string `json:"message,omitempty"` // Human-readable, English description of the problem with the contents of the field
	Path string `json:"path"` // <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Pathrelated string `json:"pathRelated,omitempty"` // JSONPath referring to a field containing an error, which is referenced by `path`
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
}

// Consent represents the Consent schema from the OpenAPI specification
type Consent struct {
	Agreedby string `json:"agreedBy"` // Originating client IP address of the end-user's computer when they consented to these legal agreements
	Agreementkeys []string `json:"agreementKeys"` // Unique identifiers of the legal agreements to which the end-user has agreed, as returned from the/domains/agreements endpoint
	Agreedat string `json:"agreedAt"` // Timestamp indicating when the end-user consented to these legal agreements
}

// ContactDomainCreate represents the ContactDomainCreate schema from the OpenAPI specification
type ContactDomainCreate struct {
	Jobtitle string `json:"jobTitle,omitempty"`
	Namelast string `json:"nameLast"`
	Namemiddle string `json:"nameMiddle,omitempty"`
	Organization string `json:"organization,omitempty"`
	Phone string `json:"phone"`
	Namefirst string `json:"nameFirst"`
	Fax string `json:"fax,omitempty"`
	Encoding string `json:"encoding"` // The encoding of the contact data<br/><ul><li><strong style='margin-left: 12px;'>ASCII</strong> - Data contains only ASCII characters that are not region or language specific</li><li><strong style='margin-left: 12px;'>UTF-8</strong> - Data contains characters that are specific to a region or language</li></ul>
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The contact eligibility data fields as specified by GET /v2/customers/{customerId}/domains/contacts/schema/{tld}
	Addressmailing Address `json:"addressMailing"`
	Email string `json:"email"`
}

// PrivacyPurchase represents the PrivacyPurchase schema from the OpenAPI specification
type PrivacyPurchase struct {
	Consent Consent `json:"consent"`
}

// DomainAvailableError represents the DomainAvailableError schema from the OpenAPI specification
type DomainAvailableError struct {
	Path string `json:"path"` // <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Status int `json:"status"` // HTTP status code that would return for a single check
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Domain string `json:"domain"` // Domain name
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
}

// MaintenanceDetail represents the MaintenanceDetail schema from the OpenAPI specification
type MaintenanceDetail struct {
	Createdat string `json:"createdAt"` // Date and time (UTC) when this maintenance was created
	Endsat string `json:"endsAt"` // Date and time (UTC) when this maintenance will complete
	Reason string `json:"reason"` // The reason for the maintenance being performed<br/><ul><li><strong style='margin-left: 12px;'>EMERGENCY</strong> - Unexpected Emergency maintenance.</li><li><strong style='margin-left: 12px;'>PLANNED</strong> - Planned system maintenance.</li></ul>
	Startsat string `json:"startsAt"` // Date and time (UTC) when this maintenance will start
	Status string `json:"status"` // The status of maintenance<br/><ul><li><strong style='margin-left: 12px;'>ACTIVE</strong> - The upcoming maintenance is active.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - The upcoming maintenance has been cancelled.</li></ul>
	Environment string `json:"environment"` // The environment on which the maintenance will be performed<br/><ul><li><strong style='margin-left: 12px;'>OTE</strong> - The Operational Testing Environment.</li><li><strong style='margin-left: 12px;'>PRODUCTION</strong> - The Live Production Environment.</li></ul>
	Modifiedat string `json:"modifiedAt"` // Date and time (UTC) when this maintenance was last modified
	Tlds []string `json:"tlds,omitempty"` // List of tlds that are in maintenance. Generally only applies when `type` is REGISTRY
	Maintenanceid string `json:"maintenanceId"` // The identifier for the system maintenance
	Summary string `json:"summary"` // A brief description of what is being performed
	Systems []MaintenanceSystem `json:"systems,omitempty"` // List of systems that are impacted by the maintenance.
	TypeField string `json:"type"` // The type of maintenance being performed<br/><ul><li><strong style='margin-left: 12px;'>API</strong> - Programmatic Api components.</li><li><strong style='margin-left: 12px;'>REGISTRY</strong> - The underlying Registry providing the tld(s).</li><li><strong style='margin-left: 12px;'>UI</strong> - User Interface components.</li></ul>
}

// DomainTransferIn represents the DomainTransferIn schema from the OpenAPI specification
type DomainTransferIn struct {
	Contactregistrant Contact `json:"contactRegistrant,omitempty"`
	Period int `json:"period,omitempty"` // Can be more than 1 but no more than 10 years total including current registration length
	Renewauto bool `json:"renewAuto,omitempty"` // Whether or not the domain should be configured to automatically renew
	Consent Consent `json:"consent"`
	Contactadmin Contact `json:"contactAdmin,omitempty"`
	Contactbilling Contact `json:"contactBilling,omitempty"`
	Contacttech Contact `json:"contactTech,omitempty"`
	Privacy bool `json:"privacy,omitempty"` // Whether or not privacy has been requested
	Authcode string `json:"authCode"` // Authorization code from registrar for transferring a domain
}

// VerificationDomainName represents the VerificationDomainName schema from the OpenAPI specification
type VerificationDomainName struct {
	Status string `json:"status"` // Status of the domain name verification
}

// DomainRenewV2 represents the DomainRenewV2 schema from the OpenAPI specification
type DomainRenewV2 struct {
	Period int `json:"period,omitempty"` // Number of years to extend the Domain. Must not exceed maximum for TLD. When omitted, defaults to `period` specified during original purchase
	Consent ConsentRenew `json:"consent"`
	Expires string `json:"expires"` // Current date when this domain will expire
}

// RenewalDetails represents the RenewalDetails schema from the OpenAPI specification
type RenewalDetails struct {
	Currency string `json:"currency"` // Currency in which the `price` is listed
	Price int `json:"price"` // Price for the domain renewal excluding taxes or fees
	Renewable bool `json:"renewable,omitempty"` // Whether or not the domain is eligble for renewal based on status
}

// DNSRecord represents the DNSRecord schema from the OpenAPI specification
type DNSRecord struct {
	Port int `json:"port,omitempty"` // Service port (SRV only)
	Ttl int `json:"ttl,omitempty"`
	Name string `json:"name"`
	Priority int `json:"priority,omitempty"` // Record priority (MX and SRV only)
	Service string `json:"service,omitempty"` // Service type (SRV only)
	TypeField string `json:"type"`
	Weight int `json:"weight,omitempty"` // Record weight (SRV only)
	Data string `json:"data"`
	Protocol string `json:"protocol,omitempty"` // Service protocol (SRV only)
}

// DomainDetail represents the DomainDetail schema from the OpenAPI specification
type DomainDetail struct {
	Exposewhois bool `json:"exposeWhois,omitempty"` // Whether or not the domain contact details should be shown in the WHOIS
	Registrarcreatedat string `json:"registrarCreatedAt,omitempty"` // Date and time when this domain was created by the registrar
	Verifications VerificationsDomain `json:"verifications,omitempty"`
	Holdregistrar bool `json:"holdRegistrar"` // Whether or not the domain is on-hold by the registrar
	Subaccountid string `json:"subaccountId,omitempty"` // Reseller subaccount shopperid who can manage the domain
	Contactbilling Contact `json:"contactBilling"`
	Contactadmin Contact `json:"contactAdmin"`
	Domainid float64 `json:"domainId"` // Unique identifier for this Domain
	Transferawayeligibleat string `json:"transferAwayEligibleAt,omitempty"` // Date and time when this domain is eligible to transfer
	Nameservers []string `json:"nameServers"` // Fully-qualified domain names for DNS servers
	Contacttech Contact `json:"contactTech"`
	Expirationprotected bool `json:"expirationProtected"` // Whether or not the domain is protected from expiration
	Contactregistrant Contact `json:"contactRegistrant"`
	Deletedat string `json:"deletedAt,omitempty"` // Date and time when this domain was deleted
	Privacy bool `json:"privacy"` // Whether or not the domain has privacy protection
	Createdat string `json:"createdAt"` // Date and time when this domain was created
	Renewauto bool `json:"renewAuto"` // Whether or not the domain is configured to automatically renew
	Renewdeadline string `json:"renewDeadline"` // Date the domain must renew on
	Status string `json:"status"` // Processing status of the domain<br/><ul> <li><strong style='margin-left: 12px;'>ACTIVE</strong> - All is well</li> <li><strong style='margin-left: 12px;'>AWAITING*</strong> - System is waiting for the end-user to complete an action</li> <li><strong style='margin-left: 12px;'>CANCELLED*</strong> - Domain has been cancelled, and may or may not be reclaimable</li> <li><strong style='margin-left: 12px;'>CONFISCATED</strong> - Domain has been confiscated, usually for abuse, chargeback, or fraud</li> <li><strong style='margin-left: 12px;'>DISABLED*</strong> - Domain has been disabled</li> <li><strong style='margin-left: 12px;'>EXCLUDED*</strong> - Domain has been excluded from Firehose registration</li> <li><strong style='margin-left: 12px;'>EXPIRED*</strong> - Domain has expired</li> <li><strong style='margin-left: 12px;'>FAILED*</strong> - Domain has failed a required action, and the system is no longer retrying</li> <li><strong style='margin-left: 12px;'>HELD*</strong> - Domain has been placed on hold, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>LOCKED*</strong> - Domain has been locked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PARKED*</strong> - Domain has been parked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PENDING*</strong> - Domain is working its way through an automated workflow</li> <li><strong style='margin-left: 12px;'>RESERVED*</strong> - Domain is reserved, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>REVERTED</strong> - Domain has been reverted, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>SUSPENDED*</strong> - Domain has been suspended, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>TRANSFERRED*</strong> - Domain has been transferred out</li> <li><strong style='margin-left: 12px;'>UNKNOWN</strong> - Domain is in an unknown state</li> <li><strong style='margin-left: 12px;'>UNLOCKED*</strong> - Domain has been unlocked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UNPARKED*</strong> - Domain has been unparked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UPDATED*</strong> - Domain ownership has been transferred to another account</li> </ul>
	Expires string `json:"expires,omitempty"` // Date and time when this domain will expire
	Locked bool `json:"locked"` // Whether or not the domain is locked to prevent transfers
	Authcode string `json:"authCode"` // Authorization code for transferring the Domain
	Transferprotected bool `json:"transferProtected"` // Whether or not the domain is protected from transfer
	Domain string `json:"domain"` // Name of the domain
}

// ErrorFieldDomainContactsValidate represents the ErrorFieldDomainContactsValidate schema from the OpenAPI specification
type ErrorFieldDomainContactsValidate struct {
	Pathrelated string `json:"pathRelated,omitempty"` // JSONPath referring to the field on the object referenced by `path` containing an error
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Domains []string `json:"domains"` // An array of domain names the error is for. If tlds are specified in the request, `domains` will contain tlds. For example, if `domains` in request is ["test1.com", "test2.uk", "net"], and the field is invalid for com and net, then one of the `fields` in response will have ["test1.com", "net"] as `domains`
	Message string `json:"message,omitempty"` // Human-readable, English description of the problem with the contents of the field
	Path string `json:"path"` // 1) JSONPath referring to the field within the data containing an error<br/>or<br/>2) JSONPath referring to an object containing an error
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Namelast string `json:"nameLast"`
	Namemiddle string `json:"nameMiddle,omitempty"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Jobtitle string `json:"jobTitle,omitempty"`
	Namefirst string `json:"nameFirst"`
	Organization string `json:"organization,omitempty"`
	Addressmailing Address `json:"addressMailing"`
	Fax string `json:"fax,omitempty"`
}

// DomainForwardingMask represents the DomainForwardingMask schema from the OpenAPI specification
type DomainForwardingMask struct {
	Description string `json:"description,omitempty"` // A short description of your website to display in search engine results.
	Keywords string `json:"keywords,omitempty"` // A list of comma-separated keywords that describes the content and purpose of your website.
	Title string `json:"title,omitempty"` // Displays at the top of the browser window and in search results.
}

// DomainSummary represents the DomainSummary schema from the OpenAPI specification
type DomainSummary struct {
	Domainid float64 `json:"domainId"` // Unique identifier for this Domain
	Nameservers []string `json:"nameServers,omitempty"` // Fully-qualified domain names for DNS servers
	Registrarcreatedat string `json:"registrarCreatedAt,omitempty"` // Date and time when this domain was created by the registrar
	Expires string `json:"expires,omitempty"` // Date and time when this domain will expire
	Authcode string `json:"authCode,omitempty"` // Authorization code for transferring the Domain
	Contactregistrant Contact `json:"contactRegistrant"`
	Contactbilling Contact `json:"contactBilling,omitempty"`
	Expirationprotected bool `json:"expirationProtected"` // Whether or not the domain is protected from expiration
	Renewable bool `json:"renewable,omitempty"` // Whether or not the domain is eligble for renewal based on status
	Status string `json:"status"` // Processing status of the domain<br/><ul> <li><strong style='margin-left: 12px;'>ACTIVE</strong> - All is well</li> <li><strong style='margin-left: 12px;'>AWAITING*</strong> - System is waiting for the end-user to complete an action</li> <li><strong style='margin-left: 12px;'>CANCELLED*</strong> - Domain has been cancelled, and may or may not be reclaimable</li> <li><strong style='margin-left: 12px;'>CONFISCATED</strong> - Domain has been confiscated, usually for abuse, chargeback, or fraud</li> <li><strong style='margin-left: 12px;'>DISABLED*</strong> - Domain has been disabled</li> <li><strong style='margin-left: 12px;'>EXCLUDED*</strong> - Domain has been excluded from Firehose registration</li> <li><strong style='margin-left: 12px;'>EXPIRED*</strong> - Domain has expired</li> <li><strong style='margin-left: 12px;'>FAILED*</strong> - Domain has failed a required action, and the system is no longer retrying</li> <li><strong style='margin-left: 12px;'>HELD*</strong> - Domain has been placed on hold, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>LOCKED*</strong> - Domain has been locked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PARKED*</strong> - Domain has been parked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PENDING*</strong> - Domain is working its way through an automated workflow</li> <li><strong style='margin-left: 12px;'>RESERVED*</strong> - Domain is reserved, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>REVERTED</strong> - Domain has been reverted, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>SUSPENDED*</strong> - Domain has been suspended, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>TRANSFERRED*</strong> - Domain has been transferred out</li> <li><strong style='margin-left: 12px;'>UNKNOWN</strong> - Domain is in an unknown state</li> <li><strong style='margin-left: 12px;'>UNLOCKED*</strong> - Domain has been unlocked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UNPARKED*</strong> - Domain has been unparked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UPDATED*</strong> - Domain ownership has been transferred to another account</li> </ul>
	Transferawayeligibleat string `json:"transferAwayEligibleAt,omitempty"` // Date and time when this domain is eligible to transfer
	Transferprotected bool `json:"transferProtected"` // Whether or not the domain is protected from transfer
	Renewauto bool `json:"renewAuto"` // Whether or not the domain is configured to automatically renew
	Contacttech Contact `json:"contactTech,omitempty"`
	Exposewhois bool `json:"exposeWhois,omitempty"` // Whether or not the domain contact details should be shown in the WHOIS
	Createdat string `json:"createdAt"` // Date and time when this domain was created
	Locked bool `json:"locked"` // Whether or not the domain is locked to prevent transfers
	Renewdeadline string `json:"renewDeadline"` // Date the domain must renew on
	Deletedat string `json:"deletedAt,omitempty"` // Date and time when this domain was deleted
	Holdregistrar bool `json:"holdRegistrar"` // Whether or not the domain is on-hold by the registrar
	Privacy bool `json:"privacy"` // Whether or not the domain has privacy protection
	Contactadmin Contact `json:"contactAdmin,omitempty"`
	Domain string `json:"domain"` // Name of the domain
}

// MaintenanceSystem represents the MaintenanceSystem schema from the OpenAPI specification
type MaintenanceSystem struct {
	Name string `json:"name"` // The name of the system affected by the maintenance<br/><ul><li><strong style='margin-left: 12px;'>DOMAIN_CHECKS</strong> - Refers to domain availability checks.</li><li><strong style='margin-left: 12px;'>DOMAIN_MANAGEMENT</strong> - Refers to domain management options including various update options on the domain, contacts, records, etc.</li><li><strong style='margin-left: 12px;'>DOMAIN_REGISTRATION</strong> - Refers to domain registrations, renewals, transfers.</li><li><strong style='margin-left: 12px;'>DOMAIN_REGISTRATION_DATA</strong> - Refers to RDAP and WHOIS Service queries for domains.</li><li><strong style='margin-left: 12px;'>DOMAIN_RESOLUTION</strong> - Refers to DNS resolution for domains.</li><li><strong style='margin-left: 12px;'>RESELLER_ADMIN_PORTAL</strong> - Refers to Admin portals to manage the reseller account and settings.</li><li><strong style='margin-left: 12px;'>RESELLER_STOREFRONT</strong> - Refers to the Reseller Storefront features (Standard and Custom).</li></ul>
	Impact []string `json:"impact"` // The impact of the maintenance to the system<br/><ul><li><strong style='margin-left: 12px;'>DELAYED</strong> - This response generally applies to systems where the request is queued up and processed once the system is back online.</li><li><strong style='margin-left: 12px;'>DOWN</strong> - The system will be entirely offline; errors are expected.</li><li><strong style='margin-left: 12px;'>NON_AUTHORITATIVE</strong> - This response generally applies to DOMAIN_CHECKS and DOMAIN_MANAGEMENT `system` values where a cached answer will be supplied.</li><li><strong style='margin-left: 12px;'>PARTIAL</strong> - The system will experience partial feature outages; some errors are expected.</li></ul>
}

// VerificationsDomain represents the VerificationsDomain schema from the OpenAPI specification
type VerificationsDomain struct {
	Domainname VerificationDomainName `json:"domainName,omitempty"`
	Realname VerificationRealName `json:"realName,omitempty"`
}

// DomainContactsV2 represents the DomainContactsV2 schema from the OpenAPI specification
type DomainContactsV2 struct {
	Admin ContactDomain `json:"admin,omitempty"`
	Billing ContactDomain `json:"billing,omitempty"`
	Registrant ContactDomain `json:"registrant,omitempty"`
	Tech ContactDomain `json:"tech,omitempty"`
}

// DomainRedeemV2 represents the DomainRedeemV2 schema from the OpenAPI specification
type DomainRedeemV2 struct {
	Consent ConsentRedemption `json:"consent"`
}

// Domain represents the Domain schema from the OpenAPI specification
type Domain struct {
	Id float64 `json:"id,omitempty"`
}

// VerificationsDomainV2 represents the VerificationsDomainV2 schema from the OpenAPI specification
type VerificationsDomainV2 struct {
	Domainname string `json:"domainName,omitempty"` // Status of the verification of the domain name against a prohibited list maintained by the government
	Icann string `json:"icann,omitempty"` // Status of the Icann verification of domain registrant contact by completing email and/or phone verification<br/><ul><li><strong style='margin-left: 12px;'>COMPLETED</strong> - Icann verification has been completed.</li><li><strong style='margin-left: 12px;'>PENDING</strong> - Icann verification has not been completed.</li><li><strong style='margin-left: 12px;'>UNABLE_TO_RETRIEVE_STATUS</strong> - Icann verification not supported for specified TLD.</li></ul>
	Realname string `json:"realName,omitempty"` // Status of the real name verification of an identity by comparing registration data against government issued documents<br/><ul><li><strong style='margin-left: 12px;'>APPROVED</strong> - All is well</li><li><strong style='margin-left: 12px;'>PENDING</strong> - Real name verification is working its way through the workflow</li><li><strong style='margin-left: 12px;'>REJECTED_DOCUMENT_OUTDATED</strong> - Local government verification shows there is a newer version of your document. Upload the latest version of the document and retry real name verification</li><li><strong style='margin-left: 12px;'>REJECTED_EXPIRED_BUSINESS_LICENSE</strong> - Business license is expired</li><li><strong style='margin-left: 12px;'>REJECTED_EXPIRED_ORGANIZATION_CODE</strong> - Organization code certificate number has expired</li><li><strong style='margin-left: 12px;'>REJECTED_ILLEGIBLE_DOCUMENT_NAME</strong> - There isn’t a clear name on your uploaded document, please upload a different document to retry real name verification</li><li><strong style='margin-left: 12px;'>REJECTED_ILLEGIBLE_IDENTIFICATION</strong> - Registrant identification is not clear. Upload a better image to retry</li><li><strong style='margin-left: 12px;'>REJECTED_INCOMPLETE_IDENTIFICATION</strong> - Registrant identification is incomplete</li><li><strong style='margin-left: 12px;'>REJECTED_INCOMPLETE_REGISTRATION_LETTER</strong> - Registration letter is incomplete</li><li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_IDENTITY_CARD</strong> - Provided identity card is inconsistent with the identity card on record</li><li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_ORGANIZATION_CODE</strong> - Provided organization information is inconsistent with the results obtained using the submitted organization code</li><li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_REGISTRANT_NAME</strong> - Name on the registrant identification does not match the name in the system</li><li><strong style='margin-left: 12px;'>REJECTED_INVALID_BUSINESS_LICENSE_OR_ORGANIZATION_CODE</strong> - Your document contains an invalid business license or organization code certificate number</li><li><strong style='margin-left: 12px;'>REJECTED_INVALID_DOCUMENT</strong> - Document is invalid. Please upload another document to retry real name verification</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_BUSINESS_ID</strong> - Business id does not match the business id in the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_BUSINESS_NAME</strong> - Business name does not match the business name in the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_ID</strong> - Document id does not match the id in the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_NAME</strong> - Document name does not match the name in the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_TYPE</strong> - Document type does not match the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_INFO</strong> - The information provided for the registrant does not match the document</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_LOCALITY</strong> - Registrant region is overseas, but a local identity document was provided</li><li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_NAME</strong> - Registrant name has been changed, so the request must be resubmitted</li><li><strong style='margin-left: 12px;'>REJECTED_UNABLE_TO_OPEN</strong> - Registrant identification could not be opened. Please upload the document again to retry real name verification</li><li><strong style='margin-left: 12px;'>REJECTED_UNABLE_TO_VERIFY</strong> - Unable to initiate verification. Please upload the document again to retry real name verification</li><li><strong style='margin-left: 12px;'>REJECTED_UNKNOWN_ERROR</strong> - Document was rejected due to an unknown error. For more information, contact customer support</li><li><strong style='margin-left: 12px;'>UNABLE_TO_RETRIEVE_STATUS</strong> - Unable to retrieve status for the real name verification process. Retry, if this status persists, contact customer support</li></ul>
}

// DomainAvailableBulkMixed represents the DomainAvailableBulkMixed schema from the OpenAPI specification
type DomainAvailableBulkMixed struct {
	Domains []DomainAvailableResponse `json:"domains"` // Domain available response array
	Errors []DomainAvailableError `json:"errors,omitempty"` // Errors encountered while performing a domain available check
}

// RealNameValidation represents the RealNameValidation schema from the OpenAPI specification
type RealNameValidation struct {
	Status string `json:"status,omitempty"`
}

// DomainForwardingCreate represents the DomainForwardingCreate schema from the OpenAPI specification
type DomainForwardingCreate struct {
	Url string `json:"url"` // Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)
	Mask DomainForwardingMask `json:"mask,omitempty"`
	TypeField string `json:"type"` // The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>
}

// DomainPurchaseResponse represents the DomainPurchaseResponse schema from the OpenAPI specification
type DomainPurchaseResponse struct {
	Orderid int64 `json:"orderId"` // Unique identifier of the order processed to purchase the domain
	Total int `json:"total"` // Total cost of the domain and any selected add-ons
	Currency string `json:"currency,omitempty"` // Currency in which the `total` is listed
	Itemcount int `json:"itemCount"` // Number items included in the order
}

// DomainPurchaseV2 represents the DomainPurchaseV2 schema from the OpenAPI specification
type DomainPurchaseV2 struct {
	Period int `json:"period,omitempty"`
	Privacy bool `json:"privacy,omitempty"`
	Renewauto bool `json:"renewAuto,omitempty"`
	Consent ConsentV2 `json:"consent"`
	Contacts DomainContactsCreateV2 `json:"contacts,omitempty"`
	Domain string `json:"domain"` // For internationalized domain names with non-ascii characters, the domain name is converted to punycode before format and pattern validation rules are checked
	Metadata map[string]interface{} `json:"metadata,omitempty"` // The domain eligibility data fields as specified by GET /v2/customers/{customerId}/domains/register/schema/{tld}
	Nameservers []string `json:"nameServers,omitempty"`
}

// ConsentRedemption represents the ConsentRedemption schema from the OpenAPI specification
type ConsentRedemption struct {
	Agreedat string `json:"agreedAt"` // Timestamp indicating when the end-user consented to these legal agreements
	Agreedby string `json:"agreedBy"` // Originating client IP address of the end-user's computer when they consented to these legal agreements
	Currency string `json:"currency"` // Currency in which the `price` and `fee` are listed
	Fee int `json:"fee"` // Fee charged for the domain redemption. Please use GET /v2/customers/{customerId}/domains/{domain} to retrieve the redemption fee and currency for the domain
	Price int `json:"price"` // Price for the domain renewal (if domain renewal required for redemption). Please use GET /v2/customers/{customerId}/domains/{domain} to retrieve the redemption price and currency for the domain
}

// DomainPurchase represents the DomainPurchase schema from the OpenAPI specification
type DomainPurchase struct {
	Renewauto bool `json:"renewAuto,omitempty"`
	Nameservers []string `json:"nameServers,omitempty"`
	Period int `json:"period,omitempty"`
	Contactadmin Contact `json:"contactAdmin,omitempty"`
	Consent Consent `json:"consent"`
	Contactbilling Contact `json:"contactBilling,omitempty"`
	Domain string `json:"domain"` // For internationalized domain names with non-ascii characters, the domain name is converted to punycode before format and pattern validation rules are checked
	Contactregistrant Contact `json:"contactRegistrant,omitempty"`
	Contacttech Contact `json:"contactTech,omitempty"`
	Privacy bool `json:"privacy,omitempty"`
}

// ErrorLimit represents the ErrorLimit schema from the OpenAPI specification
type ErrorLimit struct {
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Retryaftersec int `json:"retryAfterSec"` // Number of seconds to wait before attempting a similar request
}

// JsonSchema represents the JsonSchema schema from the OpenAPI specification
type JsonSchema struct {
	Id string `json:"id"`
	Models map[string]interface{} `json:"models"`
	Properties map[string]interface{} `json:"properties"`
	Required []string `json:"required"`
}

// ConsentDomainUpdate represents the ConsentDomainUpdate schema from the OpenAPI specification
type ConsentDomainUpdate struct {
	Agreementkeys []string `json:"agreementKeys"` // Unique identifiers of the agreements to which the end-user has agreed, as required by the elements being updated<br/><ul><li><strong style='margin-left: 12px;'>EXPOSE_WHOIS</strong> - Required when the exposeWhois field is updated to true</li></ul>
	Agreedat string `json:"agreedAt"` // Timestamp indicating when the end-user consented to these agreements
	Agreedby string `json:"agreedBy"` // Originating client IP address of the end-user's computer when they consented to the agreements
}

// VerificationRealName represents the VerificationRealName schema from the OpenAPI specification
type VerificationRealName struct {
	Status string `json:"status"` // Status of the real name verification<br/><ul> <li><strong style='margin-left: 12px;'>APPROVED</strong> - All is well</li> <li><strong style='margin-left: 12px;'>PENDING</strong> - Real name verification is working its way through the workflow</li> <li><strong style='margin-left: 12px;'>REJECTED_DOCUMENT_OUTDATED</strong> - Local government verification shows there is a newer version of your document. Upload the latest version of the document and retry real name verification</li> <li><strong style='margin-left: 12px;'>REJECTED_EXPIRED_BUSINESS_LICENSE</strong> - Business license is expired</li> <li><strong style='margin-left: 12px;'>REJECTED_EXPIRED_ORGANIZATION_CODE</strong> - Organization code certificate number has expired</li> <li><strong style='margin-left: 12px;'>REJECTED_ILLEGIBLE_DOCUMENT_NAME</strong> - There isn’t a clear name on your uploaded document, please upload a different document to retry real name verification</li> <li><strong style='margin-left: 12px;'>REJECTED_ILLEGIBLE_IDENTIFICATION</strong> - Registrant identification is not clear. Upload a better image to retry</li> <li><strong style='margin-left: 12px;'>REJECTED_INCOMPLETE_IDENTIFICATION</strong> - Registrant identification is incomplete</li> <li><strong style='margin-left: 12px;'>REJECTED_INCOMPLETE_REGISTRATION_LETTER</strong> - Registration letter is incomplete</li> <li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_IDENTITY_CARD</strong> - Provided identity card is inconsistent with the identity card on record</li> <li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_ORGANIZATION_CODE</strong> - Provided organization information is inconsistent with the results obtained using the submitted organization code</li> <li><strong style='margin-left: 12px;'>REJECTED_INCONSISTENT_REGISTRANT_NAME</strong> - Name on the registrant identification does not match the name in the system</li> <li><strong style='margin-left: 12px;'>REJECTED_INVALID_BUSINESS_LICENSE_OR_ORGANIZATION_CODE</strong> - Your document contains an invalid business license or organization code certificate number</li> <li><strong style='margin-left: 12px;'>REJECTED_INVALID_DOCUMENT</strong> - Document is invalid. Please upload another document to retry real name verification</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_BUSINESS_ID</strong> - Business id does not match the business id in the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_BUSINESS_NAME</strong> - Business name does not match the business name in the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_ID</strong> - Document id does not match the id in the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_NAME</strong> - Document name does not match the name in the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_DOCUMENT_TYPE</strong> - Document type does not match the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_INFO</strong> - The information provided for the registrant does not match the document</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_LOCALITY</strong> - Registrant region is overseas, but a local identity document was provided</li> <li><strong style='margin-left: 12px;'>REJECTED_MISMATCH_REGISTRANT_NAME</strong> - Registrant name has been changed, so the request must be resubmitted</li> <li><strong style='margin-left: 12px;'>REJECTED_UNABLE_TO_OPEN</strong> - Registrant identification could not be opened. Please upload the document again to retry real name verification</li> <li><strong style='margin-left: 12px;'>REJECTED_UNABLE_TO_VERIFY</strong> - Unable to initiate verification. Please upload the document again to retry real name verification</li> <li><strong style='margin-left: 12px;'>REJECTED_UNKNOWN_ERROR</strong> - Document was rejected due to an unknown error. For more information, contact customer support</li> <li><strong style='margin-left: 12px;'>UNABLE_TO_RETRIEVE_STATUS</strong> - Unable to retrieve status for the real name verification process. Retry, if this status persists, contact customer support</li> </ul>
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	City string `json:"city"`
	Country string `json:"country"` // Two-letter ISO country code to be used as a hint for target region<br/><br/> NOTE: These are sample values, there are many <a href='http://www.iso.org/iso/country_codes.htm'>more</a>
	Postalcode string `json:"postalCode"` // Postal or zip code
	State string `json:"state"` // State or province or territory
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"`
}

// DomainRenew represents the DomainRenew schema from the OpenAPI specification
type DomainRenew struct {
	Period int `json:"period,omitempty"` // Number of years to extend the Domain. Must not exceed maximum for TLD. When omitted, defaults to `period` specified during original purchase
}

// DomainForwarding represents the DomainForwarding schema from the OpenAPI specification
type DomainForwarding struct {
	Fqdn string `json:"fqdn"` // The fqdn (domain or sub domain) to forward (ex somedomain.com or sub.somedomain.com)
	Mask DomainForwardingMask `json:"mask,omitempty"`
	TypeField string `json:"type"` // The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>
	Url string `json:"url"` // Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)
}

// DNSRecordCreateTypeName represents the DNSRecordCreateTypeName schema from the OpenAPI specification
type DNSRecordCreateTypeName struct {
	Port int `json:"port,omitempty"` // Service port (SRV only)
	Priority int `json:"priority,omitempty"` // Record priority (MX and SRV only)
	Protocol string `json:"protocol,omitempty"` // Service protocol (SRV only)
	Service string `json:"service,omitempty"` // Service type (SRV only)
	Ttl int `json:"ttl,omitempty"`
	Weight int `json:"weight,omitempty"` // Record weight (SRV only)
	Data string `json:"data"`
}
