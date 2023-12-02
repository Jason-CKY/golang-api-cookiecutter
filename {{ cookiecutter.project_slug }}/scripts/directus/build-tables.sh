
ADMIN_ACCESS_TOKEN=$(curl -X POST -H "Content-Type: application/json" \
                        -d '{"email": "admin@example.com", "password": "d1r3ctu5"}' \
                        localhost:8055/auth/login \
                        | jq .data.access_token | cut -d '"' -f2)

# task table
curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"collection":"task","fields":[{"field":"id","type":"uuid","meta":{"hidden":true,"readonly":true,"interface":"input","special":["uuid"]},"schema":{"is_primary_key":true,"length":36,"has_auto_increment":false}},{"field":"date_created","type":"timestamp","meta":{"special":["date-created"],"interface":"datetime","readonly":true,"hidden":true,"width":"half","display":"datetime","display_options":{"relative":true}},"schema":{}},{"field":"date_updated","type":"timestamp","meta":{"special":["date-updated"],"interface":"datetime","readonly":true,"hidden":true,"width":"half","display":"datetime","display_options":{"relative":true}},"schema":{}}],"schema":{},"meta":{"singleton":false}}' \
    localhost:8055/collections

# task fields
curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"field":"title","type":"string","schema":{},"meta":{"interface":"input","special":null,"required":true},"collection":"task"}' \
    localhost:8055/fields/task \

curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"field":"description","type":"string","schema":{},"meta":{"interface":"input","special":null,"required":true},"collection":"task"}' \
    localhost:8055/fields/task \

curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"field":"status","type":"string","schema":{},"meta":{"interface":"select-dropdown","special":null,"required":true,"options":{"choices":[{"text":"backlog","value":"backlog"},{"text":"progress","value":"progress"},{"text":"done","value":"done"}]}},"collection":"task"}' \
    localhost:8055/fields/task \

# task_sorting table
curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"collection":"task_sorting","fields":[{"field":"id","type":"uuid","meta":{"hidden":true,"readonly":true,"interface":"input","special":["uuid"]},"schema":{"is_primary_key":true,"length":36,"has_auto_increment":false}},{"field":"date_created","type":"timestamp","meta":{"special":["date-created"],"interface":"datetime","readonly":true,"hidden":true,"width":"half","display":"datetime","display_options":{"relative":true}},"schema":{}},{"field":"date_updated","type":"timestamp","meta":{"special":["date-updated"],"interface":"datetime","readonly":true,"hidden":true,"width":"half","display":"datetime","display_options":{"relative":true}},"schema":{}}],"schema":{},"meta":{"singleton":false}}' \
    localhost:8055/collections \

# task_sorting fields
curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"field":"status","type":"string","schema":{},"meta":{"interface":"select-dropdown","special":null,"required":true,"options":{"choices":[{"text":"backlog","value":"backlog"},{"text":"progress","value":"progress"},{"text":"done","value":"done"}]}},"collection":"task_sorting"}' \
    localhost:8055/fields/task_sorting \

curl -X POST -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d '{"field":"sorting_order","type":"json","schema":{},"meta":{"interface":"input-code","special":["cast-json"],"required":true},"collection":"task_sorting"}' \
    localhost:8055/fields/task_sorting \








