## WebRTC sub repos

* All sub repo version metadata is stored in `.gclient_entries`
* WebRTC core: `https://chromium.googlesource.com/external/webrtc.git`
* chromium base: `https://chromium.googlesource.com/chromium/src/base`
* chromium build: `https://chromium.googlesource.com/chromium/src/build`
* chromium buildtools: `https://chromium.googlesource.com/chromium/buildtools.git`
* chromium testing: `https://chromium.googlesource.com/chromium/src/testing`
* chromium third_party: `https://chromium.googlesource.com/chromium/src/third_part`
* chromium tools: `https://chromium.googlesource.com/chromium/src/tools`

## source tree

```
❯ tree -L 1
.
├── AUTHORS
├── BUILD.gn
├── DEPS // dep sub repos address and version
├── LICENSE
├── LICENSE_THIRD_PARTY
├── OWNERS
├── PATENTS
├── PRESUBMIT.py
├── README.md
├── WATCHLISTS
├── base // from chromium, only to be used by Android unit tests (and the libevent dependency)
├── build // from chromium, contains most of the build toolchain.
├── build_overrides
├── buildtools
├── check_root_dir.py
├── cleanup_links.py
├── codereview.settings
├── data
├── infra
├── ios
├── license_template.txt
├── out // build target files
├── pylintrc
├── resources
├── testing // test-related scripts
├── third_party // lots of third_party code and foremost the BUILD.gn files for all.
├── tools // tools required for build and development.
├── tools-webrtc // webrtc's own tools
└── webrtc

14 directories, 15 files
```