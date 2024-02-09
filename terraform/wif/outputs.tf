output "provider_name" {
  value = module.gh_oidc.provider_name
}

output "gh_service_account_email" {
  value = module.gh_service_account.email
}

output "cr_service_account_email" {
  value = module.cr_service_account.email
}
