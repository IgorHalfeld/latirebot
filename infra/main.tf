provider "azurerm" {
  version = "~>2.0"
  features {}
}

resource "azurerm_resource_group" "latire_resource_group" {
  name     = "latire-resource-group"
  location = "westus2"
}

resource "azurerm_virtual_network" "latire_virtual_network" {
  name                = "latire-network"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.latire_resource_group.location
  resource_group_name = azurerm_resource_group.latire_resource_group.name

  tags = {
    environment = "terraform"
  }
}

resource "azurerm_subnet" "latire_subnet" {
  name                 = "latire-subnet"
  resource_group_name  = azurerm_resource_group.latire_resource_group.name
  virtual_network_name = azurerm_virtual_network.latire_virtual_network.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_public_ip" "latire_publicip" {
    name                         = "latire-public-ip"
    location                     = azurerm_resource_group.latire_resource_group.location
    resource_group_name          = azurerm_resource_group.latire_resource_group.name
    allocation_method            = "Dynamic"

    tags = {
        environment = "terraform"
    }
}

# Create Network Security Group and rule
resource "azurerm_network_security_group" "latire_security_group" {
    name                = "latire-secgroup"
    location            = azurerm_resource_group.latire_resource_group.location
    resource_group_name = azurerm_resource_group.latire_resource_group.name
    
    security_rule {
        name                       = "SSH"
        priority                   = 1001
        direction                  = "Inbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "22"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
    }

    tags = {
        environment = "terraform"
    }
}

resource "azurerm_network_interface" "latire_network_interface" {
  name                = "latire-network-interface"
  location            = azurerm_resource_group.latire_resource_group.location
  resource_group_name = azurerm_resource_group.latire_resource_group.name

  ip_configuration {
    name                          = "ip_configs"
    subnet_id                     = azurerm_subnet.latire_subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.latire_publicip.id
  }
}

resource "azurerm_network_interface_security_group_association" "latire_nic_sec_association" {
    network_interface_id      = azurerm_network_interface.latire_network_interface.id
    network_security_group_id = azurerm_network_security_group.latire_security_group.id
}

resource "tls_private_key" "latire_ssh" {
  algorithm = "RSA"
  rsa_bits = 4096
}
output "tls_private_key" { value = "${tls_private_key.latire_ssh.private_key_pem}" }

resource "azurerm_linux_virtual_machine" "latirevm" {
  name                  = "latire-latirevm"
  location              = azurerm_resource_group.latire_resource_group.location
  resource_group_name   = azurerm_resource_group.latire_resource_group.name
  network_interface_ids = [azurerm_network_interface.latire_network_interface.id]
  size               = "Standard_DS1_v2"

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  
  os_disk {
    name                 = "myosdisk1"
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  admin_username = "latire"
  admin_ssh_key {
      username       = "latire"
      public_key     = tls_private_key.latire_ssh.public_key_openssh
  }

  tags = {
    environment = "terraform"
  }
}

