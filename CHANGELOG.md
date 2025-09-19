# Changelog

## 0.4.2 (19-09-2025)

- Middleware execution optimized: avoids unnecessary slice copying
- Added thread-safety for middleware and handler execution using mutex/atomic
- `Next()` flow updated for better performance and safety

## 0.4.1 (17-09-2025)

- Support multiple handlers in command registration:
  - `OnStartCommand`, `OnHelpCommand`, and `RegisterCommandHandler` now accept variadic `...Handler`.
  - Middleware and handlers are combined into a single execution stack inside `RegisterHandler`.
  - Backward compatibility with existing single-handler usage is preserved.

## 0.4.0 (16-09-2025)

- Refactored middleware system to per-command middleware:
  - Each registered command now maintains its own middleware stack.
  - Global middlewares are still supported via `Use()`.
  - Middleware chain is now executed sequentially per command without rebuilding the full stack on each update.
- Introduced `Context` struct to wrap `bot` and `update` parameters:
  - All handlers now receive `*Context` instead of `(ctx, bot, update)`.
  - Added `Next()` and `Abort()` methods to control middleware flow.
- Simplified handler signatures and improved error propagation within middleware chain.

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