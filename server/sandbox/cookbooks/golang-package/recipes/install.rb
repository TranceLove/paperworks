dest_file = "go#{node["golang"]["version"]}.#{node["golang"]["arch"]}.tar.gz"

tar_extract "#{node["golang"]["download_url_prefix"]}/#{dest_file}" do
    target_dir node["golang"]["extract_to"]
    creates node["golang"]["dist_path"]
    not_if { ::File.exist?(node["golang"]["dist_path"]) }
end

bash "Update system-wide $PATH environment variable" do
    code <<-EOH
if [ -z `grep '#{node["golang"]["binpath"]}' /etc/profile.d/apps-bin-path.sh` ]; then
echo "PATH=$PATH:#{node["golang"]["binpath"]}" >> /etc/profile.d/apps-bin-path.sh
fi
    EOH
end

file node["golang"]["gopath_profile"] do
  content <<-EOF
export GOPATH=#{node['golang']['default_gopath']}
PATH=$GOPATH/bin:$PATH
EOF
  mode '0644'
  owner 'root'
  group 'root'
  action :create
  not_if { ::File.exist?(node["golang"]["gopath_profile"]) }
end
