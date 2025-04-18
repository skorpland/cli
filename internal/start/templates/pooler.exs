{:ok, _} = Application.ensure_all_started(:powerpooler)

{:ok, version} =
  case Powerpooler.Repo.query!("select version()") do
    %{rows: [[ver]]} -> Powerpooler.Helpers.parse_pg_version(ver)
    _ -> nil
  end

params = %{
  "external_id" => "{{ .ExternalId }}",
  "db_host" => "{{ .DbHost }}",
  "db_port" => {{ .DbPort }},
  "db_database" => "{{ .DbDatabase }}",
  "require_user" => false,
  "auth_query" => "SELECT * FROM pgbouncer.get_auth($1)",
  "default_max_clients" => {{ .DefaultMaxClients }},
  "default_pool_size" => {{ .DefaultPoolSize }},
  "default_parameter_status" => %{"server_version" => version},
  "users" => [%{
    "db_user" => "pgbouncer",
    "db_password" => "{{ .DbPassword }}",
    "mode_type" => "{{ .ModeType }}",
    "pool_size" => {{ .DefaultPoolSize }},
    "is_manager" => true
  }]
}

if !Powerpooler.Tenants.get_tenant_by_external_id(params["external_id"]) do
  {:ok, _} = Powerpooler.Tenants.create_tenant(params)
end
