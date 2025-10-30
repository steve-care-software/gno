#!/usr/bin/env bash
set -euo pipefail

OLD="github.com/gnolang/gno"
NEW="github.com/steve-care-software/gno"

echo "==> Snapshot current state"
git add -A && git commit -m "pre-rename snapshot" || true

echo "==> Update module paths in ALL go.mod files"
while IFS= read -r -d '' gomod; do
  moddir="$(dirname "$gomod")"
  oldmod="$(awk '/^module /{print $2}' "$gomod")"
  newmod="${oldmod/$OLD/$NEW}"
  if [[ "$oldmod" != "$newmod" ]]; then
    ( cd "$moddir" && go mod edit -module "$newmod" )
    echo "    $oldmod  ->  $newmod"
  fi
done < <(find . -type f -name go.mod -not -path '*/.git/*' -print0)

echo "==> Rewrite ALL references to imports/paths"
find . -type f -not -path '*/.git/*' -print0 \
| xargs -0 grep -IlZ "$OLD" \
| xargs -0 sed -i '' "s|$OLD|$NEW|g"

echo "==> Tidy each module"
while IFS= read -r -d '' gomod; do
  moddir="$(dirname "$gomod")"
  ( cd "$moddir" && go mod tidy )
done < <(find . -type f -name go.mod -not -path '*/.git/*' -print0)

echo "==> Done. Verify with:"
echo "    grep -R \"$OLD\" . || echo 'No old imports left.'"
