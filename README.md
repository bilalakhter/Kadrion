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

tconfig.yaml needs to be places in current directory where the command is executed and needs to follow this structure to work correctly

![](https://hslytkjiifoqndixiofd.supabase.co/storage/v1/object/public/files/Untitled%20design.png)

# Feature

Perform load testing targeting different endpoints and check there response time at different attempts under different load

![](https://hslytkjiifoqndixiofd.supabase.co/storage/v1/object/public/files/Screenshot_20231012_014947.png)

# Continuous testing

Easily add a stage in your CI/CD pipeline by placing the tconfig.yaml file in your repository
