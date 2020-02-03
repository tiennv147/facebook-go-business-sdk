package adaccount

const (
	prefix = "act_%s"
)

var (
	fields = []string{
		"account_id",
		"account_status",
		//"ad_account_creation_request",
		//"ad_account_promotable_objects",
		"age",
		"agency_client_declaration",
		"amount_spent",
		"attribution_spec",
		"balance",
		"business",
		"business_city",
		"business_country_code",
		"business_name",
		"business_state",
		"business_street",
		"business_street2",
		"business_zip",
		"capabilities",
		"created_time",
		"currency",
		"disable_reason",
		"end_advertiser",
		"end_advertiser_name",
		"extended_credit_invoice_group",
		"failed_delivery_checks",
		"fb_entity",
		"funding_source",
		"funding_source_details",
		"has_migrated_permissions",
		//"has_page_authorized_adaccount",
		"id",
		"io_number",
		"is_attribution_spec_system_default",
		"is_direct_deals_enabled",
		"is_in_3ds_authorization_enabled_market",
		"is_in_middle_of_local_entity_migration",
		"is_notifications_enabled",
		"is_personal",
		"is_prepay_account",
		"is_tax_id_required",
		"line_numbers",
		"media_agency",
		"min_campaign_group_spend_cap",
		"min_daily_budget",
		"name",
		"offsite_pixels_tos_accepted",
		"owner",
		"partner",
		"rf_spec",
		//"show_checkout_experience",
		"spend_cap",
		"tax_id",
		"tax_id_status",
		"tax_id_type",
		"timezone_id",
		"timezone_name",
		"timezone_offset_hours_utc",
		"tos_accepted",
		"user_tasks",
		"user_tos_accepted",
	}
)
