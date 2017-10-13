bash "Create paperworks database for tests" do
    user "postgres"
    code "psql -c \"CREATE DATABASE #{node['paperworks']['database_name']}_test\""
    not_if "psql -tc \"SELECT 1 FROM pg_database WHERE datname = '#{node["paperworks"]["database_name"]}_test'\" | grep -q 1", :user => "postgres"
end
