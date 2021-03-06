# Dumbphone

![Screenshot](docs/screenshot.png)

This is a command-line tool for making my Alcatel A392CC flip phone just a little bit smarter. It builds on a few other projects in various states of completion:

 - https://github.com/andyinabox/dumbp-a392cc-cli
 - https://github.com/andyinabox/dumbp-gmap-directions
 - https://github.com/andyinabox/dumbp-read

To setup the project:

```bash
make setup
```

This will prompt you for your Google Maps API key, and install the `go-bindata` dependency

To build:

```bash
make
```

To install into `/usr/local/bin` (requires `sudo` access):

```bash
make install
```

**Note:** This is a Go learner project so use at your own risk.

# Goals

 - [x] Basic cli framework (`dumbp` command)
 - [x] `directions` command that generates HTML directions to be transfered to a phone (already implemented in python in the [dumbp-gmap-directions](https://github.com/andyinabox/dumbp-gmap-directions) repo
 - [x] bluetooth file transfer (already in [dumbp-gmap-directions](https://github.com/andyinabox/dumbp-gmap-directions))
 - [ ] filesystem file transfer (when using phone as external drive)
 - [ ] `podcast` command that will download podcasts to be transfered to phone
 - [ ] `calendar` command that will sync with Google Calendar using `.vcf` files
 - [ ] `music` command that makes managing music on phone easier
 - [ ] `sync` command that will automatically sync multiple functions
 - [ ] `config` command for configuring different functions
 - [x] `reader` command for syncing articles to phone (there's already a python implementation in [dumbp-read](https://github.com/andyinabox/dumbp-read)

