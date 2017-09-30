remote_file "#{node["golang"]["binpath"]}/#{node["golang"]["dep"]["binname"]}" do
    source "#{node["golang"]["dep"]["download_url_prefix"]}/v#{node["golang"]["dep"]["version"]}/#{node["golang"]["dep"]["download_filename"]}"
    mode '0755'
    action :create
    not_if { ::File.exist?("#{node["golang"]["binpath"]}/#{node["golang"]["dep"]["binname"]}") }
end
