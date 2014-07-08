VAGRANTFILE_API_VERSION = "2"

ENV['VAGRANT_DEFAULT_PROVIDER'] ||= 'docker'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.define "go" do |web|
    web.vm.provider 'docker' do |d|
      d.image   = "busybox"
      d.has_ssh = true
      d.cmd     = ["/srv/main", "http://cloudspace.com"]
      d.args    "-v '/srv:/srv/go-url-lengthener'"
    end
    
    # web.vm.synced_folder ".", "/srv", owner: 'root', group: 'root'

    web.ssh.username = "root"
    # config.ssh.private_key_path = [File.join(ENV['HOME'], '.ssh', 'id_rsa')]

  end
end