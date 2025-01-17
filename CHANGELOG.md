# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].

## 1.0.0-rc

### Fixed

* Remove duplicated line on light theme
* Lights background on light theme
* Form line color on light mac
* Remove shadow from select pick

### Added

* Linux AppImage build (Ubuntu, Fedora, etc GTK-based)
* Signing and notarization application for macOS
* The OS value selection is now stored persistently and restored after loading

## [1.0.0-beta7]

### Changed

* Sidebar design

### Add

* Mechanism for searching for application updates
* `About` menu
* OS handling
* Light theme support

### Fixed

* Improved metadata in Info.plist file (The permission would need to be reissued).

## [1.0.0-beta6]

### Changed

* "Fixed On" mode renamed to "Static"

### Added

* Displaying the firmware version on the Device view
* Dynamic keyboard preview. Halo75, Halo65 and Halo96 support with correct layout.
* Help message for Sidelight. It may seem like the settings don't apply, but it's more likely that your keyboard isn't charged to the max

### Fixed

* Disabled custom mode selection. I'll bring it back when the support is complete

## [1.0.0-beta5]

### Added

* Real Halo65 support (the preview still shows the keyboard at 75%)

### Fixed

* Application signature
* Version on UI

## [1.0.0-beta4]

### Added

* Initial dirty support for Halo65

### [1.0.0-beta3]

### Added

* Color changing modal transition

### Changed

* Now the simulation runs on <kbd>Option</kbd> + <kbd>Shift</kbd> + <kbd>S</kbd>
* Fine-tuning fonts, borders and margins
* Simplified ranges

### Fixed

* Minimal window height (set to 410 pixels)
* Panic on simulation rejection
* Cyclic pop-up window on simulation rejection
* Application title
* Improve memory usage

## [1.0.0-beta2]

### Added

* Color edit button. Right click on a color still works.

### Fixed

* Window drag on loading state
* Add missing window border
* Unfocused window state

## [1.0.0-beta1]

Initial beta release

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
[1.0.0-beta1]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta1
[1.0.0-beta2]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta2
[1.0.0-beta3]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta3
[1.0.0-beta4]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta4
[1.0.0-beta5]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta5
[1.0.0-beta6]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta6
[1.0.0-beta7]: https://github.com/mishamyrt/Nuga/releases/tag/v1.0.0-beta7
