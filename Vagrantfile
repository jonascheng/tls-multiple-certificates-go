$go = <<-SCRIPT
/vagrant/provision/go.sh
/vagrant/provision/protoc.sh
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.box = "boxomatic/ubuntu-18.04"
  config.vm.box_version = "20210723.0.1"

  config.vm.define "server" do |node|
    node.vm.hostname = "server"
    node.vm.provision "shell", inline: $go, privileged: false
    node.vm.network "private_network", ip: "172.31.1.10", hostname: true, netmask: '255.255.255.0'
  end

  config.vm.define "client1" do |node|
    node.vm.hostname = "client1"
    node.vm.provision "shell", inline: $go, privileged: false
    node.vm.network "private_network", ip: "172.31.1.20", hostname: true, netmask: '255.255.255.0'
  end

  config.vm.define "client2" do |node|
    node.vm.hostname = "client2"
    node.vm.provision "shell", inline: $go, privileged: false
    node.vm.network "private_network", ip: "172.31.1.30", hostname: true, netmask: '255.255.255.0'
  end

end
