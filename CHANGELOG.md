# Changelog

## v0.3.1 (12-09-2025)

- Refactored `InlineQueryResult` into an interface.
- Each `InlineQueryResult` type (`Article`, `Photo`, `Video`, etc.) now implements:
  - `isInlineQueryResult()` marker
  - `MarshalCustom()` for JSON serialization with the correct `"type"` field.
- This improves extensibility and clarity over the previous single-struct + `MarshalJSON` override approach.

## v0.3.0 (12-09-2025)

- Added support for `InputMedia` (photo, video, animation, document & audio) in Telegram bot API
- Added support for `InputPaidMedia` (photo & video) in Telegram bot API
- Implemented handling of slices for both `InputMedia` and `InputPaidMedia`

## v0.2.0 (11-09-2025)

- Support media upload in form building
- Added parsing methods for union-type fields

## v0.1.0 (09-09-2025)

- initial release