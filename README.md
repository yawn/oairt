# OpenAI Realtime API

OpenAI [Realtime API](https://openai.com/index/introducing-the-realtime-api/) type definitions for Go. WIP because I just got access and not tested anything yet.

Wherever noted in the [documentation](https://platform.openai.com/docs/api-reference/realtime-client-events) (or somewhat obvious), optional values are references.

Updates (2024-10-04):

- Currently (for me) the API seems partially broken, at least for the text example. I'll add voice (aka the intended use case) later today.
- Hit the 100 RPD limit - voice client works already but the limit is too slow for client development, will continue tomorrow