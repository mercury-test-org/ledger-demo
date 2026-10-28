# Secrets demo

Live credentials stay in Key Vault; git only holds references. Click **Reveal** to fetch a value (audited, hidden again after 30 s).

Demo database password, revealable by `juniors` and up:

<secret ref="kv:demo-db-password" />

Admin API key. The tag says `juniors`, but the secret itself is stored as `admins`, and the secret's own role wins:

<secret ref="kv:demo-admin-key" role="juniors" />

The same database password, but this tag raises the bar to `seniors`, so juniors see a lock:

<secret ref="kv:demo-db-password" role="seniors" />

See also the [architecture overview](architecture.md) and the [ledger code](../README.md).
