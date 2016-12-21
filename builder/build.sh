#!/usr/bin/env bash

# INSTALL GIT
sudo apt-get update && sudo apt-get -y install git 

# INSTALL DEPOT TOOLS
git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
export PATH=`pwd`/depot_tools:"$PATH"

# ADD DEPOT TOOLS PATH PERMANENTLY
# echo 'export PATH='`pwd`'/depot_tools:"$PATH"' >> ~/.profile
export PATH='`pwd`'/depot_tools:"$PATH"

# DOWNLOAD AND CHECKOUT SOURCE CODE
mkdir webrtc-checkout
cd webrtc-checkout
fetch --nohooks webrtc
gclient sync
cd src
git checkout master

# BUILDING
# NOTICE: Debug builds are component builds (shared libraries) 
# by default unless is_component_build=false is passed to gn gen --args. 
# Release builds are static by default.
# To generate ninja project files for a Release build instead:

# DEFAULT DEBUG BUILD
gn gen out/Debug

# OPTIONAL RELEASE BUILD
# gn gen out/Release --args='is_debug=false'

# THIS IS THE COMMAND TO CLEAN THE BUILD
# To clean all build artifacts in a directory but leave the current GN configuration untouched (stored in the args.gn file), do:
# gn clean out/Debug
# gn clean out/Release

# COMPILE
ninja -C out/Debug
