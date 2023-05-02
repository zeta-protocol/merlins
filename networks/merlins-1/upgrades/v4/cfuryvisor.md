# Install and setup Cfuryvisor

We highly recommend validators use cfuryvisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can preinstall
new binaries, and cfuryvisor will automatically update them based on
on-chain SoftwareUpgrade proposals.

You should review the docs for cfuryvisor located here:
<https://docs.cosmos.network/master/run-node/cfuryvisor.html>

If you choose to use cfuryvisor, please continue with these
instructions:

To install Cfuryvisor:

    git clone https://github.com/cosmos/cosmos-sdk
    cd cosmos-sdk
    git checkout v0.42.9
    make cfuryvisor
    cp cfuryvisor/cfuryvisor $GOPATH/bin/cfuryvisor
    cd $HOME

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
    source ~/.profile

Finally, you should copy the current merlinsd binary into the
cfuryvisor/genesis folder.

    cp $GOPATH/bin/merlinsd ~/.merlinsd/cfuryvisor/genesis/bin

Prepare for upgrade (v4)
------------------------

To prepare for the upgrade, you need to create some folders, and build
and install the new binary.

    mkdir -p ~/.merlinsd/cfuryvisor/upgrades/v4/bin
    git clone https://github.com/merlins-labs/merlins
    cd merlins
    git checkout v4.0.0
    make build
    cp build/merlinsd ~/.merlinsd/cfuryvisor/upgrades/v4/bin

Now cfuryvisor will run with the current binary, and will automatically
upgrade to this new binary at the appropriate height if run with:

    cfuryvisor start

Please note, this does not automatically update your
`$GOPATH/bin/merlinsd` binary, to do that after the upgrade, please run
`make install` in the merlins source folder.
