# Invoice service architecture

`invoice` computes subtotal, tax and total in integer cents, so no floating point ever touches money.

## Components

- **API**: stateless HTTP service, scales horizontally.
- **Rates**: tax rates loaded at start-up from `rates.json`.

## Rounding

Tax is computed per line and rounded half-to-even, then summed.

<restricted role="seniors" id="rb_01M5S3BQCKVRHZHMYB303V2SMG">
Known edge case: invoices with more than 10,000 lines can drift by up to 1 cent against the finance team's
spreadsheet, because they round once on the total. Finance has accepted this; don't "fix" it without them.
</restricted>

## Enterprise pricing

<restricted role="admins" id="rb_01M5S3BNTZ6SN50KGDC4A3KH9E">
Enterprise customers get a negotiated discount applied before tax. Current tiers: 8% above 50k invoices/month,
12% above 200k. Contract owners: the sales lead and the CFO.
</restricted>

<restricted role="seniors" visibility="hidden" id="rb_01M5S3BJFZG045EK0TZKFHXG13">
Security review notes (2026-09): rate files are not signed yet; anyone with repo write access could change tax
rates. Tracked as LEDGER-142, target Q1.
</restricted>
