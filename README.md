# ledger-demo

Sandbox service for the Countersign POC. `invoice` computes invoice totals (subtotal, tax, total) in integer cents. Merges to `main` run a simulated deploy that creates a GitHub Deployment to `production`.




## Try sdocs

This repo doubles as the sandbox for sdocs (the keddaro project): the pages in `docs/` have
sections only some teams can read. In git they are encrypted; `sdocs` decrypts the ones your GitHub team allows.

| Doc | juniors | seniors | admins |
|---|---|---|---|
| `docs/runbook.md` | on-call rota | rollback steps | break-glass access (hidden) |
| `docs/architecture.md` | | rounding edge case, security notes (hidden) | enterprise pricing |
| `docs/onboarding.md` | first-week plan | | |

Roles come from teams in `mercury-test-org` (`juniors` < `seniors` < `admins`); a higher role sees everything below it.

**Setup** (once):

1. Start the portal: in the keddaro repo, `dotnet run --project src/Sdocs.AppHost --launch-profile http` (portal on http://localhost:5180).
2. Get the CLI on your PATH: in the keddaro repo, `dotnet publish src/Sdocs.Cli -c Release -r osx-arm64 -o ~/.local/bin`.
3. In this repo: `sdocs init` (installs the git hooks) and `sdocs login`.

**Use**:

```sh
sdocs whoami                 # your login and role
sdocs unseal docs/*.md       # decrypt what your role allows; the rest says "skipped ... requires <role>"
# edit, then commit as usual: the pre-commit hook re-seals everything
git checkout -- docs         # or throw the plaintext away
```

To see a section disappear, move yourself to a lower team in `mercury-test-org` and run `sdocs unseal` again
(takes effect immediately). Sign-in lasts an hour; run `sdocs login` again when told to.
