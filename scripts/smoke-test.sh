#!/usr/bin/env bash

# set -x
set -e

BASE_URL=${BASE_URL:-http://localhost:8888}

OUTPUT_DIR=scripts/smoke-test

URL="$BASE_URL/content"
echo -e "\nCreate - POST $URL\n"
curl -s --show-error -X POST $URL -H "Content-Type: application/json" -d '{"title":"Test Content","body":"This is test content","author":"Test Author","status":"draft"}' -o $OUTPUT_DIR/create.json -D /tmp/headers
cat /tmp/headers
cat $OUTPUT_DIR/create.json | jq

URL="$BASE_URL/content"
echo -e "\nList - GET $URL\n"
curl -s --show-error $URL -o $OUTPUT_DIR/list.json -D /tmp/headers
cat /tmp/headers
cat $OUTPUT_DIR/list.json | jq

CONTENT_ID=$(curl -s --show-error $BASE_URL/content | jq -r '.[0].id')

URL="$BASE_URL/content/$CONTENT_ID"
echo -e "\nGet - GET $URL\n"
curl -s --show-error $URL -o $OUTPUT_DIR/get.json -D /tmp/headers
cat /tmp/headers
cat $OUTPUT_DIR/get.json | jq

URL="$BASE_URL/content/$CONTENT_ID"
echo -e "\nUpdate - PUT $URL\n"
curl -s --show-error -X PUT $URL \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "body": "Updated body text",
    "author": "Updated Author",
    "status": "published",
    "data": {"foo": "bar"}
  }' -o $OUTPUT_DIR/update.json -D /tmp/headers
cat /tmp/headers
cat $OUTPUT_DIR/update.json | jq

URL="$BASE_URL/content/$CONTENT_ID"
echo -e "\nDelete - DELETE $URL\n"
curl -s --show-error -X DELETE $URL -o $OUTPUT_DIR/delete.json -D /tmp/headers
cat /tmp/headers
cat $OUTPUT_DIR/delete.json | jq
status_code=$(curl -s --show-error $URL -w '%{http_code}' -o /dev/null)
if [ "$status_code" -ne "404" ]; then
  echo -e "\nFAILURE!: expected status code 404 after delete but got $status_code\n"
  exit 1
fi

echo -e "\nSUCCESS!"
