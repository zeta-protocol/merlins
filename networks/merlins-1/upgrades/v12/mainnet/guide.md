# v11 to v12 Mainnet Upgrade Guide

Merlins v12 Gov Prop: <https://www.mintscan.io/merlins/proposals/335>

Countdown: <https://www.mintscan.io/merlins/blocks/6246000>

Height: 6246000

## Memory Requirements

This upgrade will **not** be resource intensive. With that being said, we still recommend having 64GB of memory. If having 64GB of physical memory is not possible, the next best thing is to set up swap.

Short version swap setup instructions:

``` {.sh}
sudo swapoff -a
sudo fallocate -l 32G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

To persist swap after restart:

``` {.sh}
sudo cp /etc/fstab /etc/fstab.bak
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

In depth swap setup instructions:
<https://www.digitalocean.com/community/tutorials/how-to-add-swap-space-on-ubuntu-20-04>

## First Time Cfuryvisor Setup

If you have never setup Cfuryvisor before, follow the following instructions.

If you have already setup Cfuryvisor, skip to the next section.

We highly recommend validators use cfuryvisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can
pre-install new binaries, and cfuryvisor will automatically update them
based on on-chain SoftwareUpgrade proposals.

You should review the docs for cfuryvisor located here:
<https://docs.cosmos.network/master/run-node/cfuryvisor.html>

If you choose to use cfuryvisor, please continue with these
instructions:

To install Cfuryvisor:

``` {.sh}
go install github.com/cosmos/cosmos-sdk/cfuryvisor/cmd/cfuryvisor@v1.0.0
```

After this, you must make the necessary folders for cosmosvisor in your
daemon home directory (\~/.merlinsd).

``` {.sh}
mkdir -p ~/.merlinsd
mkdir -p ~/.merlinsd/cfuryvisor
mkdir -p ~/.merlinsd/cfuryvisor/genesis
mkdir -p ~/.merlinsd/cfuryvisor/genesis/bin
mkdir -p ~/.merlinsd/cfuryvisor/upgrades
```

Copy the current v11 merlinsd binary into the
cfuryvisor/genesis folder and v11 folder.

```{.sh}
cp $GOPATH/bin/merlinsd ~/.merlinsd/cfuryvisor/genesis/bin
mkdir -p ~/.merlinsd/cfuryvisor/upgrades/v11/bin
cp $GOPATH/bin/merlinsd ~/.merlinsd/cfuryvisor/upgrades/v11/bin
```

Cfuryvisor is now ready to be set up for v12.

Set these environment variables:

```{.sh}
echo "# Setup Cfuryvisor" >> ~/.profile
echo "export DAEMON_NAME=merlinsd" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.merlinsd" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

## Cfuryvisor Upgrade

Create the v12 folder, make the build, and copy the daemon over to that folder

```{.sh}
mkdir -p ~/.merlinsd/cfuryvisor/upgrades/v12/bin
cd $HOME/merlins
git pull
git checkout v12.0.0
make build
cp build/merlinsd ~/.merlinsd/cfuryvisor/upgrades/v12/bin
```

Now, at the upgrade height, Cfuryvisor will upgrade to the v12 binary

## Manual Option

1. Wait for Merlins to reach the upgrade height (FILL ME IN WITH UPGRADE HEIGHT)

2. Look for a panic message, followed by endless peer logs. Stop the daemon

3. Run the following commands:

```{.sh}
cd $HOME/merlins
git pull
git checkout v12.0.0
make install
```

4. Start the merlins daemon again, watch the upgrade happen, and then continue to hit blocks

## Further Help

If you need more help, please go to <https://docs.merlins.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.
