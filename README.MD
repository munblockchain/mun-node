# MUN Blockchain Mainnet Node

Welcome to the MUN Blockchain Mainnet Node repository. MUN Blockchain is a state-of-the-art decentralized network designed to support fast, secure, and reliable transactions. By running your own mainnet node, you are participating in and contributing to the robustness, security, and overall health of the network.

## Prerequisites

For optimal operation of your MUN Blockchain Node, we recommend the following system configuration:

- CPU: 6 cores or more
- Memory: 16GB or more
- Operating System: Ubuntu 22.04 LTS
- Network: Allow all incoming connections on TCP ports 26656 and 26657
- IP: Static IP address
- Storage: At least 1TB of SSD storage

## Installation

### System Update and Installing Essential Tools

Make sure your system packages are up to date and install the necessary tools:

```bash
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```

## Golang Installation:

MUN Blockchain Node runs on Go. Install Go 1.20 as follows:
```
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.20
source ~/.profile
```

To verify Go installation, run:
```
go version
```
// This command should return go version go1.20 linux/amd64.

## MUN Node Setup
Clone the MUN Node repository:

```
git clone https://github.com/munblockchain/mun-node
cd mun-node
```

Next, install the executables:

```
sudo rm -rf ~/.mun
go mod tidy
make install
```
Upon successful installation, proceed with the following directory setup:

```
clear
mkdir -p ~/.mun/upgrade_manager/upgrades
mkdir -p ~/.mun/upgrade_manager/genesis/bin
```

## Genesis Binary and Upgrade Manager Setup
Make a symbolic link of the genesis binary to the upgrade manager:

```
cp $(which mund) ~/.mun/upgrade_manager/genesis/bin
sudo cp $(which mund-manager) /usr/bin
```

## Node Initialization
Initialize the validator with a moniker of your choice (example: solid-moon-rock):

```
mund init [moniker_name] --chain-id mun-1
```

Create a new wallet address, store the generated seeds, and add MUN to your wallet:

```
mund keys add [wallet_name] --keyring-backend test
```

## Genesis and P2P Configuration
Fetch the genesis.json file from the genesis node:

curl https://mainnet1rpc.mun.money/genesis | jq ".result.genesis" > ~/.mun/config/genesis.json

Next, update the seed in config.toml for P2P connection:
```
nano ~/.mun/config/config.toml
```
Add the following line:
```
seeds = "SEEDS"
```
The SEEDS information can be requested from our Discord community.

##  MUN Node Service Setup
Ensure to replace [user_name], [group_name], and [DAEMON_HOME_PATH] with your specific system information.



Create the service file /etc/systemd/system/mund.service with the following content:


```
sudo nano /etc/systemd/system/mund.service
```


```
[Unit]
Description=mund
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=root
Group=root
Environment=DAEMON_NAME=mund
Environment=DAEMON_HOME=/root/.mun
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/mund-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"
StandardOutput=file:/var/log/mund/mund.log
StandardError=file:/var/log/mund/mund_error.log
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
User, group, and DAEMON_HOME path can be obtained as follows:


```
whoami  # Outputs user_name
cd ~/.mun
pwd  # Outputs DAEMON_HOME path
```

Create log files and start running the node:

```
make log-files
sudo systemctl enable mund
sudo systemctl start mund
```

To verify that your node is running properly:
```
mund status
```

## Becoming a Validator
Once your node is fully synchronized with the network and you've secured some MUN, you can become a validator by staking your MUN. To check synchronization status, visit https://mainnet1rpc.mun.money/status and compare the latest block height with your node's block height, obtainable via http://[Your_Node_IP]:26657/status.

When ready, stake your MUN:
```
mund tx staking create-validator --from [wallet_name] --moniker [moniker_name] --pubkey $(mund tendermint show-validator) --chain-id mun-1 --keyring-backend test --amount 2000000000000000umun --commission-max-change-rate 0.01 --commission-max-rate 0.2 --commission-rate 0.1 --min-self-delegation 1 --fees 20000umun -y
```
Congratulations, you're now an active participant in the MUN Blockchain network! Your contributions make the network more secure, robust, and reliable.

## Community Support

If you encounter any issues while setting up your MUN Blockchain Node, don't hesitate to reach out to our friendly and supportive community on Discord. Our community members and developers are always ready to assist you!
