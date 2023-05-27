// fetch-polyfill.js
import fetch, { Headers } from 'node-fetch'

if (!globalThis.fetch) {
  globalThis.fetch = fetch
  globalThis.Headers = Headers
}
