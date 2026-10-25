# Ledger runbook

How we keep the `invoice` service healthy. Everyone on the team can read this page; some sections open up
depending on your role.

## Health checks

- `GET /healthz` returns `200` when the service can reach its database.
- Dashboards: *Ledger / Overview* in Grafana.

## On-call

<restricted role="juniors" id="rb_01M5S3BV87DYVZV4Z3PFT3B0B0">
On-call rotation: one engineer per week, handover Monday 10:00 CET in #ledger-oncall.
Page the secondary (a senior) if an incident is not mitigated within 30 minutes.
</restricted>

## Rolling back a release

<restricted role="seniors" id="rb_01M5S3BTF3NVEP044C5SAQZ5HJ">
1. Find the last good deployment: `gh api repos/mercury-test-org/ledger-demo/deployments --jq '.[1].sha'`.
2. Re-run the `deploy` workflow for that SHA.
3. Post the rollback in #ledger-oncall with the incident link.

Never roll back past a migration that dropped a column; restore from the nightly snapshot instead.
</restricted>

## Break-glass access

<restricted role="admins" visibility="hidden" id="rb_01M5S3BSN5QEWE7D3H96XF2GFG">
Break-glass credentials for the production database live in Key Vault `ledger-prod` as `db-breakglass`.
Using them pages the security channel automatically. Rotate them within 24 hours of use.
</restricted>

## Escalation

<restricted role="seniors">
Escalate to the CTO if the ledger is down for more than an hour.
</restricted>
