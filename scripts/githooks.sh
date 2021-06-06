rm ./.git/hooks/pre-commit
ln -s ../../githooks/pre-commit ./.git/hooks/pre-commit

rm ./.git/hooks/commit-msg
ln -s ../../githooks/commit-msg ./.git/hooks/commit-msg
