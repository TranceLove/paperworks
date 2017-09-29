bash "Create paperworks database" do
    user "postgres"
    code "psql -c \"CREATE DATABASE #{node['paperworks']['database_name']}\""
    not_if "psql -tc \"SELECT 1 FROM pg_database WHERE datname = '#{node["paperworks"]["database_name"]}'\" | grep -q 1", :user => "postgres"
end
