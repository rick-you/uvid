# Getting Started

## Prerequisites

Makes sure you have golang installed

<https://go.dev>

## Setup Dashboard

1. Install server

   ```sh
   go install github.com/rick-you/uvid@latest
   ```

2. Run server

   ```sh
   go/bin/uvid
   ```

3. The dashboard will be served from `8080` port
4. `./uvid.db` is sqlite database stores all your data
5. That's all.

## Install SDK

To start track your website, you must install SDK in your website

1. Install SDK

```sh
npm install uvid-js
```

2. Initialize SDK, By default, SDK will sends page view and performance data on page load

```js
import { init } from 'uvid-js';

const sdk = init({
  host: 'YOUR_DASHBOARD_URI/:8080',
  sessionMeta: {
    // optionally, track additional meta data
    userId: '123',
  },
  // optionally, provide your website's build version
  appVersion: '1.0.0',
});

// Track an js error
sdk.error(new Error('This is an js error!'));
// Track an custom event action and value
sdk.event('register', 'some-user@email.com');
// Track an custom event by HTML attributes
`<button data-uvid-action="register" data-uvid-value="some-user@email.com">Register</button>`;
//  When user click the button, uvid-js will track it and call uvid.event('register', 'some-user@email.com')

// Track an http request
sdk.http({
  resource: 'http://some-api.com',
  status: 500,
  method: 'GET',
  headers: '',
});
```

3. Publish above change, when user visit your website, you will see data from dashboard
