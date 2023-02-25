#!/usr/bin/env sh

: ${LITESTREAM_CONFIG_PATH:=/etc/litestream.yml}
echo "Using configuration: ${LITESTREAM_CONFIG_PATH}"

echo "Expecting database: ${DB_PATH}"

# Restore the database if it does not already exist.
if [ -f "${DB_PATH}" ]; then
  echo "Database exists, skipping restore."
else
  echo "No database found, attempting to restore from a replica."
  litestream restore \
    -config "${LITESTREAM_CONFIG_PATH}" \
    -if-replica-exists "${DB_PATH}"
  echo "Finished restoring the database."
fi

# Run app through Litestream exec.
exec litestream replicate \
  -config "${LITESTREAM_CONFIG_PATH}" \
  -exec "${LITESTREAM_EXEC}"
