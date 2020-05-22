# website-builder

Automated builder for static websites.


## Description

This is a simple program that can build static websites automatically out of a source tree. It tries to detect the static site generator and runs it as required. It was designed for automated environments, like GitHub Actions/Travis CI and as a builder for software listening to GitHub webhooks.


### Supported static site generators

- `GNU make`: Any source tree containing a `Makefile` file that implements the `website-builder` rule will have this target ran. The rule should write the output files to `$OUTPUT_DIR`.
- `blogc-make`: Any source tree containing a `blogcfile` file will be built as a `blogc-make` website.

Please note that the required software must be installed.


## Usage

The basic usage is:

    $ website-builder SOURCE_DIR OUTPUT_DIR [SYMLINK]

Where `SOURCE_DIR` is the source directory, `OUTPUT_DIR` is the output directory (will be created if needed), and `SYMLINK` is an optional symbolic link to update after successful build. **Please note that the old symlink is removed and the old symlink target directory is also removed.**
