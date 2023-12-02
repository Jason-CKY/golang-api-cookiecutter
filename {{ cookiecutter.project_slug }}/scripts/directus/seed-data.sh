ADMIN_ACCESS_TOKEN=$(curl -X POST -H "Content-Type: application/json" \
                        -d '{"email": "admin@example.com", "password": "d1r3ctu5"}' \
                        localhost:8055/auth/login \
                        | jq .data.access_token | cut -d '"' -f2)

# task table

for i in {0..8}; do
    DATA=$(cat scripts/directus/data.json | jq ".[$i]")
    curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d "$DATA" \
    localhost:8055/items/task
done


curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d  '{"status":"backlog","sorting_order":"[]"}'\
    localhost:8055/items/task_sorting

curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d  '{"status":"progress","sorting_order":"[]"}'\
    localhost:8055/items/task_sorting

curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $ADMIN_ACCESS_TOKEN" \
    -d  '{"status":"done","sorting_order":"[]"}'\
    localhost:8055/items/task_sorting