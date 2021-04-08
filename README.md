# Prompt Credentials Provider for N26 API Client

[![Build Status](https://github.com/nhatthm/n26prompt/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/n26prompt/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/n26prompt/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/n26prompt)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/httpmock)](https://goreportcard.com/report/github.com/nhatthm/httpmock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/n26prompt)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=44BH3UA6UUAMJ)

`n26prompt` uses [AlecAivazis/survey](https://github.com/AlecAivazis/survey) to get credentials from the prompt.

## Prerequisites

- `Go >= 1.14`

## Install

```bash
go get github.com/nhatthm/n26prompt
```

## Usage

**Examples**

```go
package mypackage

import (
	"github.com/nhatthm/n26api"
	"github.com/nhatthm/n26prompt/credentials"
)

func buildClient() (*n26api.Client, error) {
	c := n26api.NewClient(
		credentials.WithCredentialsAtLast(),
	)

	return c, nil
}
```

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=44BH3UA6UUAMJ)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
