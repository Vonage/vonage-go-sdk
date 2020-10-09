# Vonage Go SDK

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/vonage/vonage-go-sdk)](https://pkg.go.dev/mod/github.com/vonage/vonage-go-sdk)
![Actions](https://github.com/vonage/vonage-go-sdk/workflows/Vonage%20Go%20SDK/badge.svg)

<img src="https://developer.nexmo.com/assets/images/Vonage_Nexmo.svg" height="48px" alt="Nexmo is now known as Vonage" />

This is the community-supported Golang library for [Vonage](https://vonage.com). It has support for most of our APIs, but is still under active development. Issues, pull requests and other input is very welcome. The [package documentation is available on pkg.go](https://pkg.go.dev/mod/github.com/vonage/vonage-go-sdk).

If you don't already know Vonage: We make telephony APIs. If you need to make a call, check a phone number, or send an SMS then you are in the right place! If you don't have one yet, you can [sign up for a Vonage account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&amp;utm_medium=github&amp;utm_campaign=vonage-go) and get some free credit to get you started.

  * [Installation](#installation)
  * [Usage](#usage)
  * [API Support](#api-support)
  * [Contributions](#contributions)
    * [Using a Local Branch](#using-a-local-branch)
  * [Getting Help](#getting-help)
  * [Further Reading](#further-reading)

## Installation

Find current and past releases on the [releases page](https://github.com/vonage/vonage-go-sdk/releases).

Import the package and use it in your own project

```
import ("github.com/vonage/vonage-go-sdk")
```

## Usage

Usage examples are in the `docs/` folder - also rendered via GitHub pages: <https://vonage.github.io/vonage-go-sdk/>

## API Support

Current state of API support in this library:

| API   | API Release Status |  Supported?
|----------|:---------:|:-------------:|
| Account API | General Availability |❌|
| Alerts API | General Availability |❌|
| Application API | General Availability |✅|
| Audit API | Beta |❌|
| Conversation API | Beta |❌|
| Dispatch API | Beta |❌|
| External Accounts API | Beta |❌|
| Media API | Beta | ❌|
| Messages API | Beta |❌|
| Number Insight API | General Availability |✅|
| Number Management API | General Availability |✅|
| Pricing API | General Availability |❌|
| Redact API | Developer Preview |❌|
| Reports API | Beta |❌|
| SMS API | General Availability |✅|
| Verify API | General Availability |✅|
| Voice API | General Availability |✅|

## Contributions

Yes please! This library is open source, community-driven, and benefits greatly from the input of its users.

Please make all your changes on a branch, and open a pull request, these are welcome and will be reviewed with delight! If it's a big change, it is recommended to open an issue for discussion before you start.

All changes require tests to go with them.

### Using a Local Branch

Refer to [this excellent blog post](https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/) for instructions on how to use a local clone of this repository as the import in your own project. This is really useful when you are using a version of the library other than the latest stable release - for example if you are working on a change, or testing an open pull request.

## Getting Help
 
We love to hear from you so if you have questions, comments or find a bug in the project, let us know! You can either:
 
* Open an [issue on this repository](https://github.com/Vonage/vonage-go-sdk/issues)
* Tweet at us! We're [@VonageDev on Twitter](https://twitter.com/VonageDev)
* Or [join the Vonage Community Slack](https://developer.nexmo.com/community/slack)
 
## Further Reading
 
* Check out the Developer Documentation at <https://developer.nexmo.com> - you'll find the API references for all the APIs there as well
* The documentation for the library: <https://godoc.org/github.com/vonage/vonage-go-sdk>


## License

This library is released under the [Apache 2.0 License][license]

[license]: LICENSE.txt
