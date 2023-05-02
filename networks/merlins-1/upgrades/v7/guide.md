# v6 to v7 Upgrade Guide

## Memory Requirements

For this upgrade, nodes will need a total of 64GB of memory. This must
consist of a **minimum** of 32GB of RAM, while the remaining 32GB can be
swap. For best results, use 64GB of physical memory.

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

## Install and setup Cfuryvisor

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

Cfuryvisor requires some ENVIRONMENT VARIABLES be set in order to
function properly. We recommend setting these in your `.profile` so it
is automatically set in every session.

For validators we recommmend setting

- `DAEMON_ALLOW_DOWNLOAD_BINARIES=false` for security reasons
- `DAEMON_LOG_BUFFER_SIZE=512` to avoid a bug with extra long log
    lines crashing the server.
- `DAEMON_RESTART_AFTER_UPGRADE=true` for unattended upgrades

```{=html}
<!-- -->
```

    echo "# Setup Cfuryvisor" >> ~/.profile
    echo "export DAEMON_NAME=merlinsd" >> ~/.profile
    echo "export DAEMON_HOME=$HOME/.merlinsd" >> ~/.profile
    echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
    echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
    echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
    echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
    source ~/.profile

You may leave out `UNSAFE_SKIP_BACKUP=true`, however the backup takes a
decent amount of time and public snapshots of old states are available.

Finally, you should copy the current merlinsd binary into the
cfuryvisor/genesis folder.

    cp $GOPATH/bin/merlinsd ~/.merlinsd/cfuryvisor/genesis/bin

## Prepare for upgrade (v7)

To prepare for the upgrade, you need to create some folders, and build
and install the new binary.

    mkdir -p ~/.merlinsd/cfuryvisor/upgrades/v7/bin
    cd $HOME/merlins
    git pull
    git checkout v7.0.2
    make build
    cp build/merlinsd ~/.merlinsd/cfuryvisor/upgrades/v7/bin

Now cfuryvisor will run with the current binary, and will automatically
upgrade to this new binary at the appropriate height if run with:

    cfuryvisor start

Please note, this does not automatically update your
`$GOPATH/bin/merlinsd` binary, to do that after the upgrade, please run
`make install` in the merlins source folder.

## Further Help

If you need more help, please go to <https://docs.merlins.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.
