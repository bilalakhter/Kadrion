Kadrion is CLI tool to test multiple api endpoints for performance metrics, It reads configuration from a yaml file and perform the test according to it.

# Install

To install this tool clone this directory

    git clone https://github.com/bilalakhter/kadrion

    cd kadrion

    make install

Or you can downloed the binary from github Release

# Remove

To uninstall the tool again move into kadrion git directory

    make clean

Or you can remove the exported path -> export PATH="$PATH:/opt/kadrion" from ~/.bashrc and delete /opt/kadrion

# Usage

Kadrion reads parameters from yaml file tconfig.yaml to apply changes and perform test against api endpoints use:

    kadrion apply tconfig.yaml

```

```
