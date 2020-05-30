#!/usr/bin/env zsh
export TARGET_SHELL=zsh
source "$(dirname $0)/direnv-test-common.sh"

test_start load-hooks
  direnv_eval
  test_eq "$DIRENV_POSTLOAD" "alias foo='true'"
  test_eq "$DIRENV_PREUNLOAD" "unalias foo"
  # This is aliased now and should succeed
  foo

  cd ..
  direnv_eval
  test_eq "$DIRENV_POSTLOAD" ""
  test_eq "$DIRENV_PREUNLOAD" ""
  if [[ foo == 0 ]]; then
      exit 1
  fi
test_stop
